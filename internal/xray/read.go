package xray

import (
	"os"

	"github.com/tidwall/gjson"
)

const ConfigPath = "/usr/local/etc/xray/config.json"

func Read() (gjson.Result, error) {

	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return gjson.Result{}, err
	}

	return gjson.ParseBytes(data), nil
}
