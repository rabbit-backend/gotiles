package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rabbit-backend/go-tiles/core"
	engine "github.com/rabbit-backend/template"
)

func init() {
	godotenv.Load()
}

func main() {
	config := core.GetConfig()
	connections := config.GetConnections()
	sqlEngine := engine.NewEngineWithPlaceHolder(engine.NewPostgresPlaceHolder())

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

		params := map[string]any{
			"_x": x,
			"_y": y,
			"_z": z,
		}

		for key, value := range c.QueryParams() {
			params[key] = value[0]
		}

		tileQueryPath := path.Join("tiles", "db", source, fmt.Sprintf("%s.sql", tileName))
		query, args, err := core.GetQuery(sqlEngine, tileQueryPath, params)
		if err != nil {
			log.Println("[x] error:", err)
			return c.Blob(http.StatusInternalServerError, "application/x-protobuf", []byte(""))
		}

		data, err := db.Execute(query, args...)
		if err != nil {
			log.Println("[x] error:", err)
			return c.Blob(http.StatusInternalServerError, "application/x-protobuf", []byte(""))
		}

		return c.Blob(http.StatusOK, "application/x-protobuf", data)
	})

	e.Use(middleware.Gzip())

	e.Logger.Fatal(e.Start(":3003"))
}
