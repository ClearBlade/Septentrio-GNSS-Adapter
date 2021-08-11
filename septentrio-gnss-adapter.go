package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	sbf "Septentrio-GNSS-Adapter/sbf"

	adapter_library "github.com/clearblade/adapter-go-library"
	mqttTypes "github.com/clearblade/mqtt_parsing"

	"go.bug.st/serial"
)

//TODO
//
// 1. IN PROGRESS - Implement read timeout for serial and tcp
// 2. Implement handling of prompts
// 3. Implement handling of ASCII command reply data
// 4. Implement handling of ASCII Display data
// 5. Implement handling of event data
// 6. Implement handling of formatted information block data
// 7. Implement parsing of remaining sbf blocks
// 8. Implement handling of commands published to adapter

const (
	adapterName      = "septentrio-gnss-adapter"
	portWriteRequest = "request"
	portDataReceived = "receive"
)

var (
	adapterConfig     *adapter_library.AdapterConfig
	adapterSettings   *SeptentrioGNSSAdapterSettings
	publishTopic      string
	endWorkersChannel chan string
	port              interface{}
	buffer            []byte //Buffer containing all serial/tcp data waiting to be processed
	parsingIsRunning  bool
)

func main() {
	fmt.Println("Starting septentrioGNSSAdapter...")

	err := adapter_library.ParseArguments(adapterName)
	if err != nil {
		log.Fatalf("[FATAL] Failed to parse arguments: %s\n", err.Error())
	}

	adapterConfig, err = adapter_library.Initialize()
	if err != nil {
		log.Fatalf("[FATAL] Failed to initialize: %s\n", err.Error())
	}

	adapterSettings = &SeptentrioGNSSAdapterSettings{}
	err = json.Unmarshal([]byte(adapterConfig.AdapterSettings), adapterSettings)
	if err != nil {
		log.Fatalf("[FATAL] Failed to parse Adapter Settings %s\n", err.Error())
	}

	publishTopic = adapterConfig.TopicRoot + "/" + portDataReceived

	validateAdapterSettings()

	err = adapter_library.ConnectMQTT(adapterConfig.TopicRoot+"/"+portWriteRequest, cbMessageHandler)
	if err != nil {
		log.Fatalf("[FATAL] Failed to connect MQTT: %s\n", err.Error())
	}

	defer close(endWorkersChannel)
	endWorkersChannel = make(chan string)

	//Start read loop
	go readWorker()

	//Handle OS interrupts to shut down gracefully
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c

	log.Printf("[INFO] OS signal %s received, ending go routines.", sig)

	//End the existing goRoutines
	//TODO - This does not appear to be working
	//endWorkersChannel <- "Stop"

	os.Exit(0)
}

func validateAdapterSettings() {
	//Validate connection type
	if adapterSettings.ConnectionType == "" {
		log.Fatal("[FATAL] Connection type is required in adapter settings\n")
	} else if !(adapterSettings.ConnectionType == "serial" || adapterSettings.ConnectionType == "tcp") {
		log.Fatalf("[FATAL] Invalid connection type specified in adapter settings: %s\n", adapterSettings.ConnectionType)
	}

	if adapterSettings.ConnectionType == "serial" {
		//Validate serial fields
		if adapterSettings.SerialPort == "" {
			log.Fatal("[FATAL] serialPort is required in adapter settings when connection type is set to 'serial'\n")
		}

		if adapterSettings.BaudRate == 0 {
			log.Println("[DEBUG] Defaulting baud rate to 115200")
			adapterSettings.BaudRate = 115200
		} else {
			log.Printf("[DEBUG] Setting baud rate to %d\n", adapterSettings.BaudRate)
		}

		if adapterSettings.Size == 0 {
			log.Println("[DEBUG] Defaulting data bits to 8")
			adapterSettings.Size = 8
		} else {
			log.Printf("[DEBUG] Setting data bits to %d\n", adapterSettings.Size)
		}

		if adapterSettings.Parity == "" {
			log.Println("[DEBUG] Defaulting parity to none")
			adapterSettings.Parity = "none"
		} else {
			log.Printf("[DEBUG] Setting parity to %s\n", adapterSettings.Parity)
		}

		if adapterSettings.StopBits == 0 {
			log.Println("[DEBUG] Defaulting stop bits to 1")
			adapterSettings.StopBits = 1
		} else {
			log.Printf("[DEBUG] Setting stop bits to %f\n", adapterSettings.StopBits)
		}

		if adapterSettings.Timeout == 0 {
			log.Println("[DEBUG] Defaulting timeout to 1 second")
			adapterSettings.Timeout = 1000
		} else {
			log.Printf("[DEBUG] Setting timeout to %d milliseconds\n", adapterSettings.Timeout)
		}
	} else if adapterSettings.ConnectionType == "tcp" {
		//Validate tcp fields
		if adapterSettings.TcpHost == "" {
			log.Fatal("[FATAL] host is required in adapter settings when connection type is set to 'tcp'\n")
		}

		if adapterSettings.TcpPort == 0 {
			log.Fatal("[FATAL] port is required in adapter settings when connection type is set to 'tcp'\n")
		}
	}
}

