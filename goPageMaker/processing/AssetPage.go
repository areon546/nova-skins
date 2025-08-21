package processing

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-files/formatter"
)

// ~~~~~~~~~~~~~~~~~~~ AssetPage
type AssetsPage struct {
	formatter.FormattedFile
	pageNumber int
	maxSkins   int
	skinsC     int

	skins []nova.CustomSkin
}

func NewAssetsPage(path, filename string, pageNum int) *AssetsPage {
	return &AssetsPage{FormattedFile: *formatter.NewMarkdownFile(path, filename, ""), pageNumber: pageNum, maxSkins: 10, skinsC: 0}
}

func (a *AssetsPage) String() string {
	return a.Name()
}

func (a *AssetsPage) bufferPagePreffix() error {
	// write to file:
	// Page #
	a.AppendHeading(1, format("Page %d", a.pageNumber))

	a.AppendEmptyLine()
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
		a.AppendBold(skin.Name())
		a.Append(": ", false)
		a.Append(skin.FormatCredits(a.Fmt), false)
		a.AppendEmptyLine()
		a.AppendEmptyLine()

		a.AppendNewLine(skin.ToTable(a.Fmt))
		a.AppendNewLine("Copy this: `" + skin.ToCSVLine() + "`")
		a.AppendEmptyLine()
		a.AppendLink("Download Me", skin.Zip().Name())
		a.AppendEmptyLine()

		a.AppendEmptyLine()
		if !files.FilesEqual(*skin.Body(), *files.EmptyFile()) {
			a.AppendCustomSkinFile(path, skin.Body().Name())
		}
		if !files.FilesEqual(*skin.ForceArmour(), *files.EmptyFile()) {
			a.AppendCustomSkinFile(path, skin.ForceArmour().Name())
		}

		a.Append("", true)

		a.AppendEmptyLine()
		if !files.FilesEqual(*skin.Drone(), *files.EmptyFile()) {
			a.AppendCustomSkinFile(path, skin.Drone().Name())
		}

		a.AppendEmptyLine()
	}
}

func (a *AssetsPage) AppendCustomSkinFile(path, filename string) {
	a.AppendEmbed(files.ConstructFilePath(path+"/custom_skins", filename, ""), filename)
	a.AppendEmptyLine()
}

func (a *AssetsPage) writeBuffer() {
	broadcast("Writing to: ", a)
	// print(a.contentBuffer)
	a.WriteContents()
}

func (a *AssetsPage) addCustomSkins(cs []nova.CustomSkin) {
	numSkins := min(10, len(cs))
	for a.skinsC < numSkins {
		a.skins = append(a.skins, cs[a.skinsC])
		a.skinsC++
	}
}
