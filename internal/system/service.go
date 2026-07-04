package system

import "os/exec"

func DaemonReload() error {
	return exec.Command("systemctl", "daemon-reload").Run()
}

func Enable(name string) error {
	return exec.Command("systemctl", "enable", name).Run()
}

func Disable(name string) error {
	return exec.Command("systemctl", "disable", name).Run()
}

func Start(name string) error {
	return exec.Command("systemctl", "start", name).Run()
}

func Stop(name string) error {
	return exec.Command("systemctl", "stop", name).Run()
}

func Restart(name string) error {
	return exec.Command("systemctl", "restart", name).Run()
}

func Reload(name string) error {
	return exec.Command("systemctl", "reload", name).Run()
}
