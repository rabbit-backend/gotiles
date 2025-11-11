package db

var DB_SOURCES = map[string](func() DBSource){
	"postgres": func() DBSource { return &PGSource{} },
}
