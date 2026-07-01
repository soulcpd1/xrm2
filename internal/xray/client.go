package xray

import "github.com/tidwall/gjson"

type RealityClient struct {
	ID    string
	Email string
}

func RealityPort(cfg gjson.Result) int {

	for _, inbound := range cfg.Get("inbounds").Array() {

		if inbound.Get("streamSettings.security").String() == "reality" {
			return int(inbound.Get("port").Int())
		}

	}

	return 0
}

func RealityClients(cfg gjson.Result) []RealityClient {

	var list []RealityClient

	for _, inbound := range cfg.Get("inbounds").Array() {

		if inbound.Get("streamSettings.security").String() != "reality" {
			continue
		}

		for _, c := range inbound.Get("settings.clients").Array() {

			list = append(list, RealityClient{
				ID:    c.Get("id").String(),
				Email: c.Get("email").String(),
			})

		}

		break
	}

	return list
}
