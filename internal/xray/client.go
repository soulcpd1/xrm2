package xray

import "github.com/tidwall/gjson"

type RealityClient struct {
	ID    string
	Email string
}

func RealityClients(cfg gjson.Result) []RealityClient {

	var list []RealityClient

	inbounds := cfg.Get("inbounds").Array()

	for _, inbound := range inbounds {

		if inbound.Get("streamSettings.security").String() != "reality" {
			continue
		}

		clients := inbound.Get("settings.clients").Array()

		for _, c := range clients {

			list = append(list, RealityClient{
				ID:    c.Get("id").String(),
				Email: c.Get("email").String(),
			})

		}

		break
	}

	return list
}
