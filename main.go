package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rabbit-backend/go-tiles/core"
	"github.com/rabbit-backend/go-tiles/db"
)

var DB *sql.DB

func init() {
	var err error

	godotenv.Load()
	// Connect to the Postgres Database
	if DB, err = db.ConnectDB(); err != nil {
		log.Fatalln("[x] error connecting db:", err)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/tiles/:tile/:x/:y/:z", func(c echo.Context) error {
		x := c.Param("x")
		y := c.Param("y")
		z := c.Param("z")
		tileName := c.Param("tile")

		query, _ := core.GetQuery(path.Join("tiles", "db", fmt.Sprintf("%s.sql", tileName)))
		
		var data []byte
		DB.QueryRow(query, x, y, z).Scan(&data)
	
		return c.Blob(http.StatusOK, "application/x-protobuf", data)
	})

	e.Logger.Fatal(e.Start(":3003"))
}
