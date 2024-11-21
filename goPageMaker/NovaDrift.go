package main

import (
	"fmt"
	"strings"
)

// ~~~~~~~~~~~~~~~~~ CustomSkin
type CustomSkin struct {
	// pictures []File
	// credit      string

	name        string
	body        string
	forceArmour string
	drone       string
	angle       int
	distance    int
}

func NewCustomSkin(name string, angle, distance int) *CustomSkin {
	return &CustomSkin{name: name, angle: angle, distance: distance}
}

func (c *CustomSkin) addSkin(s string) *CustomSkin {
	c.body = s
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

func (c *CustomSkin) toString() string {
	return c.name
}

func convertCSVLineToCustomSkin(s string) *CustomSkin {
	ss := strings.Split(s, ",")
	c := CustomSkin{name: ss[0], body: ss[1], forceArmour: ss[2], drone: ss[3], angle: convertToInteger(ss[4]), distance: convertToInteger(ss[5])}
	return &c
}

func (c *CustomSkin) toCSVLine() string {
	return format("%s,%s,%s,%s,%d,%d", c.name, c.body, c.forceArmour, c.drone, c.angle, c.distance)
}

// ~~~~~~~~~~~~~~~~~~~ AssetPage

type AssetsPage struct {
	MarkdownFile
	pageNumber int

	skins []CustomSkin
}

func NewAssetsPage(filename string, pageNum int, path string) *AssetsPage {
	return &AssetsPage{MarkdownFile: *NewMarkdownFile(filename, path), pageNumber: pageNum}
}

func (a *AssetsPage) bufferPagePreffix() error {
	// write to file:
	// Page #
	// prev next
	a.append(fmt.Sprintf("# Page %d", a.pageNumber))
	err := a.bufferPrevNextPage()

	return err
}

func (a *AssetsPage) bufferPrevNextPage() error {
	path := "../pages/"

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

func (a *AssetsPage) addCustomSkins(cs []CustomSkin) {
	a.skins = cs
}

func (a *AssetsPage) bufferCustomSkins() {
	// TODO this writes to the custom skins stuff and adds the data, in markdown
	path := "https://github.com/areon546/NovaDriftCustomSkinRepository/raw/main"

	for _, skin := range a.skins {
		a.appendNewLine()
		// append

		a.append(skin.toCSVLine())
		a.appendMarkdownEmbed(constructPath(path, "custom_skins", skin.body))

		// append links to media TODO

		a.appendNewLine()
	}
}

func (a *AssetsPage) writeBuffer() {
	a.writeFile()

	print(a.contentBuffer)
}
