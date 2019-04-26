package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/lucantas/game-server/pkg/easyio"
	"github.com/lucantas/game-server/pkg/strutil"
)

// connection commands
const (
	connect = "connect"
)

var (
	hub = []Client{}
)

// App holds all the settings to the service
type App struct {
}

// Conf type is a model of the json conf
type Conf struct {
	Hub  `json:"hub"`
	Conn `json:"connection"`
}

func appInit() *App {
	app := &App{}
	app.run()
	return app
}

var conf Conf

func (a *App) run() {
	// find the executable current directory
	ex, err := os.Executable()
	if err != nil {
		log.Print("Error while opening executable: ", err)
	}
	exPath := filepath.Dir(ex)
	fileName := easyio.FileName{Path: exPath}
	confFile := fmt.Sprintf("%s/config/conf.json", fileName.GetParent())
	confBytes, err := easyio.ReadFile(confFile)

	if err != nil {
		log.Print("Error while reading from file: ", err)
	}

	json.Unmarshal(confBytes, &conf)

	if hub := conf.Hub; hub.Enabled {
		server := conf.Conn
		host = server.Host
		port = server.Port
		protocol = server.Protocol
		hub.ID = "1"
		log.Println("Starting hub with Id ", hub.ID)
		listen()
	}
}

func listen() {
	l, err := net.Listen(protocol, fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Println("Error while listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	log.Println(fmt.Sprintf("Game server started on host %s:%s", host, port))
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: connection ", err.Error())
			os.Exit(1)
		}
		go handleRequests(conn)
	}
}

func handleRequests(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		log.Println("Error while reading request:", err.Error())
	}

	isConn, err := strutil.Equals(string(bytes.Trim(buf, "\x00")), connect)
	if err != nil {
		log.Println(err)
	}

	if isConn {
		handleConnection(conn.RemoteAddr())
	}

	conn.Write([]byte("Message received."))
	conn.Close()
}

func handleConnection(addr net.Addr) {
	if client := addr.String(); !badUser(client) {
		// client is not on the list of banned addresses
		// move him to the waiting list
		addToHub(newClient(client, addr.Network()))
		for _, c := range hub {
			log.Println(c)
		}
	}
}

func addToHub(c Client) {
	hub = append(hub, c)
}

func badUser(address string) bool {
	// TODO: read the list of banned addres
	return false
}
