package router

type RuleType uint8

const (
        DomainRule RuleType = iota
        CIDRRule
)

type Rule struct {
        Type   RuleType
        Value  string
        Action Action
}

var DefaultRules = []Rule{

        // localhost

        {
                Type:   CIDRRule,
                Value:  "127.0.0.0/8",
                Action: Direct,
        },

        // Telegram DC

        {
                Type:   CIDRRule,
                Value:  "149.154.160.0/20",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "91.108.4.0/22",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "91.108.8.0/22",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "91.108.12.0/22",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "91.108.16.0/22",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "91.108.56.0/22",
                Action: VPS,
        },

        // Telegram CDN

        {
                Type:   CIDRRule,
                Value:  "91.105.192.0/23",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "91.105.194.0/23",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "91.105.196.0/23",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "91.105.198.0/23",
                Action: VPS,
        },
        {
                Type:   CIDRRule,
                Value:  "194.221.250.0/24",
                Action: VPS,
        },
}
