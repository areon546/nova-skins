package main

import "reflect"

type CreditType interface {
	constructName() string
	constructLink() string
}

type Credit struct {
	name      string
	otherInfo string
	CreditType
}

func (c Credit) constructName() string { return "" }
func (c Credit) constructLink() string { return "" }

func NewCredit(name, other, cType string) CreditType {
	if reflect.DeepEqual(cType, "discord") {
		return DiscordCredit{Credit: Credit{name: name, otherInfo: other}}
	}
	return &Credit{name: name, otherInfo: other}

	return nil
}

type DiscordCredit struct{ Credit }

func (d DiscordCredit) constructName() string {
	return format("@%s", d.name)
}

func (d DiscordCredit) constructLink() string {
	return format("https://discord.com/users/%s", d.otherInfo)
}
