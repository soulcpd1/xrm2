package xray

import (
	"fmt"
	"os/exec"
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
	if err != nil {
		return fmt.Errorf("xray validate failed:\n%s", string(out))
	}

	return nil
}
