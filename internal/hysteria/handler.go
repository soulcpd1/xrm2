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

	category := h.Core.Classifier.Resolve(req)

	route, err := h.Core.Dispatcher.Resolve(
		req.SessionID,
		category,
	)
	if err != nil {
		return err
	}

	return h.Core.Stream.Bridge(
		context.Background(),
		route.Outbound,
		req.Conn,
		req.Network,
		req.Address,
	)
}
