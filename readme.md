# Septentrio GNSS Adapter

The __Septentrio GNSS__ adapter provides the ability for the ClearBlade platform to communicate with a Septentrio GNSS receiver (https://www.septentrio.com).

The adapter subscribes to MQTT topics which are used to interact with the Septentrio GNSS receiver. The adapter publishes any data retrieved from the receiver to MQTT topics so that the ClearBlade Platform is able to retrieve and process GNSS receiver related data or execute (write) commands to Septentrio GNSS receiver devices.

# MQTT Topic Structure
The Septentrio GNSS adapter utilizes MQTT messaging to communicate with the ClearBlade Platform. The Septentrio GNSS adapter will subscribe to a specific topic in order to handle Septentrio GNSS receiver requests. Additionally, the Septentrio GNSS adapter will publish messages to MQTT topics in order to send Septentrio GNSS receiver data to the ClearBlade Platform/Edge. The topic structures utilized by the Septentrio GNSS adapter are as follows:

  * Receive Septentrio GNSS data: {__TOPIC ROOT__}/receive/
  * Execute (write) receiver command request: {__TOPIC ROOT__}/request

## ClearBlade Platform Dependencies
The Septentrio GNSS adapter was constructed to provide the ability to communicate with a _System_ defined in a ClearBlade Platform instance. Therefore, the adapter requires a _System_ to have been created within a ClearBlade Platform instance.

Once a System has been created, artifacts must be defined within the ClearBlade Platform system to allow the adapters to function properly. At a minimum: 

  * A device needs to be created in the Auth --> Devices collection. The device will represent the adapter, or more importantly, the Septentrio GNSS receiver or the gateway on which the adapter is executing. The _name_ and _active key_ values specified in the Auth --> Devices collection will be used by the adapter to authenticate to the ClearBlade Platform or ClearBlade Edge. 
  * An adapter configuration data collection needs to be created in the ClearBlade Platform _system_ and populated with the data appropriate to the Septentrio GNSS receiver installation. The schema of the data collection should be as follows:


| Column Name      | Column Datatype |
| ---------------- | --------------- |
| adapter_name     | string          |
| topic_root       | string          |
| adapter_settings | string (json)   |

### adapter_settings
The adapter_settings column will need to contain a JSON object containing the following attributes:

##### connectionType
* Either __serial__ or __tcp__



##### serialPortName
* The full unix path to the xDot serial device (ex. /dev/ttyAP1)

##### transmissionDataRate
* DR0-DR15 can be used
* See https://www.multitech.com/documents/publications/manuals/s000643.pdf for further information

##### transmissionFrequency
* The transmit frequency to use in peer-to-peer mode
* Use 915.5-919.7 MhZ for US 915 devices to avoid interference with LoRaWAN networks

#### adapter_settings_examples

##### TCP connection type example
{  
  "connectionType": "tcp",
  "host": "localhost",
  "tcpPort": 28784
}

##### Serial connection type example
  * Note: hardware and software flow control is currently not supported.

{  
  "connectionType": "serial"
  "serialPort":"/dev/ttyAP1", 
  "baudRate": 115200, //75, 110, 300, 1200, 2400, 4800, 9600, 19200, 38400, 57600 and 115200 - default value is 115200
  "dataBits": 8, //5, 6, 7, 8, 9 - default value is 8
  "parity": "none", //none|odd|even|mark|space - default value is none
  "stopBits": 1, //1, 1.5, 2 - default value is 1
  "readTimeout": 1 //Number of seconds
}

## Usage

### Executing the adapter

`septentrio-gnss-adapter -systemKey=<SYSTEM_KEY> -systemSecret=<SYSTEM_SECRET> -platformURL=<PLATFORM_URL> -messagingURL=<MESSAGING_URL> -deviceName=<DEVICE_NAME> -password=<DEVICE_ACTIVE_KEY> -adapterConfigCollection=<COLLECTION_NAME> -logLevel=<LOG_LEVEL>`

   __*Where*__ 

   __systemKey__
  * REQUIRED
  * The system key of the ClearBLade Platform __System__ the adapter will connect to

   __systemSecret__
  * REQUIRED
  * The system secret of the ClearBLade Platform __System__ the adapter will connect to
   
   __deviceName__
  * The device name the adapter will use to authenticate to the ClearBlade Platform
  * Requires the device to have been defined in the _Auth - Devices_ collection within the ClearBlade Platform __System__
  * OPTIONAL
  * Defaults to __xDotSerialAdapter__
   
   __password__
  * REQUIRED
  * The active key the adapter will use to authenticate to the platform
  * Requires the device to have been defined in the _Auth - Devices_ collection within the ClearBlade Platform __System__
   
   __platformUrl__
  * The url of the ClearBlade Platform instance the adapter will connect to
  * OPTIONAL
  * Defaults to __http://localhost:9000__

   __messagingUrl__
  * The MQTT url of the ClearBlade Platform instance the adapter will connect to
  * OPTIONAL
  * Defaults to __localhost:1883__

   __adapterConfigCollection__
  * REQUIRED 
  * The collection name of the data collection used to house adapter configuration data

   __logLevel__
  * The level of runtime logging the adapter should provide.
  * Available log levels:
    * fatal
    * error
    * warn
    * info
    * debug
  * OPTIONAL
  * Defaults to __info__


## Setup
---
The xdot adapters are dependent upon the ClearBlade Go SDK and its dependent libraries being installed. The xDot adapter was written in Go and therefore requires Go to be installed (https://golang.org/doc/install).


### Adapter compilation
In order to compile the adapter for execution within mLinux, the following steps need to be performed:

 1. Retrieve the adapter source code  
    * ```git clone git@github.com:ClearBlade/xDot-Adapter.git```
 2. Navigate to the xdotadapter directory  
    * ```cd xdotadapter```
 4. Compile the adapter
    * ```GOARCH=arm GOARM=5 GOOS=linux go build```



