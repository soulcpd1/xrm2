package xray

func (c *Config) RealityInbound() *Inbound {
	for i := range c.Inbounds {
		if c.Inbounds[i].StreamSettings.Security == "reality" {
			return &c.Inbounds[i]
		}
	}
	return nil
}
