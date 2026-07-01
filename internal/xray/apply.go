package xray

import (
	"os"

	"github.com/soulcpd1/xrm2/internal/backup"
	"github.com/soulcpd1/xrm2/internal/system"
)

func Apply(fn func([]byte) ([]byte, error)) error {

	if _, err := backup.Create(); err != nil {
		return err
	}

	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return err
	}

	data, err = fn(data)
	if err != nil {
		return err
	}

	tmp := ConfigPath + ".tmp"

	if err := Write(tmp, data); err != nil {
		return err
	}

	if err := Validate(tmp); err != nil {
		_ = os.Remove(tmp)
		return err
	}

	if err := os.Rename(tmp, ConfigPath); err != nil {
		return err
	}

	return system.ReloadXray()
}
