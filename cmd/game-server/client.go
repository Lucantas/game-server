package main

import (
	"fmt"
	"net"
	"strings"
)

// Client type represents a client with a connection with the server
type Client struct {
	Conn
}

func (c *Client) connect(nc Client) {
	conn, err := net.Dial("tcp", nc.Conn.Host)

	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
}

func (c *Client) address() string {
	return c.Host + ":" + c.Port
}

func newClient(address string, network string) Client {
	var host string
	var port string
	if strings.Contains(address, "[") {
		// IPV6
		addr := strings.SplitAfter(address, "]")
		host = addr[0]
		port = addr[1]
	} else {
		// IPV4
		addr := strings.Split(address, ":")
		host = addr[0]
		port = addr[1]
	}
	return Client{Conn{host, port, network}}
}
