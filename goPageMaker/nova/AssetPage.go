package nova

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/formatter"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

// ~~~~~~~~~~~~~~~~~~~ AssetPage
type AssetsPage struct {
	formatter.FormattedFile
	pageNumber int
	maxSkins   int
	skinsC     int

	skins []CustomSkin
}

func NewAssetsPage(filename string, pageNum int, path string) *AssetsPage {
	return &AssetsPage{FormattedFile: *formatter.NewMarkdownFile(filename, path), pageNumber: pageNum, maxSkins: 10, skinsC: 0}
}

func (a *AssetsPage) String() string {
	return a.Name()
}

func (a *AssetsPage) bufferPagePreffix() error {
	// write to file:
	// Page #
	a.AppendHeading(1, format("# Page %d", a.pageNumber))
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

		a.AppendLink(prev, (path + prevF))
	}

	a.AppendLink(curr, (path + currF))
	a.AppendLink(next, (path + nextF))

	return nil
}

func (a *AssetsPage) bufferCustomSkins() {
	// this writes to the custom skins stuff and adds the data, in markdown
	path := ".."

	for _, skin := range a.skins {
		a.AppendEmptyLine()

		a.AppendHeading(2, "")
		a.AppendBold(skin.name)
		a.Append(": ", false)
		a.Append(skin.FormatCredits(a.Fmt), false)
		a.AppendEmptyLine()

		a.AppendNewLine(skin.ToTable(a.Fmt))
		a.AppendNewLine("Copy this: `" + skin.ToCSVLine() + "`")
		a.AppendEmptyLine()
		a.AppendLink("Download Me", skin.zip.GetName())

		a.AppendEmptyLine()
		if !fileIO.FilesEqual(skin.Body, *fileIO.EmptyFile()) {
			a.AppendCustomSkinFile(path, skin.Body.Name())
		}
		if !fileIO.FilesEqual(skin.ForceArmour, *fileIO.EmptyFile()) {
			a.AppendCustomSkinFile(path, skin.ForceArmour.Name())
		}

		a.AppendEmptyLine()
		if !fileIO.FilesEqual(skin.Drone, *fileIO.EmptyFile()) {
			a.AppendCustomSkinFile(path, skin.Drone.Name())
		}

		a.AppendEmptyLine()
	}
}

func (a *AssetsPage) AppendCustomSkinFile(path, filename string) {
	a.AppendEmbed(fileIO.ConstructPath(path, "custom_skins", filename))
	a.AppendEmptyLine()
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
