package main

import (
	"fmt"
	"log/slog"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-files/files/zip"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/dirs"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/log"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/processing"
)

var (
	custom_skins_dir string = dirs.Skins()
	compiled_pages   string = dirs.Pages()

	skinsToDisplay = 500000
	skinsPerPage   = 12
)

func main() {
	testing := false
	// testing = !testing

	if testing {
		fmt.Println(cred.GetDiscordUIDs())
		test()
		return
	}

	print("Running")

	log.ClearLogFile()
	log.SetLogger(slog.LevelDebug)
	// zipAllSkins()
	compileSkins()
}

func test() {
	file := files.NewTextFile("../testing.txt")
	file.Append("asdasd")
	file.Append("asdasd")
	file.Append("asdasd")

	fmt.Println(file.Contents())

	file.WriteContents()
}

func print(a ...any) {
	helpers.Broadcast(a...)
}

func compileSkins() {
	log.Info("Compiling Skins")
	// delete the entirety of the pages' folder's contents if present
	files.RemoveAllWithinDirectory(compiled_pages)
	log.Info("Removed from pages directory", "pages", compiled_pages)

	// compiles a list of skins based on the files in the custom skins directory
	skins := processing.GetCustomSkins(files.ReadDirectory(custom_skins_dir))
	log.Info("Read Custom Skin directory")
	print(len(skins), len(skins))

	skinsToDisplay = min(len(skins), skinsToDisplay)
	// the processing package creates a list of skins based on the custom skins csv in the custom skins folder and uses that to create these
	processing.ConstructAssetPages(skins[:skinsToDisplay], skinsPerPage)
}

func zipAllSkins() {
	// zips custom_skins folder
	zip.ZipFolder(custom_skins_dir, dirs.Assets()+"custom_skins")
	// TODO: currently I do not like how the two arguments have the same name, and that is because the function adds a zip at the end
	// solution: make it check if there is a zip or .ZIP in the function already
}
