package tuic

import (
	"fmt"
	"runtime"
)

func Install() error {

	fmt.Println("========================================")
	fmt.Println("           TUIC Installer")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Printf("OS:              %s\n", runtime.GOOS)
	fmt.Printf("Arch:            %s\n", runtime.GOARCH)
	fmt.Printf("Version:         %s\n", Version)
	fmt.Println()

	fmt.Printf("Binary:          %s\n", BinaryPath)
	fmt.Printf("Config:          %s\n", ConfigPath)
	fmt.Printf("Service:         %s\n", ServicePath)
	fmt.Printf("Data:            %s\n", DataDir)
	fmt.Println()

	fmt.Println("Status: READY")

	return nil
}
