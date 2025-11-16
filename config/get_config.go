package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rabbit-backend/go-tiles/db"
	engine "github.com/rabbit-backend/template"
)

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

	log.Println(config)

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
