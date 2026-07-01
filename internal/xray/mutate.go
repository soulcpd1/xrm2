package xray

import (
	"github.com/tidwall/sjson"
)

func Set(data []byte, path string, value any) ([]byte, error) {
	return sjson.SetBytes(data, path, value)
}

func Delete(data []byte, path string) ([]byte, error) {
	return sjson.DeleteBytes(data, path)
}
