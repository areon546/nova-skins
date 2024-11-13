package main

import (
	"fmt"
	"io/fs"
)

func main() {
	var cs []CustomSkin

	debug := false
	// debug = !debug

	print("Running")

	if debug {
		print("TEST:")
		runTest()
		return
	}

	// reads what files are in the assets folder
	// assets := readDirectory(skinFolder())
	// assetsAsFiles, _ := filterFiles(assets)

	// print("assets", assets)
	// printf("%s", "abba")
	// for _, v := range assetsAsFiles {
	// 	print(v.toString())
	// }

	skinsCSV := readCSV(skinFolder() + "custom_skins")
	names := skinsCSV.getIndexOfColumn("name")
	angles := skinsCSV.getIndexOfColumn("jet_angle")
	distances := skinsCSV.getIndexOfColumn("jet_distance")
	skins := skinsCSV.getIndexOfColumn("body_artwork")
	forces := skinsCSV.getIndexOfColumn("body_force_armor_artwork")
	drones := skinsCSV.getIndexOfColumn("drone_artwork")

	cs = make([]CustomSkin, skinsCSV.Rows())
	print(cs, skinsCSV.Rows())

	for i, v := range skinsCSV.contents {
		print(i, v, skins, forces, drones)

		name := v[names]
		distance := convertToInteger(v[distances])
		angle := convertToInteger(v[angles])

		c := NewCustomSkin(name, distance, angle).addSkin(v[skins]).addForceA(v[forces]).addDrone(v[drones])

		// print("c, ", c)

		cs = append(cs, *c)
	}

	print(cs)

}

func runTest() {

	testFile := NewFile("file.md")
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
	skinPicture string
	forceArmour string
	drone       string
	distance    int
	angle       int
}

func NewCustomSkin(name string, distance, angle int) *CustomSkin {
	return &CustomSkin{name: name, distance: distance, angle: angle}
}

func (c *CustomSkin) addSkin(s string) *CustomSkin {
	c.skinPicture = s
	return c
}

func (c *CustomSkin) addForceA(s string) *CustomSkin {
	c.forceArmour = s
	return c
}

func (c *CustomSkin) addDrone(s string) *CustomSkin {
	c.drone = s
	return c
}

func (c *CustomSkin) toString() string {
	return c.name
}

func convertCSVLineToCustomSkin(ar []string) {

	return
}
