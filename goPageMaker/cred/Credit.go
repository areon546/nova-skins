package cred

const (
	Unknown CreditSource = iota
	Default
	Discord
)

type CreditType interface {
	ConstructName() string
	ConstructLink() string
}

type CreditSource int

type Credit struct {
	name      Creditor
	otherInfo CreditorInfo
	CreditType
}

func (c Credit) ConstructName() string { return "unknown" }
func (c Credit) ConstructLink() string { return "/404.html" }

func NewCredit(name Creditor, other CreditorInfo, cType CreditSource) CreditType {
	switch cType {
	case Discord:
		return DiscordCredit{Credit: Credit{name: name, otherInfo: other}}
	case Default:
		return DefaultCredit{}
	case Unknown:
		fallthrough
	default:
		return &Credit{name: name, otherInfo: other}
	}
}

type DefaultCredit struct{ Credit }

func (c DefaultCredit) ConstructName() string { return "Default Skin" }
func (c DefaultCredit) ConstructLink() string { return "https://novadrift.io/" }

type DiscordCredit struct{ Credit }

func (d DiscordCredit) ConstructName() string {
	return format("@%s", d.name)
}

func (d DiscordCredit) ConstructLink() string {
	return format("https://discord.com/users/%s", d.otherInfo)
}
