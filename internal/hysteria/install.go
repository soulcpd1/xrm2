package hysteria

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Install() error {

	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return err
	}

	cfg := DefaultConfig()

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile(ConfigPath, data, 0644); err != nil {
		return err
	}

	fmt.Println("Config created:")
	fmt.Println(" ", ConfigPath)

	return nil
}
