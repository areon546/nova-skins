package main

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
)

func main() {
	testing := false
	// testing = !testing

	if testing {
		print("Testing")

		test()

		return
	}

	print("Running")

	// zips custom_skins folder
	// fileIO.ZipFolder("../custom_skins", "../custom_skins")

	// delete the entirety of the pages' folder's contents if present
	fileIO.RemoveAllWithinDirectory(nova.Pages)

	// the nova package creates a list of skins based on the custom skins csv in the custom skins folder and uses that to create these
	nova.ConstructAssetPages()

}

func test() {
	fileIO.ZipFolder("../custom_skins", "../../../assets")
}

func print(a ...any) {
	helpers.Print(a...)
}
