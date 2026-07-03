package tuic

type Config struct {
	Server      string   `json:"server"`
	Users       []User   `json:"users"`
	Certificate string   `json:"certificate"`
	PrivateKey  string   `json:"private_key"`
	Congestion  string   `json:"congestion_control"`
	ALPN        []string `json:"alpn"`
}

type User struct {
	UUID     string `json:"uuid"`
	Password string `json:"password"`
}
