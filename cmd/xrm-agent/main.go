package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/soulcpd1/xrm2/internal/agent"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	a := agent.New("192.168.100.1:7000")

	log.Println("XRM Agent listening on 192.168.100.1:7000")

	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
