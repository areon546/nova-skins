package processing

import (
	"errors"
	"io/fs"
	"os"
	"reflect"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/dirs"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/log"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-files/table"
)

var (
	defaultCredits = cred.GetDefault()
	discordUIDs    = cred.GetDiscordUIDs()

	// parallel arrays of credits to the credit source to have custom handlers of various credit types.
	creditMaps    = []cred.CreditMap{defaultCredits, discordUIDs}
	creditSources = []cred.CreditSource{cred.Default, cred.Discord}
)

// returns a list of CustomSkins based on whats in the custom_skins folder
func GetCustomSkins(custom_skin_dir []fs.DirEntry) (skins []nova.CustomSkin) {
	broadcast("Reading Skins In Directory")

	filename := dirs.Skins() + "custom_skins.csv"

	broadcast("Reading Custom Skin CSV", filename)

	skinsData, err := files.ReadCSV(filename, true)
	helpers.Handle(err)

	// Setup variables for Crediting authors

	reqLength := skinsData.Width()
	log.Debug("SkinProcessing GetCustomSkins", "expected column width:", 7, "column width", reqLength)

	skins = make([]nova.CustomSkin, 0, skinsData.Entries())
	for _, record := range skinsData.Iter() {

		skin, err := recordToCustomSkin(&record, custom_skin_dir, reqLength)
		if err != nil {
			log.Error("processing GetCustomSkins", "skin", skin.Name(), "error", err)
			helpers.HandleExcept(err, nova.ErrMalformedRow)
			continue
		}

		broadcast("Processing Skin:", skin.Name(), skin.Body().Name())
		skin.GenerateZipFile()

		skins = append(skins, *skin)
	}

	return
}

func assignCredits(credit cred.Creditor) (creditInfo cred.CreditorInfo, creditType cred.CreditSource) {
	// assign credit type based on credit info
	for i, m := range creditMaps {
		temp, exists := m[credit]
		if exists {
			creditInfo = temp
			creditType = creditSources[i]
			return
		}
	}

	creditType = cred.Unknown
	return
}

func recordToCustomSkin(record *table.Row, custom_skin_dir []os.DirEntry, reqLength int) (*nova.CustomSkin, error) {
	var err, e error
	var cs *nova.CustomSkin
	if record.Size() != reqLength {
		return nova.EmptyCustomSkin(), nova.ErrMalformedRow
	}
	// name,body_artwork,body_force_armor_artwork,drone_artwork,jet_angle,jet_distance,credit

	log.Debug("processing/skinProcessing.recordToCustomSkin", "record input", record)
	for index := range reqLength {
		s, e := record.Get(index)
		if e != nil {
			errors.Join(err, e)
		}
		switch index {
		case 0: // NAME
			cs = nova.NewCustomSkin(s)
		case 1: // BODY
			body := fileIn(s, custom_skin_dir) // NOTE: We check

			cs.AddBody(body)
		case 2: // FORCE ARMOUR
			forceArmour := fileIn(s, custom_skin_dir)

			cs.AddForceA(forceArmour)
		case 3: // DRONE
			drone := fileIn(s, custom_skin_dir)

			cs.AddDrone(drone)
		case 4: // ANGLE
			cs.AddAngle(s)
		case 5: // DISTANCE
			cs.AddDistance(s)
		case 6: // CREDITS
			credit := cred.Creditor(s)
			creditInfo, creditType := assignCredits(credit)

			cs.AddCredits(cred.NewCredit(credit, creditInfo, creditType))
		default:
		}
	}

	log.Error("processing/skinProcessing.recordToCustomSkin", "error", e)

	return cs, err
}

// TODO: replace this with the SearchWithFunc when you update the helpers library version used
func fileIn(filename string, arr []os.DirEntry) (f files.File) {
	f = *files.EmptyFile()

	// If filename empty, return the
	filenameEmpty := reflect.DeepEqual(filename, "")
	if filenameEmpty {
		return f
	}

	// Go through
	for _, dirEntry := range arr {

		filenameMatch := reflect.DeepEqual(filename, dirEntry.Name())
		notDirectory := !dirEntry.IsDir()
		if filenameMatch && notDirectory {
			return *openCustomSkin(dirEntry)
		}
	}

	// Return emptyFile to deal with potential edge casts
	return
}
