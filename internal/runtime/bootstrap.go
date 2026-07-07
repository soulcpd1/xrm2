package runtime

import (
	"log"
	"net"

	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/router"
	"github.com/soulcpd1/xrm2/internal/socks"
)

type Runtime struct {
	dispatcher *core.Dispatcher
	forwarder  *core.Forwarder
	socks      *socks.Server
}

func New() *Runtime {

	rt := router.New()

	direct := outbound.NewDirectDriver()
	vps := outbound.NewVPSDriver()

	ob := outbound.NewManager(direct, vps)

	fw := core.NewForwarder(ob)

	dp := core.NewDispatcher(fw, rt)

	return &Runtime{
		dispatcher: dp,
		forwarder:  fw,
		socks:      socks.New(dp),
	}
}

func (r *Runtime) Handle(conn net.Conn) {

	log.Println("[RUNTIME] REAL CONNECTION RECEIVED")

	req := core.ToRequest(conn)

	_, err := r.dispatcher.Handle(req)
	if err != nil {
		log.Println("[RUNTIME] dispatch error:", err)
	}
}

func (r *Runtime) Start() {

	log.Println("[RUNTIME] READY")

	go r.socks.Serve("127.0.0.1:2080")

	select {}
}

func (r *Runtime) Forwarder() *core.Forwarder {
	return r.forwarder
}
