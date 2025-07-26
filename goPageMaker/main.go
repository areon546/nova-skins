package main

import (
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-files/files/zip"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/processing"
)

var (
	custom_skins_dir string = "../custom_skins/"
	compiled_pages   string = "../compiled/pages"
)

func main() {
	testing := false
	// testing = !testing

	if testing {
		print("Testing")
		test()
	}

	print("Running")

	// zipAllSkins()
	compileSkins()
}

func test() {
	// zip.ZipFolderO("../pages", "../main")
	helpers.Print(files.ConstructFilePath(compiled_pages, "asd", ""))

	txt := files.NewTextFile("../404.md")
	print(txt.Name())

	result := txt.ReadFile()

	helpers.Print(result)
}

func print(a ...any) {
	helpers.Print(a...)
}

func compileSkins() {
	// delete the entirety of the pages' folder's contents if present
	files.RemoveAllWithinDirectory(compiled_pages)

	// compiles a list of skins based on the files in the custom skins directory
	skins := processing.GetCustomSkins(files.ReadDirectory(custom_skins_dir))

	// the processing package creates a list of skins based on the custom skins csv in the custom skins folder and uses that to create these
	processing.ConstructAssetPages(skins)
}

func zipAllSkins() {
	// zips custom_skins folder
	zip.ZipFolder(custom_skins_dir, "../custom_skins")
	// TODO: currently I do not like how the two arguments have the same name, and that is because the function adds a zip at the end
	// solution: make it check if there is a zip or .ZIP in the function already
}
