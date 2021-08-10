package main

type SeptentrioGNSSAdapterSettings struct {
	ConnectionType string  `json:"connectionType"`
	TcpHost        string  `json:"host,omitempty"`
	TcpPort        int     `json:"tcpPort,omitempty"`
	SerialPort     string  `json:"serialPort,omitempty"`
	BaudRate       int     `json:"baudRate,omitempty"`
	Size           int     `json:"dataBits,omitempty"`
	Parity         string  `json:"parity,omitempty"`
	StopBits       float32 `json:"stopBits,omitempty"`
	Timeout        int     `json:"readTimeout,omitempty"`
}
