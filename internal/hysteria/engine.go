package hysteria

import (
	"os"

	"gopkg.in/yaml.v3"
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

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (e *Engine) Save(cfg *Config) error {

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(e.configPath, data, 0644)
}

func (e *Engine) Exists() bool {
	_, err := os.Stat(e.configPath)
	return err == nil
}

func (e *Engine) ConfigPath() string {
	return e.configPath
}
