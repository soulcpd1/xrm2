package xray

import "github.com/tidwall/gjson"

func RealityInbound(cfg gjson.Result) gjson.Result {

	for _, inbound := range cfg.Get("inbounds").Array() {

		if inbound.Get("streamSettings.security").String() == "reality" {
			return inbound
		}

	}

	return gjson.Result{}
}

func FindClientByEmail(cfg gjson.Result, email string) *RealityClient {

	for _, c := range RealityClients(cfg) {

		if c.Email == email {
			client := c
			return &client
		}

	}

	return nil
}

func FindClientByUUID(cfg gjson.Result, uuid string) *RealityClient {

	for _, c := range RealityClients(cfg) {

		if c.ID == uuid {
			client := c
			return &client
		}

	}

	return nil
}
