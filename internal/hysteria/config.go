package hysteria

func DefaultConfig() *Config {

	return &Config{
		Listen: ":443",

		TLS: TLS{
			Cert: "/etc/hysteria/server.crt",
			Key:  "/etc/hysteria/server.key",
		},

		Auth: Auth{
			Type:     "password",
			Password: "CHANGE_ME",
		},

		Masquerade: Masquerade{
			Type: "proxy",
			Proxy: Proxy{
				URL: "https://www.cloudflare.com",
			},
		},

		Bandwidth: Bandwidth{
			Up:   "1 gbps",
			Down: "1 gbps",
		},

		QUIC: QUIC{
			InitStreamReceiveWindow:     8388608,
			MaxStreamReceiveWindow:      8388608,
			InitConnectionReceiveWindow: 20971520,
			MaxConnectionReceiveWindow:  20971520,
			MaxIdleTimeout:              "30s",
		},

		IgnoreClient: true,
	}
}
