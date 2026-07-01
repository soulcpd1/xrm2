package xray

import (
	"encoding/json"
	"os"

	"github.com/soulcpd1/xrm2/internal/backup"
	"github.com/soulcpd1/xrm2/internal/system"
)

func Apply(mut func(*Config) error) error {

	if _, err := backup.Create(); err != nil {
		return err
	}

	raw, err := os.ReadFile(ConfigPath)
	if err != nil {
		return err
	}

	var cfg Config

	if err := json.Unmarshal(raw, &cfg); err != nil {
		return err
	}

	if err := mut(&cfg); err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
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