func getSerialMode() *serial.Mode {
	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: adapterSettings.BaudRate,
		DataBits: adapterSettings.Size,
	}

	switch adapterSettings.Parity {
	case "none":
		mode.Parity = serial.NoParity
	case "odd":
		mode.Parity = serial.OddParity
	case "even":
		mode.Parity = serial.EvenParity
	case "mark":
		mode.Parity = serial.MarkParity
	case "space":
		mode.Parity = serial.SpaceParity
	default:
		log.Fatalf("[FATAL] Invalid parity specified: %s\n", adapterSettings.Parity)
	}

	switch adapterSettings.StopBits {
	case 1:
		mode.StopBits = serial.OneStopBit
	case 1.5:
		mode.StopBits = serial.OnePointFiveStopBits
	case 2:
		mode.StopBits = serial.TwoStopBits
	default:
		log.Fatalf("[FATAL] Invalid stop bits specified specified: %f\n", adapterSettings.StopBits)
	}

	return mode
}

func readWorker() {
	log.Println("[INFO] readWorker - Starting readWorker")
	var err error
	var n int

	//Open the serial port or tcp connection
	if adapterSettings.ConnectionType == "serial" {
		err = createSerialPort()
	} else if adapterSettings.ConnectionType == "tcp" {
		err = createTcpPort()
	}

	if err != nil {
		log.Fatalf("[FATAL] readWorker - Error creating connection to GNSS receiver: %s\n", err.Error())
	}

	//Ensure we close the serial port or tcp connection when we exit
	defer func() {
		log.Printf("[DEBUG] readWorker - Closing %s port\n", adapterSettings.ConnectionType)
		if adapterSettings.ConnectionType == "serial" {
			port.(serial.Port).Close()
		} else {
			port.(net.Conn).Close()
		}
	}()

	for {
		select {
		case <-endWorkersChannel:
			log.Println("[DEBUG] readWorker - stopping read worker")
			return
		default:
			buff := make([]byte, 4096)

			//Read from the serial port or tcp connection
			if adapterSettings.ConnectionType == "serial" {
				n, err = port.(serial.Port).Read(buff)
			} else {
				if adapterSettings.Timeout > 0 {
					port.(net.Conn).SetReadDeadline(time.Now().Add(time.Duration(adapterSettings.Timeout) * time.Second))
				}
				n, err = port.(net.Conn).Read(buff)
			}

			if err != nil {

				//Ignore TCP timeout errors
				if adapterSettings.ConnectionType == "tcp" {
					if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
						continue
					}
				}

				//TODO - Should we exit with a fatal error?
				//TODO - Need to see if we get EOF errors: Not receiving them with tcp
				log.Printf("[ERROR] readWorker - Error reading from %s port on GNSS Receiver: %+v\n", adapterSettings.ConnectionType, err)
			}

			if n > 0 {
				log.Printf("[DEBUG] readWorker - %d bytes read from %s port\n", n, adapterSettings.ConnectionType)
				buffer = append(buffer, buff[:n]...)
				parseAndPublishPayloads()
			}
		}
	}
}

