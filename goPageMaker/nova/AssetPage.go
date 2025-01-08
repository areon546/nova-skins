package nova

import (
	"fmt"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

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
	return a.Name()
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

func (a *AssetsPage) bufferCustomSkins() {
	// this writes to the custom skins stuff and adds the data, in markdown
	path := ".."

	for _, skin := range a.skins {
		a.AppendNewLine()

		a.Append(format("**%s**: %s", skin.name, skin.FormatCredits()))
		a.AppendNewLine()

		a.Append(skin.ToTable())
		a.Append("`" + skin.ToCSVLine() + "`")
		a.AppendNewLine()

		// helpers.Print("Buffering skin: ", skin)

		if !fileIO.FilesEqual(skin.Body, *fileIO.EmptyFile()) {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.Body.Name()))
			a.AppendNewLine()
		}
		if !fileIO.FilesEqual(skin.ForceArmour, *fileIO.EmptyFile()) {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.ForceArmour.Name()))
			a.AppendNewLine()
		}
		if !fileIO.FilesEqual(skin.Drone, *fileIO.EmptyFile()) {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.Drone.Name()))
			a.AppendNewLine()
		}

		a.AppendMarkdownLink("Download Me", skin.zip.GetName())

		a.AppendNewLine()
	}
}

func (a *AssetsPage) writeBuffer() {
	helpers.Print("Writing to: ", a)
	// print(a.contentBuffer)
	a.Write(a.Contents())
}

func (a *AssetsPage) addCustomSkins(cs []CustomSkin) {
	numSkins := min(10, len(cs))
	for a.skinsC < numSkins {
		a.skins = append(a.skins, cs[a.skinsC])
		a.skinsC++
	}
}
