package hysteria

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/common"
)

var parser = NewParser()

func BuildRequest(conn net.Conn) *common.Request {

	target, err := parser.Parse(conn)
	if err != nil {
		return nil
	}

	return &common.Request{
		Conn: conn,

		Network: target.Network,
		Address: target.Address,
	}
}
