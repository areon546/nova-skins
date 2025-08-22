package processing

import (
	"errors"
	"fmt"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
)

func ConstructAssetPages(skins []nova.CustomSkin) (pages []AssetsPage) {
	broadcast("Making Asset Pages")
	numSkins := len(skins)
	numFiles := numSkins / 10

	// Count number of files expected to create.
	if numSkins%10 != 0 {
		numFiles++
	}

	for i := range numFiles {
		// create a new file
		pageNum := i + 1
		a := NewAssetsPage(pagesFolder(), format("Page_%d.md", pageNum), pageNum)
		_ = a.ClearFile() // don't care about this error

		writeToAssetPage(a, skins, i)
		pages = append(pages, *a)

	}
	return
}

func writeToAssetPage(a *AssetsPage, skins []nova.CustomSkin, i int) {
	fmt.Println("Wrote to page", a)
	a.bufferPagePreffix()
	//
	skinSlice, err := getNextSlice(skins, i)
	helpers.Handle(err)

	a.addCustomSkins(skinSlice)
	a.bufferCustomSkins()
	a.bufferPageSuffix()
	//
	a.writeBuffer()
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
