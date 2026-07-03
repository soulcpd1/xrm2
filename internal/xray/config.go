package xray

type Config struct {
	Inbounds []Inbound `json:"inbounds"`
}

type Inbound struct {
	Tag            string         `json:"tag"`
	Port           int            `json:"port"`
	StreamSettings StreamSettings `json:"streamSettings"`
	Settings       Settings       `json:"settings"`
}

type Settings struct {
	Clients []Client `json:"clients"`
}

type Client struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
	Flow  string `json:"flow,omitempty"`
}

type StreamSettings struct {
	Security        string          `json:"security"`
	RealitySettings RealitySettings `json:"realitySettings"`
}

type RealitySettings struct {
	ServerName string `json:"serverName"`
}
