package nova

import (
	"errors"
	"io/fs"
	"reflect"
	"strconv"
	"strings"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
)

// ~~~~~~~~~~~~~~~~~ CustomSkin
type CustomSkin struct {
	pictures []fileIO.File
	credit   cred.CreditType

	name        string
	body        fs.DirEntry
	forceArmour fs.DirEntry
	drone       fs.DirEntry
	angle       string
	distance    string
}

func NewCustomSkin(name, angle, distance string) *CustomSkin {
	return &CustomSkin{name: name, angle: angle, distance: distance}
}

func (c *CustomSkin) addBody(f fs.DirEntry) *CustomSkin {
	c.body = f
	return c
}

func (c *CustomSkin) addForceA(s fs.DirEntry) *CustomSkin {
	c.forceArmour = s
	return c
}

func (c *CustomSkin) addDrone(f fs.DirEntry) *CustomSkin {
	c.drone = f
	return c
}

func (cs *CustomSkin) addCredits(c cred.CreditType) {
	cs.credit = c
}

func (c *CustomSkin) String() string {
	return c.name
}

func convertCSVLineToCustomSkin(s string, custom_skin_dir []fs.DirEntry, reqLength int) (c *CustomSkin, err error) {
	ss := strings.Split(s, ",")

	if len(ss) == reqLength {

		bodyS, forceArmourS, droneS := ss[1], ss[2], ss[3]

		body, forceArmour, drone := fileIn(bodyS, custom_skin_dir), fileIn(forceArmourS, custom_skin_dir), fileIn(droneS, custom_skin_dir)
		c = NewCustomSkin(ss[0], ss[4], ss[5]).addBody(body).addForceA(forceArmour).addDrone(drone)

		return
	}
	err = errors.New("malformed Row")
	// helpers.Print(len(ss))

	return
}

func fileIn(s string, arr []fs.DirEntry) fs.DirEntry {

	for _, v := range arr {
		if reflect.DeepEqual(s, v.Name()) {
			return v
		}
	}

	return nil
}

func (c *CustomSkin) toCSVLine() string {
	body, fA, drone := "", "", ""

	if c.body != nil {
		body = c.body.Name()
	}

	if c.forceArmour != nil {
		fA = c.forceArmour.Name()
	}

	if c.drone != nil {
		drone = c.drone.Name()
	}
	return format("%s,%s,%s,%s,%s,%s", c.name, body, fA, drone, c.getAngle(), c.getDistance())
}

func (c *CustomSkin) getAngle() string {
	// try to convert s to an integer, if it fails, return nothing
	_, err := strconv.Atoi(c.angle)
	if err != nil {
		return ""
	} else {
		return c.angle
	}
}

func (c *CustomSkin) getDistance() string {
	// try to convert to an integer
	_, err := strconv.Atoi(c.distance)
	if err != nil {
		return ""
	} else {
		return c.distance
	}
}

func (c *CustomSkin) FormatCredits() string {
	if c.credit == nil {
		return ""
	}
	return fileIO.ConstructMarkDownLink(false, c.credit.ConstructName(), c.credit.ConstructLink())
}

// returns a list of CustomSkins based on whats in the custom_skins folder
func GetCustomSkins(custom_skin_dir []fs.DirEntry) (skins []CustomSkin) {
	skinsData := fileIO.ReadCSV(inSkinsFolder("custom_skins"))
	credits := skinsData.GetIndexOfColumn("credit")

	discordUIDs := getDiscordUIDs()
	infoMaps := []map[string]string{discordUIDs}
	mapType := []cred.CreditSource{cred.Discord}

	reqLength := skinsData.NumHeaders()
	skins = make([]CustomSkin, 0, skinsData.Rows())

	for row := range skinsData.Rows() {
		s := skinsData.GetRow(row)
		skin, err := convertCSVLineToCustomSkin(s, custom_skin_dir, reqLength)
		if err == nil {
			// print(i, v, body, forces, drones)

			credit := skinsData.GetCell(row, credits)
			creditInfo, creditType := assignCredits(credit, infoMaps, mapType)

			if creditType != cred.Default {
				skin.addCredits(cred.NewCredit(credit, creditInfo, creditType))
			}

			skins = append(skins, *skin)

			// printf("appropriate length: %d, %s", len(v), skin)
		} else {
			// printf("malformed csv, required length: %d, length: %d, %s,", reqLength, len(s), s)
		}
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
