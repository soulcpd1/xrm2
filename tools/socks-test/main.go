package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:2080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Greeting
	_, _ = conn.Write([]byte{
		0x05,
		0x01,
		0x00,
	})

	reply := make([]byte, 2)

	if _, err := io.ReadFull(conn, reply); err != nil {
		panic(err)
	}

	fmt.Printf("Handshake: % x\n", reply)

	// CONNECT github.com:443
	req := []byte{
		0x05, // VER
		0x01, // CONNECT
		0x00, // RSV
		0x03, // DOMAIN
		0x0a, // len("github.com")
		'g', 'i', 't', 'h', 'u', 'b', '.', 'c', 'o', 'm',
		0x01, 0xbb, // 443
	}

	_, _ = conn.Write(req)

	fmt.Println("CONNECT sent.")

	select {}
}
