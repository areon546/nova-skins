package nova

import (
	"io/fs"
	"reflect"
	"strconv"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/formatter"
)

var (
	ErrMalformedRow CustomSkinError = CustomSkinError{"malformed row"}
)

type CustomSkinError struct {
	name string
}

func (cse CustomSkinError) Error() string {
	return cse.name
}

// ~~~~~~~~~~~~~~~~~ CustomSkin
type CustomSkin struct {
	pictures []fileIO.File
	credit   cred.CreditType

	name        string
	Body        fileIO.File
	ForceArmour fileIO.File
	Drone       fileIO.File
	angle       string
	distance    string

	zip fileIO.ZipFile
}

func NewCustomSkin(name, angle, distance string) (cs *CustomSkin) {
	cs = &CustomSkin{name: name, angle: angle, distance: distance}
	return
}

func (cs *CustomSkin) addBody(f fileIO.File) *CustomSkin {
	cs.Body = f
	return cs
}

func (cs *CustomSkin) addForceA(s fileIO.File) *CustomSkin {
	cs.ForceArmour = s
	return cs
}

func (cs *CustomSkin) addDrone(f fileIO.File) *CustomSkin {
	cs.Drone = f
	return cs
}

func (cs *CustomSkin) addCredits(c cred.CreditType) {
	cs.credit = c
}

func (cs *CustomSkin) addMedia(f fileIO.File) {
	cs.pictures = append(cs.pictures, f)
}

func (cs *CustomSkin) HasZip() bool {
	return reflect.DeepEqual(&cs.zip, (&fileIO.ZipFile{}))
}

// TODO This should use the fs.DirEntires to generate a zip file for the individual skin
func (cs *CustomSkin) generateZipFile() {

	path := fileIO.ConstructPath("..", "assets/zips", cs.name)
	cs.zip = *fileIO.NewZipFile(path)

	body, fA, drone := cs.getBody_FA_Drone()

	if body != "" {
		cs.zip.AddZipFile(body, cs.Body)
	}

	if fA != "" {
		cs.zip.AddZipFile(fA, cs.ForceArmour)
	}

	if drone != "" {
		cs.zip.AddZipFile(drone, cs.Drone)
	}
	// // helpers.Print(cs.forceArmour.BufferToString())
	cs.zip.WriteToZipFile()

	return
}

func (cs CustomSkin) String() string {
	return cs.name
}

func openCustomSkin(d fs.DirEntry) *fileIO.File {
	return fileIO.OpenFile("../custom_skins/", d)
}

func EmptyCustomSkin() *CustomSkin {
	return &CustomSkin{}
}

func (cs CustomSkin) getBody_FA_Drone() (body, fA, drone string) {
	body, fA, drone = "", "", ""

	if !fileIO.FilesEqual(cs.Body, *fileIO.EmptyFile()) {
		body = cs.Body.Name()
	}

	if !fileIO.FilesEqual(cs.ForceArmour, *fileIO.EmptyFile()) {
		fA = cs.ForceArmour.Name()
	}

	if !fileIO.FilesEqual(cs.Drone, *fileIO.EmptyFile()) {
		drone = cs.Drone.Name()
	}

	return
}

func (cs *CustomSkin) ToCSVLine() string {
	body, fA, drone := cs.getBody_FA_Drone()
	return format("%s,%s,%s,%s,%s,%s", cs.name, body, fA, drone, cs.getAngle(), cs.getDistance())
}

func (cs *CustomSkin) ToTable(fmt formatter.Formatter) string {
	body, fA, drone := cs.getBody_FA_Drone()

	t := formatter.NewTable(2, 0)
	bodyRow := formatter.NewRow(2)
	bodyRow.Set(0, "Body:")
	bodyRow.Set(1, body)
	t.AddRow(*bodyRow)
	faRow := formatter.NewRow(2)
	faRow.Set(0, "Fource Armour:")
	faRow.Set(1, fA)
	t.AddRow(*faRow)
	droneRow := formatter.NewRow(2)
	droneRow.Set(0, "Drone:")
	droneRow.Set(1, drone)
	t.AddRow(*droneRow)
	angleRow := formatter.NewRow(2)
	angleRow.Set(0, "Angle:")
	angleRow.Set(1, cs.getAngle())
	t.AddRow(*angleRow)
	distanceRow := formatter.NewRow(2)
	distanceRow.Set(0, "Distance:")
	distanceRow.Set(1, cs.getDistance())
	t.AddRow(*distanceRow)

	// return format("| -- | --- | \n| Body:| %s| \n| ForceArmour:| %s| \n| Drone:| %s| \n| Angle:| %s| \n| Distance:| %s| \n", body, fA, drone, cs.getAngle(), cs.getDistance())

	return fmt.FormatTable(*t)
}

func (c *CustomSkin) getAngle() string {
	// try to convert s to an integer, if it fails, return nothing
	_, err := strconv.Atoi(c.angle)
	if err != nil {
		return ""
	} else {
		return c.angle
	}
}

func (c *CustomSkin) getDistance() string {
	// try to convert to an integer
	_, err := strconv.Atoi(c.distance)
	if err != nil {
		return ""
	} else {
		return c.distance
	}
}

func (c *CustomSkin) FormatCredits(fmt formatter.Formatter) string {
	if c.credit == nil {
		return ""
	}
	return fmt.FormatLink(c.credit.ConstructName(), c.credit.ConstructLink())
}
