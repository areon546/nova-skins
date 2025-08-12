package processing

import (
	"errors"
	"io/fs"
	"os"
	"reflect"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-files/table"
)

// returns a list of CustomSkins based on whats in the custom_skins folder
func GetCustomSkins(custom_skin_dir []fs.DirEntry) (skins []nova.CustomSkin) {
	helpers.Print("Reading Skins In Directory")

	filename := inSkinsFolder("custom_skins", "csv")

	print(filename, "in GetCustomSkins")

	skinsData, err := files.ReadCSV(filename, true)

	helpers.Handle(err)

	credits := skinsData.IndexOf("credit")

	discordUIDs := getDiscordUIDs()
	infoMaps := []map[string]string{discordUIDs}
	mapType := []cred.CreditSource{cred.Discord}

	print("Skinsdata.cols", skinsData.Width())
	reqLength := skinsData.Width()
	skins = make([]nova.CustomSkin, 0, skinsData.Entries())

	for rowNumber, record := range skinsData.Iter() {

		skin, err := recordToCustomSkin(&record, custom_skin_dir, reqLength)
		if err != nil {
			helpers.Print("Get Custom Skin Error", skin)
			helpers.HandleExcept(err, nova.ErrMalformedRow)
			continue
		}

		print("Processing Skin:", skin.Name())
		skin.GenerateZipFile()
		// print("Zip complete")

		credit, _ := skinsData.Cell(rowNumber, credits)
		creditInfo, creditType := assignCredits(credit, infoMaps, mapType)

		if creditType != cred.Default {
			skin.AddCredits(cred.NewCredit(credit, creditInfo, creditType))
		}

		skins = append(skins, *skin)
	}

	return
}

func assignCredits(credit string, creditInfoMaps []map[string]string, mapTypes []cred.CreditSource) (creditInfo string, creditType cred.CreditSource) {
	// assign credit type based on credit info
	for i, m := range creditInfoMaps {
		temp, exists := m[credit]
		if exists {
			creditInfo = temp
			creditType = mapTypes[i]
			return
		}
	}

	creditType = cred.Default
	return
}

func getDiscordUIDs() map[string]string {
	discordCreditData, err := files.ReadCSV(inAssetsFolder("DISCORD_UIDS", "csv"), true)

	helpers.Handle(err)

	uidMap := make(map[string]string, discordCreditData.Entries())

	for _, row := range discordCreditData.Iter() {
		discordName, _ := row.Get(0)
		UID, _ := row.Get(1)
		uidMap[discordName] = UID
	}

	return uidMap
}

func recordToCustomSkin(record *table.Row, custom_skin_dir []os.DirEntry, reqLength int) (*nova.CustomSkin, error) {
	var err, e error
	if record.Size() != reqLength {
		return nova.EmptyCustomSkin(), nova.ErrMalformedRow
	}
	// name,body_artwork,body_force_armor_artwork,drone_artwork,jet_angle,jet_distance,credit
	var name, bodyFn, forceArmourFn, droneFn, angle, distance string

	for index := range reqLength {
		print(index)
		switch index {
		case 0:
			print("ASD")
		default:
			print("DEFAULT")
		}
	}

	name, e = record.Get(0)
	bodyFn, e = record.Get(1)
	forceArmourFn, e = record.Get(2)
	droneFn, e = record.Get(3)
	angle, e = record.Get(4)
	distance, e = record.Get(5)

	body, e := fileIn(bodyFn, custom_skin_dir)
	forceArmour, e := fileIn(forceArmourFn, custom_skin_dir)
	drone, e := fileIn(droneFn, custom_skin_dir)

	print(e)

	cs := nova.NewCustomSkin(name, angle, distance)
	cs.AddBody(body)
	cs.AddForceA(forceArmour)
	cs.AddDrone(drone)

	return cs, err
}

// TODO: replace this with the SearchWithFunc when you update the helpers library version used
func fileIn(s string, arr []os.DirEntry) (f files.File, err error) {
	f = *files.EmptyFile()
	err = errors.New("file not found")
	// TODO: why are you passing through variables that could simply be part of the nova
	// make custom_skin directory a nova variable
	if reflect.DeepEqual("", s) {
		return f, errors.New("empty file")
	}

	for _, v := range arr {
		if reflect.DeepEqual(s, v.Name()) {
			return *openCustomSkin(v), nil
		}
	}
	return
}

func emptyOSFile() os.File {
	return os.File{}
}
