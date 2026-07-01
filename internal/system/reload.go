package system

import "os/exec"

func ReloadXray() error {
	return exec.Command(
		"systemctl",
		"reload",
		"xray",
	).Run()
}
