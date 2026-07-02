package xray

import "github.com/tidwall/gjson"

// Returns gjson path to reality clients array
func RealityClientsPath(cfg gjson.Result) string {
	// ищем inbound с reality
	inbounds := cfg.Get("inbounds").Array()

	for i := range inbounds {
		if inbounds[i].Get("streamSettings.security").String() == "reality" {
			return "inbounds." + itoa(i) + ".settings.clients"
		}
	}

	return ""
}

// простая конверсия int -> string без strconv (чтобы не раздувать зависимости)
func itoa(i int) string {
	if i == 0 {
		return "0"
	}

	digits := "0123456789"
	var b [20]byte
	pos := len(b)

	for i > 0 {
		pos--
		b[pos] = digits[i%10]
		i /= 10
	}

	return string(b[pos:])
}
