package hysteria

import (
	"context"

	"github.com/soulcpd1/xrm2/internal/common"
	"github.com/soulcpd1/xrm2/internal/core"
)

type Handler struct {
	Core *core.Core
}

func NewHandler(c *core.Core) *Handler {
	return &Handler{Core: c}
}

func (h *Handler) Resolve(req *common.Request) error {

	// FIX: Resolve is void, just mutate req or ignore return
	h.Core.Classifier.Resolve(req)

	req.Category = req.Category

	return h.Core.Dispatcher.Resolve(req)
}

func (h *Handler) Forward(req *common.Request) error {

	return h.Core.Stream.BridgeConn(
		context.Background(),
		"tcp",
		req.Conn,
		req.Network,
		req.Address,
	)
}
