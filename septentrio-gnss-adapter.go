package main

import (
	// 	"bytes"
	// 	"math/rand"
	// 	"os/exec"
	// 	"sync"
	// 	"time"

	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	sbf "Septentrio-GNSS-Adapter/sbf"

	adapter_library "github.com/clearblade/adapter-go-library"
	mqttTypes "github.com/clearblade/mqtt_parsing"

	"go.bug.st/serial"
)

const (
	adapterName                    = "septentrio-gnss-adapter"
	platURL                        = "http://localhost:9000"
	messURL                        = "localhost:1883"
	msgSubscribeQos                = 0
	msgPublishQos                  = 0
	portRead                       = "receive"
	portWrite                      = "send"
	adapterConfigCollectionDefault = "adapter_config"
)

var (
	adapterConfig   *adapter_library.AdapterConfig
	adapterSettings *SeptentrioGNSSAdapterSettings

	platformURL             string //Defaults to http://localhost:9000
	messagingURL            string //Defaults to localhost:1883
	sysKey                  string
	sysSec                  string
	deviceName              string //Defaults to xDotSerialAdapter
	activeKey               string
	logLevel                string //Defaults to info
	adapterConfigCollection string
	readInterval            int
	readTimeout             int

	serialPortName = ""

	topicRoot = "wayside/lora"

	cbSubscribeChannel <-chan *mqttTypes.Publish
	endWorkersChannel  chan string

	port   interface{}
	buffer []byte
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

	validateAdapterSettings()

	err = adapter_library.ConnectMQTT(adapterConfig.TopicRoot+"/#", cbMessageHandler)
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
	endWorkersChannel <- "Stop Channel"
	endWorkersChannel <- "Stop Channel"

	//stop serial data mode when adapter is killed
	// log.Println("[INFO] Stopping Serial Data Mode...")
	// if err := serialPort.StopSerialDataMode(); err != nil {
	// 	log.Println("[WARN] initCbClient - Error stopping serial data mode: " + err.Error())
	// }

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
			log.Fatal("[FATAL] host is required in adapter settings when connection type is set to 'serial'\n")
		}

		if adapterSettings.TcpPort == 0 {
			log.Fatal("[FATAL] port is required in adapter settings when connection type is set to 'serial'\n")
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
	if adapterSettings.ConnectionType == "serial" {
		readFromSerialPort()
	} else if adapterSettings.ConnectionType == "tcp" {
		readFromTcpPort()
	}
}

// Publishes data to a topic
func publish(topic string, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("[ERROR] Failed to stringify JSON: %s\n", err.Error())
		return
	}

	log.Printf("[DEBUG] publish - Publishing to topic %s\n", topic)
	err = adapter_library.Publish(topic, b)
	if err != nil {
		log.Printf("[ERROR] Failed to publish MQTT message to topic %s: %s\n", topic, err.Error())
	}
}

func readFromSerialPort() {

	var err error
	port, err = serial.Open(adapterSettings.SerialPort, getSerialMode())
	if err != nil {
		log.Fatalf("[ERROR] readFromSerialPort - Error opening serial port: %s\n", err.Error())
		return
	}

	defer port.(serial.Port).Close()

	for {
		select {
		case <-endWorkersChannel:
			log.Println("[DEBUG] readWorker - stopping serial read worker")
			return
		default:
			buff := make([]byte, 4096)

			n, err := port.(serial.Port).Read(buff)
			if err != nil {
				log.Printf("[ERROR] readFromSerialPort - Error reading from serial port: %s\n", err.Error())
			}
			if n > 0 {
				log.Printf("[DEBUG] readFromSerialPort - %d bytes read from serial port\n", n)
				log.Printf("[DEBUG] %v\n", buff[:n])

				buffer = append(buffer, buff[:n]...)

				sbf.Parse(&buffer)
			}
		}
	}

	// data, err := serialPort.ReadSerialPort()

	// if err != nil && err != io.EOF {
	// 	log.Printf("[ERROR] readFromSerialPort - ERROR reading from serial port: %s\n", err.Error())
	// } else {
	// 	if data != "" {
	// 		//If there are any slashes in the data, we need to escape them so duktape
	// 		//doesn't throw a SyntaxError: unterminated string (line 1) error
	// 		data = strings.Replace(data, `\`, `\\`, -1)

	// 		log.Printf("[INFO] readFromSerialPort - Data read from serial port: %s\n", data)

	// 		//Publish data to message broker
	// 		err := publish(topicRoot+"/"+serialRead+"/response", data)
	// 		if err != nil {
	// 			log.Printf("[ERROR] readFromSerialPort - ERROR publishing to topic: %s\n", err.Error())
	// 		}
	// 	} else {
	// 		log.Println("[DEBUG] readFromSerialPort - No data read from serial port, skipping publish.")
	// 	}
	// }
}

func readFromTcpPort() {
	var err error
	port, err = net.Dial("tcp", adapterSettings.TcpHost+":"+strconv.Itoa(adapterSettings.TcpPort))
	if err != nil {
		log.Fatalf("[ERROR] readFromTcpPort - Error opening tcp port: %s\n", err.Error())
		return
	}

	defer port.(net.Conn).Close()

	for {
		select {
		case <-endWorkersChannel:
			log.Println("[DEBUG] readFromTcpPort - stopping tcp read worker")
			return
		default:
			buff := make([]byte, 4096)

			n, err := port.(net.Conn).Read(buff)
			if err != nil {
				log.Printf("[ERROR] readFromTcpPort - Error reading from tcp port: %s\n", err.Error())
			}
			if n > 0 {
				log.Printf("[DEBUG] readFromTcpPort - %d bytes read from tcp port\n", n)
				log.Printf("[DEBUG] %v\n", buff[:n])

				buffer = append(buffer, buff[:n]...)
				sbf.Parse(&buffer)
			}
		}
	}
}

func writeToPort(payload []byte) {
	if adapterSettings.ConnectionType == "serial" {
		if port != nil {
			writeToSerialPort(payload)
		} else {
			log.Print("[ERROR] writeToPort - Cannot write to serial port. Port not open.\n")
		}
	} else if adapterSettings.ConnectionType == "tcp" {
		if port != nil {
			writeToTcpPort(payload)
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
