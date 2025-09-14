package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/tiles/:x/:y/:z", func(c echo.Context) error {
		x := c.Param("x")
		y := c.Param("y")
		z := c.Param("z")

		return c.String(http.StatusOK, fmt.Sprintf("%s %s %s", x, y, z))
	})

	e.Logger.Fatal(e.Start(":3003"))
}
