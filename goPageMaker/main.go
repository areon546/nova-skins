package main

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"
)

func main() {
	testing := false

	if testing {
		print("Testing")

		test()

		return
	}

	print("Running")

	// zips custom_skins folder
	fileIO.ZipFolder("../custom_skins", "../custom_skins")

	// delete the entirety of the pages' folder's contents if present

	// returns a list of CustomSkins based on whats in the custom_skins folder
	print("Compiling Skins")
	skins := nova.GetCustomSkins(fileIO.ReadDirectory("../custom_skins"))

	// print(skins)
	_ = nova.ConstructZipFiles(skins) // zipFiles

	nova.ConstructAssetPages(skins)

}

func test() {
	fs := fileIO.ReadDirectory("../custom_skins")
	for _, file := range fs {
		print(file.Name())
	}
}

func print(a ...any) {
	helpers.Print(a...)
}
