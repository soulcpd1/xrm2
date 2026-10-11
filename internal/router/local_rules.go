package router

var LocalRules = []Rule{

	// Block
	{Type: DomainRule, Value: "domain:max.ru", Action: Block},
	{Type: DomainRule, Value: "domain:oneme.ru", Action: Block},

	// YouTube
	{Type: DomainRule, Value: "domain:youtube.com", Action: VPS},
	{Type: DomainRule, Value: "domain:youtu.be", Action: VPS},
	{Type: DomainRule, Value: "domain:googlevideo.com", Action: VPS},
	{Type: DomainRule, Value: "domain:ytimg.com", Action: VPS},
	{Type: DomainRule, Value: "domain:youtubei.googleapis.com", Action: VPS},

	// Google
	{Type: DomainRule, Value: "domain:googleusercontent.com", Action: VPS},
	{Type: DomainRule, Value: "domain:googleapis.com", Action: VPS},
	{Type: DomainRule, Value: "domain:gstatic.com", Action: VPS},
	{Type: DomainRule, Value: "domain:ggpht.com", Action: VPS},
	{Type: DomainRule, Value: "domain:googlesyndication.com", Action: VPS},
	{Type: DomainRule, Value: "domain:googleadservices.com", Action: VPS},
	{Type: DomainRule, Value: "domain:doubleclick.net", Action: VPS},
	{Type: DomainRule, Value: "domain:google.com", Action: VPS},
	{Type: DomainRule, Value: "domain:goo.gl", Action: VPS},
	{Type: DomainRule, Value: "domain:withgoogle.com", Action: VPS},

	// HomeWhiz (Beko/Grundig/Arcelik/Blomberg)
	{Type: DomainRule, Value: "domain:homewhiz.com", Action: VPS},
	{Type: DomainRule, Value: "domain:api.homewhiz.com", Action: VPS},
	{Type: DomainRule, Value: "domain:aws.homewhiz.com", Action: VPS},
	{Type: DomainRule, Value: "domain:beko.com", Action: VPS},
	{Type: DomainRule, Value: "domain:api.beko.com", Action: VPS},
	{Type: DomainRule, Value: "domain:grundig.com", Action: VPS},
	{Type: DomainRule, Value: "domain:arcelik.com", Action: VPS},
	{Type: DomainRule, Value: "domain:blomberg.com", Action: VPS},
	{Type: DomainRule, Value: "domain:homewhiz.net", Action: VPS},
	{Type: DomainRule, Value: "domain:homewhiz.io", Action: VPS},

	// HomeWhiz (Beko/Grundig/Arcelik/Blomberg)
	{Type: DomainRule, Value: "domain:homewhiz.com", Action: VPS},
	{Type: DomainRule, Value: "domain:api.homewhiz.com", Action: VPS},
	{Type: DomainRule, Value: "domain:aws.homewhiz.com", Action: VPS},
	{Type: DomainRule, Value: "domain:homewhiz.net", Action: VPS},
	{Type: DomainRule, Value: "domain:homewhiz.io", Action: VPS},
	{Type: DomainRule, Value: "domain:beko.com", Action: VPS},
	{Type: DomainRule, Value: "domain:api.beko.com", Action: VPS},
	{Type: DomainRule, Value: "domain:grundig.com", Action: VPS},
	{Type: DomainRule, Value: "domain:arcelik.com", Action: VPS},
	{Type: DomainRule, Value: "domain:blomberg.com", Action: VPS},

	// HomeWhiz / Arcelik IoT
	{Type: DomainRule, Value: "domain:arcelikiot.com", Action: VPS},
	{Type: DomainRule, Value: "domain:api.arcelikiot.com", Action: VPS},
	{Type: DomainRule, Value: "domain:globalconfig.arcelikiot.com", Action: VPS},
	{Type: DomainRule, Value: "domain:homewhiz.com", Action: VPS},
	{Type: DomainRule, Value: "domain:homewhiz.net", Action: VPS},
	{Type: DomainRule, Value: "domain:beko.com", Action: VPS},
	{Type: DomainRule, Value: "domain:grundig.com", Action: VPS},
	{Type: DomainRule, Value: "domain:arcelik.com", Action: VPS},
	{Type: DomainRule, Value: "domain:blomberg.com", Action: VPS},
	{Type: DomainRule, Value: "domain:execute-api.us-east-1.amazonaws.com", Action: VPS},
	{Type: DomainRule, Value: "domain:execute-api.eu-west-1.amazonaws.com", Action: VPS},
}
