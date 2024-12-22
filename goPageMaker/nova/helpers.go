package nova

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

// helper functions

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func pagesFolder() string {
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

func inSkinsFolder(file string) string {
	return fileIO.ConstructPath("..", "custom_skins", file)
}

func inAssetsFolder(file string) string {
	return fileIO.ConstructPath("..", "assets", file)
}

func AssetsCSVPath() string {
	return inAssetsFolder("assets.csv")
}
