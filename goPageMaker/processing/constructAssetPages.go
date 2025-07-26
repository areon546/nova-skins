package processing

import (
	"errors"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
)

func ConstructAssetPages(skins []nova.CustomSkin) (pages []AssetsPage) {
	helpers.Print("Making Asset Pages")
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
		a := NewAssetsPage("", pagesFolder()+format("Page_%d.md", pageNum), pageNum)

		a.bufferPagePreffix()

		skinSlice, err := getNextSlice(skins, i)
		helpers.Handle(err)

		a.addCustomSkins(skinSlice)
		a.bufferCustomSkins()
		a.bufferPageSuffix()

		pages = append(pages, *a)

		a.writeBuffer()
	}
	return
}

func getNextSlice(skins []nova.CustomSkin, i int) (subset []nova.CustomSkin, err error) {
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
