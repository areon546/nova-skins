package nova

import (
	"errors"
	"fmt"
	"io/fs"
	"reflect"
	"strconv"
	"strings"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
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

			if !reflect.DeepEqual(creditType, "default") {
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

// ~~~~~~~~~~~~~~~~~~~ AssetPage

type AssetsPage struct {
	fileIO.MarkdownFile
	pageNumber int
	maxSkins   int
	skinsC     int

	skins []CustomSkin
}

func NewAssetsPage(filename string, pageNum int, path string) *AssetsPage {
	return &AssetsPage{MarkdownFile: *fileIO.NewMarkdownFile(filename, path), pageNumber: pageNum, maxSkins: 10, skinsC: 0}
}

func (a *AssetsPage) String() string {
	return a.GetFileName()
}

func (a *AssetsPage) bufferPagePreffix() error {
	// write to file:
	// Page #
	a.Append(fmt.Sprintf("# Page %d", a.pageNumber))
	// prev next
	err := a.bufferPrevNextPage()

	return err
}

func (a *AssetsPage) bufferPageSuffix() error {
	// write to file:
	// prev next
	err := a.bufferPrevNextPage()

	return err
}

func (a *AssetsPage) bufferPrevNextPage() error {
	path := "./"

	prev := format("Page_%d", a.pageNumber-1)
	prevF := format("%s.md", prev)
	curr := format("Page_%d", a.pageNumber)
	currF := format("%s.md", curr)
	next := format("Page_%d", a.pageNumber+1)
	nextF := format("%s.md", next)

	if a.pageNumber > 1 {

		a.AppendMarkdownLink(prev, (path + prevF))
	}

	a.AppendMarkdownLink(curr, (path + currF))
	a.AppendMarkdownLink(next, (path + nextF))

	return nil
}

func (a *AssetsPage) bufferCustomSkins(download bool) {
	// this writes to the custom skins stuff and adds the data, in markdown
	path := "https://github.com/areon546/NovaDriftCustomSkinRepository/raw/main"

	for _, skin := range a.skins {
		a.AppendNewLine()

		a.Append(format("**%s**: %s", skin.name, skin.FormatCredits()))
		a.AppendNewLine()

		a.Append("`" + skin.toCSVLine() + "`")
		a.AppendNewLine()

		// helpers.Print("Buffering skin: ", skin)

		if skin.body != nil {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.body.Name()))
		}
		if skin.forceArmour != nil {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.forceArmour.Name()))
		}
		if skin.drone != nil {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.drone.Name()))
		} // TODO append links to media  but how do we determine if there are media files?

		if download {
			a.AppendMarkdownLink("Download Me", fileIO.ConstructPath(path, "assets", format("%s.zip", skin.name)))
		}

		a.AppendNewLine()
	}
}

func (a *AssetsPage) writeBuffer() {
	a.WriteFile()

	// print(a.contentBuffer)
	helpers.Print("Writing to: ", a)
}

func (a *AssetsPage) addCustomSkins(cs []CustomSkin) {
	numSkins := min(10, len(cs))
	for a.skinsC < numSkins {
		a.skins = append(a.skins, cs[a.skinsC])
		a.skinsC++
	}
}

func ConstructAssetPages(skins []CustomSkin) (pages []AssetsPage) {
	numSkins := len(skins)
	// print("skins ", numSkins)
	numFiles := numSkins / 10

	if numSkins%10 != 0 {
		numFiles++
	}
	// print("filesToCreate", numFiles)

	for i := range numFiles {
		// create a new file
		pageNum := i + 1
		a := NewAssetsPage(fileIO.ConstructPath("", pagesFolder(), format("Page_%d", pageNum)), pageNum, "2")

		a.bufferPagePreffix()

		skinSlice, err := getNextSlice(skins, i)
		helpers.Handle(err)

		a.addCustomSkins(skinSlice)
		a.bufferCustomSkins(false)
		a.bufferPageSuffix()

		pages = append(pages, *a)
		// print(a)

		a.writeBuffer()
	}

	// a := NewAssetsPage(constructPath("", getPagesFolder(), "test"), 0, "")

	// a.bufferPagePreffix()
	// a.addCustomSkins(skins)
	// a.bufferCustomSkins()
	// a.bufferPageSuffix()

	// pages = append(pages, *a)
	return
}

func getNextSlice(skins []CustomSkin, i int) (subset []CustomSkin, err error) {
	numSkins := len(skins)

	if i < 0 || i > (len(skins)/10+1) {
		err = errors.New("index out of bounds for CustomSkins array")
	}

	min, max := i*10, (i+1)*10

	if max > numSkins {
		max = numSkins
	}

	return skins[min:max], err
}

func ConstructZipFiles(skins []CustomSkin) []fileIO.File {
	return make([]fileIO.File, 0)
}
