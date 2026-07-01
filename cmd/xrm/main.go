package main

import (
	"os"

	"github.com/soulcpd1/xrm2/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