func parseAndPublishPayloads() {
	parsedPayloads := []map[string]interface{}{}

	//We need to ensure that we only have 1 instance of the parsing goroutine
	//running at any one time
	if !parsingIsRunning {
		parsingIsRunning = true

		log.Println("[DEBUG] parseAndPublishPayloads - Begin parsing")
		go func() {
			sbf.Parse(&buffer, &parsedPayloads)

			//Only attempt to publish if we actually have payloads to publish
			if len(parsedPayloads) > 0 {
				publishPayloads(&parsedPayloads)
				parsedPayloads = nil
			}
			parsingIsRunning = false
			log.Println("[DEBUG] parseAndPublishPayloads - Parsing complete")
		}()
	} else {
		log.Println("[DEBUG] parseAndPublishPayloads - Parsing is already being executed")
	}
}

// Publishes data to a topic
func publishPayloads(payloads *[]map[string]interface{}) {

	for _, payload := range *payloads {
		jsonStr, err := json.Marshal(payload)
		if err != nil {
			log.Printf("[ERROR] Failed to marshall JSON %v: %s\n", payload, err.Error())
		} else {
			log.Printf("[DEBUG] publish - Publishing JSON %s to topic %s\n", jsonStr, publishTopic)
			err = adapter_library.Publish(publishTopic, jsonStr)
			if err != nil {
				log.Printf("[ERROR] Failed to publish MQTT message to topic: %s\n", err.Error())
			}
		}
	}
}

func createSerialPort() error {
	var err error
	port, err = serial.Open(adapterSettings.SerialPort, getSerialMode())
	if err != nil {
		log.Printf("[ERROR] createSerialPort - Error opening serial port %s: %s\n", adapterSettings.SerialPort, err.Error())
		return err
	}

	log.Printf("[DEBUG] createSerialPort - Serial port %s opened\n", adapterSettings.SerialPort)
	if adapterSettings.Timeout > 0 {
		port.(serial.Port).SetReadTimeout(time.Duration(adapterSettings.Timeout) * time.Second)
		log.Printf("[DEBUG] createSerialPort - Serial port read timeout set to %d seconds\n", adapterSettings.Timeout)
	}
	return nil
}

func createTcpPort() error {
	var err error
	addr := adapterSettings.TcpHost + ":" + strconv.Itoa(adapterSettings.TcpPort)

	port, err = net.Dial("tcp", addr)
	if err != nil {
		log.Printf("[ERROR] createTcpPort - Error opening tcp port %s: %s\n", addr, err.Error())
		return err
	}

	log.Printf("[DEBUG] createSerialPort - TCP port %s opened\n", addr)
	return nil
}

func writeToPort(payload []byte) {
	//We need to append the newline character to the bytes array to that the command
	//can be processed by the septentrio gnss receiver
	if adapterSettings.ConnectionType == "serial" {
		if port != nil {
			writeToSerialPort(append(payload, "\n"...))
		} else {
			log.Print("[ERROR] writeToPort - Cannot write to serial port. Port not open.\n")
		}
	} else if adapterSettings.ConnectionType == "tcp" {
		if port != nil {
			writeToTcpPort(append(payload, "\n"...))
		} else {
			log.Print("[ERROR] writeToPort - Cannot write to tcp port. Port not open.\n")
		}
	} else {
		log.Printf("[ERROR] writeToPort - Invalid port type configured: %s\n", adapterSettings.ConnectionType)
	}
}

func writeToSerialPort(bytes []byte) {
	n, err := port.(serial.Port).Write(bytes)
	if err != nil {
		log.Printf("[ERROR] writeToSerialPort - Error writing to serial port: %s\n", err.Error())
	}
	log.Printf("[DEBUG] writeToSerialPort - Wrote %d bytes to serial port\n", n)
}

func writeToTcpPort(bytes []byte) {
	n, err := port.(net.Conn).Write(bytes)
	if err != nil {
		log.Printf("[ERROR] writeToSerialPort - Error writing to tcp port: %s\n", err.Error())
	}
	log.Printf("[DEBUG] writeToSerialPort - Wrote %d bytes to tcp port\n", n)
}

func cbMessageHandler(message *mqttTypes.Publish) {
	//TODO - Add code to construct the appropriate payload
	writeToPort(message.Payload)
}
