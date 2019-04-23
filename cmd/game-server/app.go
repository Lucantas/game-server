package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lucantas/game-server/pkg/easyio"
)

// App holds all the settings to the service
type App struct {
}

// Conf type is a model of the json conf
type Conf struct {
	Hub `json:"hub"`
}

func appInit() *App {
	app := &App{}
	app.run()
	return app
}

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
	var conf Conf

	if err != nil {
		log.Print(err)
	}

	json.Unmarshal(confBytes, &conf)

	if hub := conf.Hub; hub.Enabled {
		hub.ID = "1"
		log.Println("Starting hub with Id ", hub.ID)
	}
}
