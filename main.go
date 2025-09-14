package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rabbit-backend/go-tiles/core"
)

func init() {
	godotenv.Load()
}

func main() {
	config := core.GetConfig()
	connections := config.GetConnections()

	e := echo.New()
	e.Use(middleware.CORS())

	e.Static("/static", path.Join("tiles", "static"))

	e.GET("/tiles/:source/:tile/:x/:y/:z", func(c echo.Context) error {
		x := c.Param("x")
		y := c.Param("y")
		z := c.Param("z")
		tileName := c.Param("tile")
		source := c.Param("source")

		db := connections[source]
		query, err := core.GetQuery(path.Join(
				"tiles", 
				"db", 
				source, 
				fmt.Sprintf("%s.sql", tileName),
			),
		)

		if err != nil {
			return c.Blob(http.StatusInternalServerError, "application/x-protobuf", []byte(""))
		}
		
		var data []byte
		
		err = db.QueryRow(query, x, y, z).Scan(&data)
		if err != nil {
			return c.Blob(http.StatusInternalServerError, "application/x-protobuf", []byte(""))
		}

		return c.Blob(http.StatusOK, "application/x-protobuf", data)
	})

	e.Use(middleware.Gzip())

	e.Logger.Fatal(e.Start(":3003"))
}
