package main

import (
	"fmt"
	"net"
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
