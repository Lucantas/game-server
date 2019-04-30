package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucantas/game-server/pkg/easyio"
)

// connection commands
const (
	connect    = "connect"
	spread     = "spread"
	disconnect = "disconnect"
)

var (
	hub   Hub
	games = []Match{}
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

	if hub = conf.Hub; hub.Enabled {
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
	var p Payload
	var buf bytes.Buffer
	_, err := io.Copy(&buf, conn)

	if err != nil {
		log.Println("Error while reading from buffer: ", err)
	}

	content := buf.Bytes()
	err = json.Unmarshal(content, &p)

	if err != nil {
		log.Println("Error while unmarshalling: ", err)
	}

	switch strings.ToLower(p.Method) {
	case connect:
		handleConnection(conn.RemoteAddr())
	case spread:
		spreadData(conn.RemoteAddr(), p)
	case disconnect:
		conn.Close()
	}
}

func handleConnection(addr net.Addr) {
	if client := addr.String(); !badUser(client) {
		// client is not on the list of banned addresses
		// move him to the waiting list
		hub.add(newClient(client, addr.Network()))
	}
}

func spreadData(addr net.Addr, p Payload) {
	if validMatch(p.MatchID, addr.String()) {
		// only spread data if the payload have a valid matchID
	}
}

func badUser(address string) bool {
	// TODO: read the list of banned addres
	return false
}

func validMatch(matchID int64, addr string) bool {
	if matchID > 0 {
		for _, match := range games {
			if match.ID == matchID {
				// valid match
				for _, client := range match.Clients {
					if client.address() == addr {
						// client belong to that match
						return true
					}
				}
			}
		}
	}
	return false
}
