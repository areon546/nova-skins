package main

import (
	"fmt"
	"io/fs"
)

// tell program page or have it have a csv of pages
// use csv and have like 10 max per page
// have it then use assets to place into page

func main() {
	debug := false
	// debug = !debug

	print("Running")

	if debug {
		print("TEST:")
		runTest()
		return
	}

	// reads what files are in the assets folder
	assets := readDirectory(skinFolder())
	assetsAsFiles, _ := filterFiles(assets)

	print("assets", assets)
	printf("%s", "abba")

	skinsCSV := readCSV(skinFolder() + "custom_skins")
	skinsCSV.printHeaders()

	for _, v := range assetsAsFiles {
		print(v.toString())
	}

	// for _, folder := range folders {
	// 	print(folder)
	// 	files, _ := filterFiles(readDirectory(skinFolder() + folder.Name()))
	// 	print(files)

	// 	c := CustomSkin{name: folder.Name()}
	// 	print(c.toString())
	// }

}

func runTest() {

	testFile := createFile("file.md", "md")
	fmt.Print(testFile.readLine(1))

	return
}

func checkNewAssets(preExistingAssets []string, assets []fs.DirEntry, newAssets []bool) {
	// loop through assets and determine if any assets have been added

	for _, v := range assets {
		// loopts through assets

		location := search(v.Name(), preExistingAssets)
		if location >= 0 {

			print("inArray", v.Name())
		}

	}

	// for i, v := range preExistingAssets {

	// }
}

type CustomSkin struct {
	// pictures []File

	name        string
	credit      string
	skinPicture File
	forceArmour File
	drone       File
	thrust      int
	angle       int
}

func (c CustomSkin) toString() string {
	return c.name
}
