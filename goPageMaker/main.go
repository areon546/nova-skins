package main

func main() {
	var skins []CustomSkin

	print("Running")

	// reads what files are in the assets folder
	// assets := readDirectory(skinFolder())
	// assetsAsFiles, _ := filterFiles(assets)

	// print("assets", assets)
	// printf("%s", "abba")
	// for _, v := range assetsAsFiles {
	// 	print(v.toString())
	// }

	skinsData := readCSV(skinFolder() + "custom_skins")
	names := skinsData.getIndexOfColumn("name")
	angles := skinsData.getIndexOfColumn("jet_angle")
	distances := skinsData.getIndexOfColumn("jet_distance")
	body := skinsData.getIndexOfColumn("body_artwork")
	forces := skinsData.getIndexOfColumn("body_force_armor_artwork")
	drones := skinsData.getIndexOfColumn("drone_artwork")

	skins = make([]CustomSkin, skinsData.Rows())
	print(skins, skinsData.Rows())

	for i, v := range skinsData.contents {
		print(i, v, body, forces, drones)

		name := v[names]
		distance := convertToInteger(v[distances])
		angle := convertToInteger(v[angles])

		skin := NewCustomSkin(name, distance, angle).addSkin(v[body]).addForceA(v[forces]).addDrone(v[drones])

		// print("c, ", c)

		skins[i] = *skin

		print(skin.toString())
	}

	print(skins)

	a := NewAssetsPage(constructPath("", getPagesFolder(), "test"), 0, "")

	a.bufferPagePreffix()
	a.bufferPrevNextPage()
	a.addCustomSkins(skins)
	a.bufferCustomSkins()
	a.bufferPrevNextPage()

	a.writeBuffer()

}

func getPagesFolder() string {
	return "../pages"
}
