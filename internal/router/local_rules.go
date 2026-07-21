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
                Value:  "domain:youtube.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:youtu.be",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:googlevideo.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:ytimg.com",
                Action: VPS,
        },
        {
                Type:   DomainRule,
                Value:  "domain:youtubei.googleapis.com",
                Action: VPS,
        },
}
