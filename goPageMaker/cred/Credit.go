package cred

import (
	"reflect"
)

const (
	Default CreditSource = 0
	Discord CreditSource = 1
)

type CreditType interface {
	ConstructName() string
	ConstructLink() string
}

type CreditSource int

// TODO make this an interface with constructName in it
type Credit struct {
	name      string
	otherInfo string
	CreditType
}

func (c Credit) ConstructName() string { return "" }
func (c Credit) ConstructLink() string { return "" }

func NewCredit(name, other string, cType CreditSource) CreditType {
	if reflect.DeepEqual(cType, "discord") {
		return DiscordCredit{Credit: Credit{name: name, otherInfo: other}}
	}
	return &Credit{name: name, otherInfo: other}
}

type DiscordCredit struct{ Credit }

func (d DiscordCredit) ConstructName() string {
	return format("@%s", d.name)
}

func (d DiscordCredit) ConstructLink() string {
	return format("https://discord.com/users/%s", d.otherInfo)
}
