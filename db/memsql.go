package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLSource struct {
	db *sql.DB
}

func (s *MySQLSource) Open(conn string) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}

	s.db = db
}

func (s *MySQLSource) Execute(query string, args ...any) ([]byte, error) {
	var buf []byte
	if err := s.db.QueryRow(query, args...).Scan(&buf); err != nil {
		return nil, err
	}

	return buf, nil
}
