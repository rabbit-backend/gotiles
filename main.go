package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rabbit-backend/go-tiles/db"
)

var DB *sql.DB

func init() {
	var err error

	// Connect to the Postgres Database
	if DB, err = db.ConnectDB(); err != nil {
		log.Fatalln("[x] error connecting db:", err)
	}
}

func main() {
	e := echo.New()

	e.GET("/tiles/:tile/:x/:y/:z", func(c echo.Context) error {
		x := c.Param("x")
		y := c.Param("y")
		z := c.Param("z")

		return c.String(http.StatusOK, fmt.Sprintf("%s %s %s", x, y, z))
	})

	e.Logger.Fatal(e.Start(":3003"))
}
