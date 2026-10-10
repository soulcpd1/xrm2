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
}
