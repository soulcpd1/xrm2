package xray

func FindRealityInbound(cfg *Config) *Inbound {
	for i := range cfg.Inbounds {
		if cfg.Inbounds[i].StreamSettings.Security == "reality" {
			return &cfg.Inbounds[i]
		}
	}
	return nil
}
