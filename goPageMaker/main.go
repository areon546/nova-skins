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
	fileIO.ZipFolder("../custom_skins", "../custom_skins")

	// delete the entirety of the pages' folder's contents if present
	fileIO.RemoveAllWithinDirectory(nova.Pages)

	custom_skin_dir := fileIO.ReadDirectory("../custom_skins")
	// returns a list of CustomSkins based on whats in the custom_skins folder
	print("Compiling Skins")
	skins := nova.GetCustomSkins(custom_skin_dir)

	print("Making Files")
	nova.ConstructAssetPages(skins[:])

}

func test() {
	fileIO.ZipFolder("../custom_skins", "../../../assets")
}

func print(a ...any) {
	helpers.Print(a...)
}
