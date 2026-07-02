package hysteria

type Config struct {
	Listen       string       `yaml:"listen"`
	TLS          TLS          `yaml:"tls"`
	Auth         Auth         `yaml:"auth"`
	Masquerade   Masquerade   `yaml:"masquerade"`
	Bandwidth    Bandwidth    `yaml:"bandwidth"`
	QUIC         QUIC         `yaml:"quic"`
	IgnoreClient bool         `yaml:"ignoreClientBandwidth,omitempty"`
}

type TLS struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

type Auth struct {
	Type     string `yaml:"type"`
	Password string `yaml:"password"`
}

type Masquerade struct {
	Type string `yaml:"type"`
	Proxy Proxy  `yaml:"proxy"`
}

type Proxy struct {
	URL string `yaml:"url"`
}

type Bandwidth struct {
	Up   string `yaml:"up"`
	Down string `yaml:"down"`
}

type QUIC struct {
	InitStreamReceiveWindow     uint64 `yaml:"initStreamReceiveWindow"`
	MaxStreamReceiveWindow      uint64 `yaml:"maxStreamReceiveWindow"`
	InitConnectionReceiveWindow uint64 `yaml:"initConnReceiveWindow"`
	MaxConnectionReceiveWindow  uint64 `yaml:"maxConnReceiveWindow"`
	MaxIdleTimeout              string `yaml:"maxIdleTimeout"`
}
