package main

func main() {

	print("Running")
	// delete the entirety of the pages' folder's contents if present

	// returns a list of CustomSkins based on whats in the custom_skins folder
	skins := getCustomSkins()

	// print(skins)

	constructAssetPages(skins[:])

}

func getPagesFolder() string {
	return "../pages"

	// reads what files are in the assets folder
	// assets := readDirectory(skinFolder())
	// assetsAsFiles, _ := filterFiles(assets)

	// print("assets", assets)
	// printf("%s", "abba")
	// for _, v := range assetsAsFiles {
	// 	print(v.String())
	// }

}

func skinFolder() string {
	return "../custom_skins/"
}
