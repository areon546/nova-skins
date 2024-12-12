package main

// TODO make this an interface with constructName in it
type Credit struct {
	name      string
	otherInfo string
	cType     CreditType
}

func NewCredit(name, other string) *Credit {
	return &Credit{name: name, otherInfo: other}
}

func (c *Credit) getCredit() string {
	return c.cType.constructName(*c)
}

type CreditType interface {
	constructName(Credit) string
	constructLink(Credit) string
}

type DiscordCredit struct {
}

func (d *DiscordCredit) constructName(c Credit) string {
	return format("@%s", c.name)
}

func (d *DiscordCredit) constructLink(c Credit) string {
	return format("discordapp.com/users/%s", c.otherInfo)
}
