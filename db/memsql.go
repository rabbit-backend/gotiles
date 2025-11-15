package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	engine "github.com/rabbit-backend/template"
)

type MySQLSource struct {
	db     *sql.DB
	engine *engine.Engine
}

func (s *MySQLSource) Open(conn string) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}

	s.db = db
}

func (s *MySQLSource) Execute(queryPath string, params any) ([]byte, error) {
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
