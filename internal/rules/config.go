package rules

type Config struct {
	Rules []ConfigRule `yaml:"rules"`
}

type ConfigRule struct {
	Type   string `yaml:"type"`
	Value  string `yaml:"value"`
	Policy string `yaml:"policy"`
}
