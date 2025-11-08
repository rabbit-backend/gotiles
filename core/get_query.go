package core

import (
	engine "github.com/rabbit-backend/template"
)

func GetQuery(e *engine.Engine, path string, params any) (string, []any) {
	query, args := e.Execute(path, params)
	return query, args
}
