package doctor

import "fmt"

type HysteriaChecker struct{}

func NewHysteriaChecker() *HysteriaChecker {
	return &HysteriaChecker{}
}

func (h *HysteriaChecker) Check() error {
	// TODO: real check later
	fmt.Println("Hysteria: OK")
	return nil
}
