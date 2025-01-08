package nova

import (
	"io/fs"
	"reflect"
	"strconv"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"
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

func (cs *CustomSkin) ToTable() string {
	body, fA, drone := cs.getBody_FA_Drone()

	return format("| -- | --- | \n| Body:| %s| \n| ForceArmour:| %s| \n| Drone:| %s| \n| Angle:| %s| \n| Distance:| %s| \n", body, fA, drone, cs.getAngle(), cs.getDistance())
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

func (c *CustomSkin) FormatCredits() string {
	if c.credit == nil {
		return ""
	}
	return fileIO.ConstructMarkDownLink(false, c.credit.ConstructName(), c.credit.ConstructLink())
}
