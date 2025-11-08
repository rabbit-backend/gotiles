package core

import (
	engine "github.com/rabbit-backend/template"
)

func GetQuery(e *engine.Engine, path string, params any) (string, []any, error) {
	query, args, err := e.Execute(path, params)
	return query, args, err
}
