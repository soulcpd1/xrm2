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

	// Слушаем на двух адресах: Nebula и AmneziaWG
	addrs := []string{
		"192.168.100.1:7000", // Nebula (основной)
		"10.9.9.1:7000",
		"10.8.0.2:7000",
		"10.99.0.2:7000",      // AmneziaWG (резерв)
		"10.147.17.1:7000",    // ZeroTier
	}

	a := agent.New(addrs)

	log.Printf("XRM Agent listening on: %v", addrs)

	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
