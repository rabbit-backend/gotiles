package main

import (
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
	e.GET("/tiles/:source/:tile/:x/:y/:z", core.NewTileController(connections, sqlEngine))

	e.Use(middleware.Gzip())

	e.Logger.Fatal(e.Start(":3003"))
}
