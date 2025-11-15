package db

import engine "github.com/rabbit-backend/template"

var DB_SOURCES = map[string](func(e *engine.Engine) DBSource){
	"postgres": func(e *engine.Engine) DBSource { return &PGSource{engine: e} },
	"mysql":    func(e *engine.Engine) DBSource { return &MySQLSource{engine: e} },
}
