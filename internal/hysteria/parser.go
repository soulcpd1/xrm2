package hysteria

import "net"

type Target struct {
	Network string
	Address string
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(conn net.Conn) (*Target, error) {

	// Пока заглушка.
	// Следующим коммитом здесь будет разбор
	// настоящего Hysteria handshake.

	return &Target{
		Network: "tcp",
		Address: "8.8.8.8:443",
	}, nil
}
