package db

import engine "github.com/rabbit-backend/template"

var DB_SOURCES = map[string](func(e *engine.Engine) DBSource){
	"postgres": func(e *engine.Engine) DBSource { return &PGSource{engine: e} },
	"memsql":   func(e *engine.Engine) DBSource { return &MemSQLSource{engine: e} },
}
