package xray

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/soulcpd1/xrm2/internal/backup"
	"github.com/soulcpd1/xrm2/internal/system"
)

type Engine struct {
	configPath string
}

func NewEngine() *Engine {
	return &Engine{
		configPath: ConfigPath,
	}
}

func (e *Engine) Load() (*Config, error) {

	data, err := os.ReadFile(e.configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (e *Engine) Save(cfg *Config) error {

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	tmp := filepath.Join(
		filepath.Dir(e.configPath),
		"config.tmp.json",
	)

	if err := Write(tmp, data); err != nil {
		return err
	}

	if err := Validate(tmp); err != nil {
		_ = os.Remove(tmp)
		return err
	}

	if _, err := backup.Create(); err != nil {
		_ = os.Remove(tmp)
		return err
	}

	if err := os.Rename(tmp, e.configPath); err != nil {
		return err
	}

	return system.ReloadXray()
}

func ApplyWithEngine(fn func(*Config) error) error {

	engine := NewEngine()

	cfg, err := engine.Load()
	if err != nil {
		return err
	}

	if err := fn(cfg); err != nil {
		return err
	}

	return engine.Save(cfg)
}
