package xray

import (
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

	return cmd.Run()
}
