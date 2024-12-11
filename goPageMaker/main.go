package main

func main() {

	print("Running")

	// reads what files are in the assets folder
	// assets := readDirectory(skinFolder())
	// assetsAsFiles, _ := filterFiles(assets)

	// print("assets", assets)
	// printf("%s", "abba")
	// for _, v := range assetsAsFiles {
	// 	print(v.String())
	// }

	// returns a list of CustomSkins based on whats in the custom_skins folder
	skins := getCustomSkins()

	constructAssetPages(skins[:20])

}

func getPagesFolder() string {
	return "../pages"
}
