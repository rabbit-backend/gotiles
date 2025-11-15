package db

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	engine "github.com/rabbit-backend/template"
)

type PGSource struct {
	db     *sql.DB
	engine *engine.Engine
}

func (s *PGSource) Open(conn string) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}

	s.db = db
}

func (s *PGSource) Execute(_ echo.Context, queryPath string, params any) ([]byte, error) {
	var buf []byte

	query, args, err := s.engine.Execute(queryPath, params)
	if err != nil {
		return nil, err
	}

	if err := s.db.QueryRow(query, args...).Scan(&buf); err != nil {
		return nil, err
	}

	return buf, nil
}
