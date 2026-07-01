package xray

import (
	"fmt"
	"os"
	"path/filepath"
)

func Write(path string, data []byte) error {

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("mkdir failed: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write failed: %w (path=%s)", err, path)
	}

	return nil
}
