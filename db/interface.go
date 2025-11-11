package db

type DBSource interface {
	Open(conn string)
	Execute(query string, args ...any) ([]byte, error)
}
