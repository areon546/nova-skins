package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ~~~~~~~~~~~~~~~~~ CustomSkin
type CustomSkin struct {
	pictures []File
	credit   *Credit

	name        string
	body        string
	forceArmour string
	drone       string
	angle       string
	distance    string
}

func NewCustomSkin(name, angle, distance string) *CustomSkin {
	return &CustomSkin{name: name, angle: angle, distance: distance}
}

func (c *CustomSkin) addBody(b string) *CustomSkin {
	c.body = b
	return c
}

func (c *CustomSkin) addForceA(s string) *CustomSkin {
	c.forceArmour = s
	return c
}

func (c *CustomSkin) addDrone(s string) *CustomSkin {
	c.drone = s
	return c
}

func (c *CustomSkin) String() string {
	return c.name
}

func convertCSVLineToCustomSkin(s string) *CustomSkin {
	ss := strings.Split(s, ",")
	c := CustomSkin{name: ss[0], body: ss[1], forceArmour: ss[2], drone: ss[3], angle: ss[4], distance: ss[5]}
	return &c
}

func (c *CustomSkin) toCSVLine() string {
	return format("%s,%s,%s,%s,%s,%s", c.name, c.body, c.forceArmour, c.drone, c.getAngle(), c.getDistance())
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

// returns a list of CustomSkins based on whats in the custom_skins folder
func getCustomSkins() (skins []CustomSkin) {
	skinsData := readCSV(skinFolder() + "custom_skins")
	names := skinsData.getIndexOfColumn("name")
	angles := skinsData.getIndexOfColumn("jet_angle")
	distances := skinsData.getIndexOfColumn("jet_distance")
	body := skinsData.getIndexOfColumn("body_artwork")
	forces := skinsData.getIndexOfColumn("body_force_armor_artwork")
	drones := skinsData.getIndexOfColumn("drone_artwork")

	print(skinsData.headings)

	skins = make([]CustomSkin, 0, skinsData.Rows())
	print(skinsData.Rows())
	reqLength := skinsData.numHeaders()

	for _, v := range skinsData.contents {
		if len(v) == reqLength || len(v) == 7 {
			// print(i, v, body, forces, drones)

			name := v[names]
			distance := v[distances]
			angle := v[angles]

			skin := NewCustomSkin(name, distance, angle).addBody(v[body]).addForceA(v[forces]).addDrone(v[drones])
			skins = append(skins, *skin)

			printf("appropriate length: %d, %s", len(v), skin)
		} else {
			printf("malformed csv, required length: %d, length: %d, %s,", reqLength, len(v), v)
		}
	}

	return
}

func (c *CustomSkin) formatCredits() string {
	if c.credit == nil {
		return ""
	}
	return constructMarkDownLink(false, c.credit.getCredit(), "")
}

// ~~~~~~~~~~~~~~~~~~~ AssetPage

type AssetsPage struct {
	MarkdownFile
	pageNumber int
	maxSkins   int
	skinsC     int

	skins []CustomSkin
}

func NewAssetsPage(filename string, pageNum int, path string) *AssetsPage {
	return &AssetsPage{MarkdownFile: *NewMarkdownFile(filename, path), pageNumber: pageNum, maxSkins: 10, skinsC: 0}
}

func (a *AssetsPage) String() string {
	return a.filename
}

func (a *AssetsPage) bufferPagePreffix() error {
	// write to file:
	// Page #
	a.append(fmt.Sprintf("# Page %d", a.pageNumber))
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

		a.appendMarkdownLink(prev, (path + prevF))
	}

	a.appendMarkdownLink(curr, (path + currF))
	a.appendMarkdownLink(next, (path + nextF))

	return nil
}

func (a *AssetsPage) bufferCustomSkins() {
	// this writes to the custom skins stuff and adds the data, in markdown
	path := "https://github.com/areon546/NovaDriftCustomSkinRepository/raw/main"

	for _, skin := range a.skins {
		a.appendNewLine()
		// append

		a.append(format("*%s*: %s", skin.name, skin.formatCredits()))
		a.appendNewLine()

		a.append("`" + skin.toCSVLine() + "`")
		a.appendNewLine()

		a.appendMarkdownEmbed(constructPath(path, "custom_skins", skin.body))
		a.appendMarkdownEmbed(constructPath(path, "custom_skins", skin.forceArmour))
		a.appendMarkdownEmbed(constructPath(path, "custom_skins", skin.drone))
		// TODO append links to media  but how do we determine if there are media files?

		a.appendNewLine()
	}
}

func (a *AssetsPage) writeBuffer() {
	a.writeFile()

	// print(a.contentBuffer)
	print("Writing to: ", a)
}

func (a *AssetsPage) addCustomSkins(cs []CustomSkin) {
	numSkins := min(10, len(cs))
	for a.skinsC < numSkins {
		a.skins = append(a.skins, cs[a.skinsC])
		a.skinsC++
	}
}

func constructAssetPages(skins []CustomSkin) (pages []AssetsPage) {
	numSkins := len(skins)
	print("aa ", numSkins)
	numFiles := numSkins / 10

	if numSkins%10 != 0 {
		numFiles++
	}
	print(numFiles)

	for i := range numFiles {
		// create a new file
		pageNum := i + 1
		a := NewAssetsPage(constructPath("", getPagesFolder(), format("Page_%d", pageNum)), pageNum, "2")

		a.bufferPagePreffix()

		skinSlice, err := getNextSlice(skins, i)
		handle(err)

		a.addCustomSkins(skinSlice)
		a.bufferCustomSkins()
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
