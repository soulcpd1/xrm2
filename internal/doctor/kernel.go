package doctor

import (
	"os/exec"
	"strings"
)

type KernelChecker struct{}

func (KernelChecker) Run() Result {

	out, err := exec.Command("uname", "-r").Output()

	if err != nil {
		return Result{
			Name: "Kernel",
			OK:   false,
			Detail: err.Error(),
		}
	}

	return Result{
		Name:   "Kernel",
		OK:     true,
		Detail: strings.TrimSpace(string(out)),
	}
}
