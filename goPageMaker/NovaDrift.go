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
	File
	pageNumber int

	skins []CustomSkin
}

func NewAssetsPage(filename string, pageNum int, path string) *AssetsPage {
	return &AssetsPage{File: *NewFileWithSuffix(filename, "md", path), pageNumber: pageNum}
}

func (a *AssetsPage) writePagePreffix() error {
	// write to file:
	// Page #
	// prev next
	a.append(fmt.Sprintf("Page %d", a.pageNumber))
	err := a.writePrevNextPage()

	return err
}

func (a *AssetsPage) writePrevNextPage() error {
	path := "../pages/"
	links := ""

	prev := format("Pade%d.md", a.pageNumber-1)
	prevD := format("Page %d", a.pageNumber-1)
	curr := format("Pade%d.md", a.pageNumber)
	currD := format("Page %d", a.pageNumber-1)
	next := format("Pade%d.md", a.pageNumber+1)
	nextD := format("Page %d", a.pageNumber+1)

	if a.pageNumber > 1 {

		links += constructMarkdownLink(false, prevD, (path + prev))
	}

	links += constructMarkdownLink(false, currD, (path + curr))
	links += constructMarkdownLink(false, nextD, (path + next))

	a.append(links)

	return nil
}

func (a *AssetsPage) addCustomSkins(cs []CustomSkin) {
	a.skins = cs
}

func (a *AssetsPage) writeCustomSkins(cs []CustomSkin) {
	// TODO this writes to the custom skins stuff and adds the data, in markdown
}
