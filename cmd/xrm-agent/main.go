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

	a := agent.New("0.0.0.0:7000")

	log.Println("XRM Agent listening on :7000")

	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
