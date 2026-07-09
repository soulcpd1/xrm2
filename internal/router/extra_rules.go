package router

var ExtraRules = []Rule{

	// Telegram
	{Type: CIDRRule, Value: "91.105.192.0/23", Action: VPS},
	{Type: CIDRRule, Value: "91.108.4.0/22", Action: VPS},
	{Type: CIDRRule, Value: "91.108.8.0/22", Action: VPS},
	{Type: CIDRRule, Value: "91.108.12.0/22", Action: VPS},
	{Type: CIDRRule, Value: "91.108.16.0/22", Action: VPS},
	{Type: CIDRRule, Value: "91.108.56.0/22", Action: VPS},
	{Type: CIDRRule, Value: "149.154.160.0/20", Action: VPS},

	// Meta
	{Type: CIDRRule, Value: "31.13.64.0/18", Action: VPS},
	{Type: CIDRRule, Value: "66.220.144.0/20", Action: VPS},
	{Type: CIDRRule, Value: "69.63.176.0/20", Action: VPS},
	{Type: CIDRRule, Value: "69.171.224.0/19", Action: VPS},
	{Type: CIDRRule, Value: "157.240.0.0/16", Action: VPS},
	{Type: CIDRRule, Value: "129.134.0.0/16", Action: VPS},
	{Type: CIDRRule, Value: "185.60.216.0/22", Action: VPS},
}
