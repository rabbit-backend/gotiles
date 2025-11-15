package core

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rabbit-backend/go-tiles/db"
)

func NewTileController(connections map[string]db.DBSource) func(c echo.Context) error {
	return func(c echo.Context) error {
		_x := c.Param("x")
		_y := c.Param("y")
		_z := c.Param("z")

		x, _ := strconv.Atoi(_x)
		y, _ := strconv.Atoi(_y)
		z, _ := strconv.Atoi(_z)

		tileName := c.Param("tile")
		source := c.Param("source")

		db := connections[source]

		params := map[string]any{
			"_x":     x,
			"_y":     y,
			"_z":     z,
			"_layer": tileName,
		}

		for key, value := range c.QueryParams() {
			params[key] = value[0]
		}

		tileQueryPath := path.Join("tiles", "db", source, fmt.Sprintf("%s.sql", tileName))
		data, err := db.Execute(c, tileQueryPath, params)

		if err != nil {
			log.Println("[x] error:", err)
			return c.Blob(http.StatusInternalServerError, "application/x-protobuf", []byte(""))
		}

		return c.Blob(http.StatusOK, "application/x-protobuf", data)
	}
}
