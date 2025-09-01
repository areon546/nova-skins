package processing

import (
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

func (a *AssetsPage) AddCustomSkins(cs []nova.CustomSkin) {
	numSkins := min(10, len(cs))
	for a.skinsC < numSkins {
		a.skins = append(a.skins, cs[a.skinsC])
		a.skinsC++
	}
}

func (a *AssetsPage) BufferPagePreffix() error {
	// write to file:
	// Page #
	a.AppendHeading(1, format("Page %d", a.pageNumber))

	a.AppendEmptyLine()
	// prev next
	err := a.BufferPrevNextPage()

	return err
}

func (a *AssetsPage) BufferPageSuffix() error {
	// write to file:
	// prev next
	err := a.BufferPrevNextPage()

	return err
}

func (a *AssetsPage) BufferPrevNextPage() error {
	path := "./"
	suff := "html"

	prev := format("Page_%d", a.pageNumber-1)
	prevF := format("%s.%s", prev, suff)
	curr := format("Page_%d", a.pageNumber)
	currF := format("%s.%s", curr, suff)
	next := format("Page_%d", a.pageNumber+1)
	nextF := format("%s.%s", next, suff)

	a.Append("<section class=\"nav\">")
	if a.pageNumber > 1 {
		a.AppendLink(prev, (path + prevF))
	}
	a.AppendLink(curr, (path + currF))
	a.AppendLink(next, (path + nextF))
	a.Append("</section>")

	return nil
}

// Adds the relevant custom skin data in markdown.
func (a *AssetsPage) BufferCustomSkins() {
	// this writes to the custom skins stuff and adds the data, in markdown
	a.Append("<section class='skins'>")

	for _, skin := range a.skins {
		a.bufferSkin(skin)
	}

	a.Append("</section")
}

func (a *AssetsPage) bufferSkin(skin nova.CustomSkin) {
	a.Append("<section class='skin'>")

	a.appendHeading(skin)

	a.appendTable(skin)
	a.appendCopyText(skin)
	a.appendDownloadLink(skin)

	a.appendMedia(skin)

	a.Append("</section>")
}

func (a *AssetsPage) appendHeading(skin nova.CustomSkin) {
	heading := a.Fmt.Bold(skin.Name()) + ":"
	a.AppendHeading(2, heading)
	a.Append(skin.FormatCredits(a.Fmt))
	a.AppendEmptyLine()
	a.AppendEmptyLine()
}

func (a *AssetsPage) appendTable(skin nova.CustomSkin) {
	a.Append(skin.ToTable(a.Fmt))
}

func (a *AssetsPage) appendCopyText(skin nova.CustomSkin) {
	a.Append("Copy CSV: <button class='copier' csv='" + skin.ToCSVLine() + "'><img src='/static/svg/copy.svg' alt='Clipboard SVG'></img></button>") // Copy Button, works in conjunction to /static/copy.js in order to copy the skin csv line
	a.AppendEmptyLine()
	a.Append("<code class='csv'>" + skin.ToCSVLine() + "</code>")
	a.AppendEmptyLine()
}

func (a *AssetsPage) appendDownloadLink(skin nova.CustomSkin) {
	a.AppendLink("Download Me", dirs.WwwAssets()+"zips/"+skin.Name()+".zip")
	a.AppendEmptyLine()
}

func (a *AssetsPage) appendMedia(skin nova.CustomSkin) {
	a.Append("<section class=\"media\">")
	if !files.FilesEqual(*skin.Body(), *files.EmptyFile()) {
		a.appendCustomSkinFile(skin.Body())
	}
	if !files.FilesEqual(*skin.ForceArmour(), *files.EmptyFile()) {
		a.appendCustomSkinFile(skin.ForceArmour())
	}

	if !files.FilesEqual(*skin.Drone(), *files.EmptyFile()) {
		a.appendCustomSkinFile(skin.Drone())
	}

	a.AppendEmptyLine()
	a.Append("</section>")
}

func (a *AssetsPage) appendCustomSkinFile(f *files.File) {
	a.AppendEmbed(dirs.WwwSkins()+f.Name(), f.Name())
}

func (a *AssetsPage) writeBuffer() {
	broadcast("Writing to: ", a.FullName())
	// print(a.contentBuffer)
	a.WriteContents()
}
