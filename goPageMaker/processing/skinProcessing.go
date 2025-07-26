package processing

import (
	"errors"
	"io/fs"
	"os"
	"reflect"
	"strings"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
	"github.com/areon546/go-files/files"
)

// returns a list of CustomSkins based on whats in the custom_skins folder
func GetCustomSkins(custom_skin_dir []fs.DirEntry) (skins []nova.CustomSkin) {
	helpers.Print("Reading Skins In Directory")

	filename := inSkinsFolder("custom_skins", "csv")

	print(filename, "in GetCustomSkins")

	skinsData, err := files.ReadCSV(filename, true)

	helpers.Handle(err)

	credits := skinsData.IndexOfCol("credit")

	discordUIDs := getDiscordUIDs()
	infoMaps := []map[string]string{discordUIDs}
	mapType := []cred.CreditSource{cred.Discord}

	print("Skinsdata.cols", skinsData.Cols())
	reqLength := skinsData.Cols()
	skins = make([]nova.CustomSkin, 0, skinsData.Rows())

	for row := range skinsData.Rows() {
		s := skinsData.Row(row)

		// Headers
		if strings.HasPrefix(s, "name,body_artwork,body_force_armor_artwork,drone_artwork,jet_angle,jet_distance") {
			continue
		}

		skin, err := CSVLineToCustomSkin(s, custom_skin_dir, reqLength)
		if err != nil {
			helpers.Print("Get Custom Skin Error", s)
			helpers.HandleExcept(err, nova.ErrMalformedRow)
			continue
		}

		print("Processing Skin:", skin.Name())
		skin.GenerateZipFile()
		// print("Zip complete")

		credit := skinsData.Cell(row, credits)
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
	fileContents := discordCreditData.Contents()

	helpers.Handle(err)

	uidMap := make(map[string]string, discordCreditData.Rows())

	for _, row := range fileContents {
		discordName := row[0]
		UID := row[1]
		uidMap[discordName] = UID
	}

	return uidMap
}

func CSVLineToCustomSkin(s string, custom_skin_dir []os.DirEntry, reqLength int) (cs *nova.CustomSkin, err error) {
	ss := strings.Split(s, ",")

	if len(ss) != reqLength {
		return nova.EmptyCustomSkin(), nova.ErrMalformedRow
	}

	bodyS, forceArmourS, droneS := ss[1], ss[2], ss[3]

	body, _ := fileIn(bodyS, custom_skin_dir)
	forceArmour, _ := fileIn(forceArmourS, custom_skin_dir)
	drone, _ := fileIn(droneS, custom_skin_dir)

	cs = nova.NewCustomSkin(ss[0], ss[4], ss[5]).AddBody(body).AddForceA(forceArmour).AddDrone(drone)

	return
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
