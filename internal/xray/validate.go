package xray

import (
	"fmt"
	"os/exec"
	"strings"
)

func Validate(path string) error {

	cmd := exec.Command(
		"/usr/local/bin/xray",
		"run",
		"-test",
		"-config",
		path,
	)

	out, err := cmd.CombinedOutput()

	output := string(out)

	fmt.Println("========== XRAY VALIDATE OUTPUT ==========")
	fmt.Println(output)
	fmt.Println("==========================================")

	// ❗ Xray returns non-zero even on success
	if strings.Contains(output, "Configuration OK") {
		return nil
	}

	if err != nil {
		return fmt.Errorf("xray validate failed")
	}

	return nil
}
