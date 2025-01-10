package nova

import (
	"errors"
	"io/fs"
	"os"
	"reflect"
	"strings"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

// returns a list of CustomSkins based on whats in the custom_skins folder
func GetCustomSkins(custom_skin_dir []fs.DirEntry) (skins []CustomSkin) {
	helpers.Print("Compiling Skins")
	skinsData := fileIO.ReadCSV(inSkinsFolder("custom_skins"))

	credits := skinsData.GetIndexOfColumn("credit")

	discordUIDs := getDiscordUIDs()
	infoMaps := []map[string]string{discordUIDs}
	mapType := []cred.CreditSource{cred.Discord}

	reqLength := skinsData.NumHeaders()
	skins = make([]CustomSkin, 0, skinsData.Rows())

	for row := range skinsData.Rows() {
		s := skinsData.GetRow(row)
		skin, err := CSVLineToCustomSkin(s, custom_skin_dir, reqLength)

		if err != nil {
			helpers.Print("Get Custom Skin Error", s)
			helpers.HandleExcept(err, ErrMalformedRow)
			continue
		}

		credit := skinsData.GetCell(row, credits)
		creditInfo, creditType := assignCredits(credit, infoMaps, mapType)

		if creditType != cred.Default {
			skin.addCredits(cred.NewCredit(credit, creditInfo, creditType))
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
	discordCreditData := fileIO.ReadCSV(inAssetsFolder("DISCORD_UIDS"))
	fileContents := discordCreditData.GetContents()

	uidMap := make(map[string]string, discordCreditData.Rows())

	for _, row := range fileContents {
		discordName := row[0]
		UID := row[1]
		uidMap[discordName] = UID
	}

	return uidMap
}

func CSVLineToCustomSkin(s string, custom_skin_dir []os.DirEntry, reqLength int) (cs *CustomSkin, err error) {
	ss := strings.Split(s, ",")

	if len(ss) != reqLength {
		return EmptyCustomSkin(), ErrMalformedRow
	}

	bodyS, forceArmourS, droneS := ss[1], ss[2], ss[3]

	body, _ := fileIn(bodyS, custom_skin_dir)
	forceArmour, _ := fileIn(forceArmourS, custom_skin_dir)
	drone, _ := fileIn(droneS, custom_skin_dir)

	cs = NewCustomSkin(ss[0], ss[4], ss[5]).addBody(body).addForceA(forceArmour).addDrone(drone)

	cs.generateZipFile()

	return
}

// TODO replace this with the SearchWithFunc when you update the helpers library version used
func fileIn(s string, arr []os.DirEntry) (f fileIO.File, err error) {
	f = *fileIO.EmptyFile()
	err = errors.New("file not found")
	// TODO why are you passing through variables that could simply be part of the nova
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
