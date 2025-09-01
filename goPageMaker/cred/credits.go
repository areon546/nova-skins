package cred

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/dirs"
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

type (
	Creditor     string
	CreditorInfo string
	CreditMap    map[Creditor]CreditorInfo
)

func GetDefault() CreditMap {
	return CreditMap{Creditor("default"): CreditorInfo("https://blog.novadrift.io/customskins/")}
}

func GetDiscordUIDs() CreditMap {
	discordCreditData, err := files.ReadCSV(dirs.Assets()+"DISCORD_UIDS.csv", true)

	helpers.Handle(err)

	uidMap := make(CreditMap, discordCreditData.Entries())

	for _, row := range discordCreditData.Iter() {
		discordName, _ := row.Get(0)
		UID, _ := row.Get(1)
		uidMap[Creditor(discordName)] = CreditorInfo(UID)
	}

	return uidMap
}
