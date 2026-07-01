package xray

import (
	"strconv"

	"github.com/tidwall/gjson"
)

func RealityInboundIndex(cfg gjson.Result) int {

	for i, inbound := range cfg.Get("inbounds").Array() {

		if inbound.Get("streamSettings.security").String() == "reality" {
			return i
		}

	}

	return -1
}

func RealityClientsPath(cfg gjson.Result) string {

	idx := RealityInboundIndex(cfg)

	if idx < 0 {
		return ""
	}

	return "inbounds." + strconv.Itoa(idx) + ".settings.clients"
}
