package hysteria

import (
	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/policy"
)

type Handler struct {
	Core *core.Core
}

func NewHandler(c *core.Core) *Handler {
	return &Handler{
		Core: c,
	}
}

func (h *Handler) Resolve(
	sessionID string,
	category policy.Category,
) error {

	route, err := h.Core.Dispatcher.Resolve(
		sessionID,
		category,
	)

	if err != nil {
		return err
	}

	_, err = h.Core.Outbound.Driver(route.Outbound)

	if err != nil {
		return err
	}

	return nil
}
