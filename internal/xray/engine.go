package xray

import (
	"encoding/json"
	"os"
)

type Engine struct {
	configPath string
}

func NewEngine() *Engine {
	return &Engine{
		configPath: "/usr/local/etc/xray/config.json",
	}
}

func (e *Engine) Load() (*Config, error) {
	data, err := os.ReadFile(e.configPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	return &cfg, err
}

func (e *Engine) Save(cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(e.configPath, data, 0644)
}

func ApplyWithEngine(fn func(cfg *Config) error) error {
	e := NewEngine()

	cfg, err := e.Load()
	if err != nil {
		return err
	}

	if err := fn(cfg); err != nil {
		return err
	}

	return e.Save(cfg)
}
