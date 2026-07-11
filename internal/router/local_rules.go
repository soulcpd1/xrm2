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
                Value:  "domain:chatgpt.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:openai.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:oaistatic.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:oaiusercontent.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:claude.ai",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:anthropic.com",
                Action: VPS,
        },
}
