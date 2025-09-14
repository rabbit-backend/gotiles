package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		// x := c.Param("x")
		// y := c.Param("y")
		// z := c.Param("z")

		return c.Blob(http.StatusOK, "text/html", []byte(""))
	})

	e.Logger.Fatal(e.Start(":3003"))
}
