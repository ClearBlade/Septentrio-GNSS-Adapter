package main

type SeptentrioGNSSAdapterSettings struct {
	ConnectionType string  `json:"connectionType"`
	TcpHost        string  `json:"host"`
	TcpPort        int     `json:"tcpPort"`
	SerialPort     string  `json:"serialPort"`
	BaudRate       int     `json:"baudRate"`
	Size           int     `json:"dataBits"`
	Parity         string  `json:"parity"`
	StopBits       float32 `json:"stopBits"`
	Timeout        int     `json:"readTimeout"`
}
