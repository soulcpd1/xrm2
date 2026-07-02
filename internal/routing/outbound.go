package routing

type OutboundType string

const (
	Direct OutboundType = "direct"
	Nebula OutboundType = "nebula"
)

type Outbound struct {
	Name      string
	Type      OutboundType
	Interface string
	Gateway   string
	Metric    int
	Enabled   bool
}
