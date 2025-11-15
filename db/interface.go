package db

type DBSource interface {
	Open(conn string)
	Execute(path string, params any) ([]byte, error)
}
