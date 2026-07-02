package tuic

import (
	"encoding/json"
	"os"
)

func DefaultConfig() *Config {
	return &Config{
		Server:      ":443",
		Certificate: "/etc/tuic/server.crt",
		PrivateKey:  "/etc/tuic/server.key",
		Congestion:  "bbr",
		ALPN:        []string{"h3"},
		Users:       []User{},
	}
}

func Save(cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigPath, data, 0644)
}
