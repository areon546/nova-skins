package nova

import (
	"errors"
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
		fmt.Println("Buffering skin: ", skin.String())

		if !fileIO.FilesEqual(skin.body, *fileIO.EmptyFile()) {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.body.Name()))
		}
		if !fileIO.FilesEqual(skin.forceArmour, *fileIO.EmptyFile()) {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.forceArmour.Name()))
		}
		if !fileIO.FilesEqual(skin.drone, *fileIO.EmptyFile()) {
			a.AppendMarkdownEmbed(fileIO.ConstructPath(path, "custom_skins", skin.drone.Name()))
		}

		// a.AppendMarkdownLink("Download Me", skin.zip.GetName())

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
		a.bufferCustomSkins()
		a.bufferPageSuffix()

		pages = append(pages, *a)

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
