package main

import "github.com/areon546/NovaDriftCustomSkins/goPageMaker/nova"

func main() {

	print("Running")
	// delete the entirety of the pages' folder's contents if present

	// returns a list of CustomSkins based on whats in the custom_skins folder
	skins := nova.GetCustomSkins()

	// print(skins)

	nova.ConstructAssetPages(skins[:])

}
