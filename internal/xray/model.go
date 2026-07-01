package xray

type Config struct {
	Inbounds  []Inbound  `json:"inbounds"`
	Outbounds []Outbound `json:"outbounds"`
}

type Inbound struct {
	Tag      string `json:"tag"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`

	Settings struct {
		Clients []Client `json:"clients"`
	} `json:"settings"`

	StreamSettings struct {
		Security string `json:"security"`

		RealitySettings struct {
			ServerNames []string `json:"serverNames"`
			ShortIDs    []string `json:"shortIds"`
			PublicKey   string   `json:"publicKey,omitempty"`
			PrivateKey  string   `json:"privateKey,omitempty"`
		} `json:"realitySettings"`
	} `json:"streamSettings"`
}

type Client struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
	Flow  string `json:"flow,omitempty"`
}

type Outbound struct {
	Tag      string `json:"tag"`
	Protocol string `json:"protocol"`

	Settings struct {
		VNext []struct {
			Address string `json:"address"`
			Port    int    `json:"port"`
		} `json:"vnext"`
	} `json:"settings"`
}

type ClientMutation struct {
	Email string
	Flow  string
}
