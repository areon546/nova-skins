package processing

import (
	"io/fs"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/go-files/files"
)

// helper functions

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func print(a ...any) {
	helpers.Print(a...)
}

func pagesFolder() string {
	return "../pages/"
}

func inSkinsFolder(filename, filetype string) string {
	s := files.ConstructFilePath("../custom_skins/", filename, filetype)

	print(s)

	return s
}

func inAssetsFolder(file, filetype string) string {
	return files.ConstructFilePath("../assets", file, filetype)
}

func AssetsCSVPath() string {
	return inAssetsFolder("assets", "csv")
}

func openCustomSkin(d fs.DirEntry) *files.File {
	f, _ := files.OpenFile(files.ConstructFilePath("../custom_skins", d.Name(), ""))

	return f
}
