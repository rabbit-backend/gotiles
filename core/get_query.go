package core

import "os"

func GetQuery(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}