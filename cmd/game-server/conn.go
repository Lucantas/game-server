package main

// Conn type represents the information of a connection
// that could be used by a client or services
type Conn struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
}
