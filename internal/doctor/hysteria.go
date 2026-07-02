package doctor

import (
	"os"
)

type HysteriaChecker struct{}

func (HysteriaChecker) Run() Result {

	if _, err := os.Stat("/usr/local/bin/hysteria"); err == nil {
		return Result{
			Name:   "Hysteria Binary",
			OK:     true,
			Detail: "installed",
		}
	}

	return Result{
		Name:   "Hysteria Binary",
		OK:     false,
		Detail: "not installed",
	}
}
