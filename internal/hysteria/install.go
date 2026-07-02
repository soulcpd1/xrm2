package hysteria

import (
	"fmt"
	"runtime"
)

func PrintInstallPlan() {

	fmt.Println("========================================")
	fmt.Println("        Hysteria2 Installer")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf("%-16s %s\n", "OS:", runtime.GOOS)
	fmt.Printf("%-16s %s\n", "Arch:", runtime.GOARCH)
	fmt.Printf("%-16s %s\n", "Version:", Version)
	fmt.Println()

	fmt.Printf("%-16s %s\n", "Binary:", BinaryPath)
	fmt.Printf("%-16s %s\n", "Config:", ConfigPath)
	fmt.Printf("%-16s %s\n", "Service:", ServicePath)
	fmt.Printf("%-16s %s\n", "Data:", DataDir)
	fmt.Println()

	fmt.Println("Status: READY")
}
