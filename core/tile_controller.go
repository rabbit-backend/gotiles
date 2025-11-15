package core

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/rabbit-backend/go-tiles/db"
)

func NewTileController(connections map[string]db.DBSource) func(c echo.Context) error {
	return func(c echo.Context) error {
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
		data, err := db.Execute(tileQueryPath, params)
		if err != nil {
			log.Println("[x] error:", err)
			return c.Blob(http.StatusInternalServerError, "application/x-protobuf", []byte(""))
		}

		return c.Blob(http.StatusOK, "application/x-protobuf", data)
	}
}
