package doctor

import (
	"os/exec"
	"strings"
)

type ServiceChecker struct{}

func (ServiceChecker) Run() Result {

	out, err := exec.Command("systemctl", "is-active", "xray").Output()

	if err != nil {
		return Result{
			Name:   "Xray Service",
			OK:     false,
			Detail: strings.TrimSpace(string(out)),
		}
	}

	status := strings.TrimSpace(string(out))

	return Result{
		Name:   "Xray Service",
		OK:     status == "active",
		Detail: status,
	}
}
