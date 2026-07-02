package hysteria

type Config struct {
	Listen string `yaml:"listen"`

	TLS TLS `yaml:"tls"`

	Auth Auth `yaml:"auth"`

	Transport Transport `yaml:"transport"`

	Masquerade Masquerade `yaml:"masquerade"`
}

type TLS struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

type Auth struct {
	Type     string `yaml:"type"`
	Password string `yaml:"password"`
}

type Transport struct {
	Type string `yaml:"type"`
}

type Masquerade struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}
