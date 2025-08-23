package processing

import (
	"io/fs"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/dirs"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/log"
	"github.com/areon546/go-files/files"
)

// helper functions

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func broadcast(a ...any) {
	helpers.Broadcast(a...)

	logString := format("Broadcasting: %s", a)
	log.Info(logString)
}

func inSkinsFolder(filename string) string {
	return dirs.SkinsFolder() + filename
}

func inAssetsFolder(filename string) string {
	return "../assets/" + filename
}

func AssetsCSVPath() string {
	return inAssetsFolder("assets.csv")
}

func openCustomSkin(d fs.DirEntry) *files.File {
	f, _ := files.OpenFile(dirs.SkinsFolder() + d.Name())

	return f
}
