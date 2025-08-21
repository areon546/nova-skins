package nova

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/go-files/files"
)

// helper functions

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func broadcast(a ...any) {
	helpers.Broadcast(a...)
}

func inAssetsFolder(file string) string {
	return files.ConstructFilePath("../assets", file, "")
}

func AssetsCSVPath() string {
	return inAssetsFolder("assets.csv")
}
