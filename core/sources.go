package core

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/rabbit-backend/go-tiles/db"
	engine "github.com/rabbit-backend/template"
)

type DBConnection struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Source struct {
	Name       string       `json:"name"`
	Connection DBConnection `json:"connection"`
	Type       string       `json:"type"`
}

type GoTilesConfig struct {
	Sources []Source `json:"sources"`
}

func (conn DBConnection) GetConnectionURL() string {
	if conn.Type == "env" {
		return os.Getenv(conn.Value)
	}

	return conn.Value
}

func GetConfig() GoTilesConfig {
	file, err := os.Open("gotiles.json")
	if err != nil {
		log.Fatalln(err)
	}

	var config GoTilesConfig
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalln(err)
	}

	return config
}

func (config GoTilesConfig) GetConnections(e *engine.Engine) map[string]db.DBSource {
	connections := make(map[string]db.DBSource, 0)

	for _, source := range config.Sources {
		driver, ok := db.DB_SOURCES[source.Type]
		if !ok {
			log.Fatalln("[x] Invalid datasource:", source.Type, " / Try implementing the driver for this type")
		}

		conn := driver(e)
		conn.Open(source.Connection.GetConnectionURL()) // try to connect to the database

		connections[source.Name] = conn
	}

	return connections
}
