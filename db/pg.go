package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PGSource struct {
	db *sql.DB
}

func (s *PGSource) Open(conn string) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	s.db = db
}

func (s *PGSource) Execute(query string, args ...any) ([]byte, error) {
	var buf []byte
	if err := s.db.QueryRow(query, args...).Scan(&buf); err != nil {
		return nil, err
	}

	return buf, nil
}
