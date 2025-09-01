package processing

import (
	"errors"
	"fmt"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/dirs"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
)

func ConstructAssetPages(skins []nova.CustomSkin, skinsPerPage int) (pages []AssetsPage) {
	broadcast("Making Asset Pages")
	numSkins := len(skins)
	numFiles := numSkins / skinsPerPage

	// Count number of files expected to create.
	if numSkins%skinsPerPage != 0 {
		numFiles++
	}

	for i := range numFiles {
		// create a new file
		pageNum := i + 1
		a := NewAssetsPage(dirs.Pages(), format("Page_%d.md", pageNum), pageNum)
		_ = a.ClearFile() // don't care about this error

		writeToAssetPage(a, skins, i, skinsPerPage)
		pages = append(pages, *a)

	}
	return
}

func writeToAssetPage(a *AssetsPage, skins []nova.CustomSkin, i, l int) {
	fmt.Println("Wrote to page", a)
	a.BufferPagePreffix()
	//
	skinSlice, err := getNextSlice(skins, i, l)
	helpers.Handle(err)

	a.AddCustomSkins(skinSlice)
	a.BufferCustomSkins()
	a.BufferPageSuffix()
	//
	a.writeBuffer()
}

func getNextSlice(skins []nova.CustomSkin, i, l int) (subset []nova.CustomSkin, err error) {
	numSkins := len(skins)

	if i < 0 || i > (len(skins)/l+1) {
		err = errors.New("index out of bounds for CustomSkins array")
	}

	min, max := i*l, (i+1)*l
	fmt.Println(min, max)

	if max > numSkins {
		max = numSkins
	}

	return skins[min:max], err
}
