package backup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	ConfigFile = "/usr/local/etc/xray/config.json"
	BackupDir  = "/var/lib/xrm2/backups"
)

func Create() (string, error) {

	if err := os.MkdirAll(BackupDir, 0755); err != nil {
		return "", err
	}

	name := time.Now().Format("20060102-150405") + ".json"

	dst := filepath.Join(BackupDir, name)

	srcFile, err := os.Open(ConfigFile)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", err
	}

	return fmt.Sprintf("Backup created: %s", dst), nil
}
