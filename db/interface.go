package db

import "github.com/labstack/echo/v4"

type DBSource interface {
	Open(conn string)
	Execute(c echo.Context, path string, params any) ([]byte, error)
}
