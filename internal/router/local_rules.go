package router

var LocalRules = []Rule{

        {
                Type:   DomainRule,
                Value:  "domain:max.ru",
                Action: Block,
        },

        {
                Type:   DomainRule,
                Value:  "domain:oneme.ru",
                Action: Block,
        },

        {
                Type:   DomainRule,
                Value:  "domain:whatsapp.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:whatsapp.net",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:signal.org",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:whispersystems.org",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:fbcdn.net",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:cdninstagram.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:fbsbx.com",
                Action: VPS,
        },
}
