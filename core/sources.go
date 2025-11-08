package core

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Source struct {
	Name       string       `json:"name"`
	Connection DBConnection `json:"connection"`
}

type GoTilesConfig struct {
	Sources []Source `json:"sources"`
}

func (conn DBConnection) GetConnection() (*sql.DB, error) {
	if conn.Type == "env" {
		return sql.Open("postgres", os.Getenv(conn.Value))
	}

	return sql.Open("postgres", conn.Value)
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

func (config GoTilesConfig) GetConnections() map[string]*sql.DB {
	connections := make(map[string]*sql.DB, 0)

	for _, source := range config.Sources {
		conn, err := source.Connection.GetConnection()
		if err != nil {
			log.Fatalln(err)
		}

		connections[source.Name] = conn
	}

	return connections
}
