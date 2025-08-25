package processing

import (
	"fmt"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/dirs"
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
	suff := "html"

	fmt.Println(a.pageNumber)

	prev := format("Page_%d", a.pageNumber-1)
	prevF := format("%s.%s", prev, suff)
	curr := format("Page_%d", a.pageNumber)
	currF := format("%s.%s", curr, suff)
	next := format("Page_%d", a.pageNumber+1)
	nextF := format("%s.%s", next, suff)

	if a.pageNumber > 1 {
		a.AppendLink(prev, (path + prevF))
	}

	a.AppendLink(curr, (path + currF))
	a.AppendLink(next, (path + nextF))

	return nil
}

func (a *AssetsPage) bufferCustomSkins() {
	// this writes to the custom skins stuff and adds the data, in markdown

	for _, skin := range a.skins {
		a.AppendEmptyLine()

		heading := a.Fmt.Bold(skin.Name()) + ":"
		a.AppendHeading(2, heading)
		a.Append(skin.FormatCredits(a.Fmt))
		a.AppendEmptyLine()
		a.AppendEmptyLine()

		a.AppendNewLine(skin.ToTable(a.Fmt))
		a.AppendNewLine("Copy this: `" + skin.ToCSVLine() + "`")
		a.AppendEmptyLine()
		a.AppendLink("Download Me", dirs.WwwAssetsFolder()+"zips/"+skin.Name()+".zip")
		a.AppendEmptyLine()

		a.AppendEmptyLine()
		if !files.FilesEqual(*skin.Body(), *files.EmptyFile()) {
			a.AppendCustomSkinFile(skin.Body())
		}
		if !files.FilesEqual(*skin.ForceArmour(), *files.EmptyFile()) {
			a.AppendCustomSkinFile(skin.ForceArmour())
		}

		a.AppendEmptyLine()
		if !files.FilesEqual(*skin.Drone(), *files.EmptyFile()) {
			a.AppendCustomSkinFile(skin.Drone())
		}

		a.AppendEmptyLine()
	}
}

func (a *AssetsPage) AppendCustomSkinFile(f *files.File) {
	a.AppendEmbed(dirs.WwwSkinsFolder()+f.Name(), f.Name())
	a.AppendEmptyLine()
}

func (a *AssetsPage) writeBuffer() {
	broadcast("Writing to: ", a.FullName())
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
