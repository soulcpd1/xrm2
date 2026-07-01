package doctor

import (
	"os/exec"
	"strings"
)

type XrayChecker struct{}

func (XrayChecker) Run() Result {

	cmd := exec.Command("/usr/local/bin/xray", "version")

	out, err := cmd.Output()
	if err != nil {
		return Result{
			Name:   "Xray",
			OK:     false,
			Detail: err.Error(),
		}
	}

	line := strings.Split(strings.TrimSpace(string(out)), "\n")[0]

	return Result{
		Name:   "Xray",
		OK:     true,
		Detail: line,
	}
}
