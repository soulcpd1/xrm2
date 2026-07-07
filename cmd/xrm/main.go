package main

import (
	"log"

	"github.com/soulcpd1/xrm2/internal/inbound"
	"github.com/soulcpd1/xrm2/internal/runtime"
)

func main() {

	rt := runtime.New()

	log.Println("[XRM] STARTING REAL DATAPATH")

	go rt.Start()

	in := inbound.New(rt)

	log.Println("[XRM] inbound listening on 127.0.0.1:22345")

	log.Fatal(in.Listen("127.0.0.1:22345"))
}
