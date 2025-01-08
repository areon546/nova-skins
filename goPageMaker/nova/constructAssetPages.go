package nova

import (
	"errors"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

func ConstructAssetPages() (pages []AssetsPage) {
	helpers.Print("Making Files")
	numSkins := len(Skins)
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

		skinSlice, err := getNextSlice(Skins, i)
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
