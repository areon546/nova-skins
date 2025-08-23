package nova

import (
	"reflect"
	"strconv"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-files/files/zip"
	"github.com/areon546/go-files/formatter"
	"github.com/areon546/go-files/table"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/dirs"
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/log"
)

var (
	emptySkinFile = files.EmptyFile().Name()

	ErrMalformedRow CustomSkinError = CustomSkinError{"malformed row"}

	defaultSkinNames   map[string]string = map[string]string{"body": emptySkinFile, "forceArmour": emptySkinFile, "drone": emptySkinFile, "angle": "", "distance": ""}
	missingBody        string            = defaultSkinNames["body"]
	missingForceArmour string            = defaultSkinNames["forceArmour"]
	missingDrone       string            = defaultSkinNames["drone"]
	missingAngle       string            = defaultSkinNames["angle"]

	missingDistance string = defaultSkinNames["distance"]

	missingCredits string = ""
)

type CustomSkinError struct {
	name string
}

func (err CustomSkinError) Error() string {
	return err.name
}

// ~~~~~~~~~~~~~~~~~ CustomSkin
type CustomSkin struct {
	pictures []files.File
	credit   []cred.CreditType

	name        string
	body        files.File
	forceArmour files.File
	drone       files.File
	angle       string
	distance    string

	zip zip.ZipFile
}

func NewCustomSkin(name string) (cs *CustomSkin) {
	cs = &CustomSkin{name: name, credit: []cred.CreditType{}}
	return
}

// ~~~ Setters
func (cs *CustomSkin) AddBody(f files.File) {
	cs.body = f
}

func (cs *CustomSkin) AddForceA(s files.File) {
	cs.forceArmour = s
}

func (cs *CustomSkin) AddDrone(f files.File) {
	cs.drone = f
}

func (cs *CustomSkin) AddAngle(s string) {
	cs.angle = s
}

func (cs *CustomSkin) AddDistance(s string) {
	cs.distance = s
}

func (cs *CustomSkin) AddCredits(c cred.CreditType) {
	cs.credit = append(cs.credit, c)
}

func (cs *CustomSkin) AddMedia(f files.File) {
	cs.pictures = append(cs.pictures, f)
}

// ~~~ Getters
// func (cs CustomSkin) String() string {
// 	return cs.name
// }

func (cs CustomSkin) Name() string {
	return cs.name
}

func (cs *CustomSkin) ToCSVLine() string {
	body, fA, drone := cs.getBody_FA_Drone()
	return format("%s,%s,%s,%s,%s,%s", cs.name, body, fA, drone, cs.Angle(), cs.Distance())
}

func (cs CustomSkin) getBody_FA_Drone() (body, fA, drone string) {
	body = cs.body.Name()
	fA = cs.forceArmour.Name()
	drone = cs.drone.Name()

	return
}

func (c *CustomSkin) Body() *files.File {
	return &c.body
}

func (c *CustomSkin) ForceArmour() *files.File {
	return &c.forceArmour
}

func (c *CustomSkin) Drone() *files.File {
	return &c.drone
}

func (c *CustomSkin) Angle() string {
	// try to convert s to an integer, if it fails, return nothing
	_, err := strconv.Atoi(c.angle)
	if err != nil {
		return missingAngle
	} else {
		return c.angle
	}
}

func (c *CustomSkin) Distance() string {
	// try to convert to an integer
	_, err := strconv.Atoi(c.distance)
	if err != nil {
		return missingDistance
	} else {
		return c.distance
	}
}

func (cs *CustomSkin) HasZip() bool {
	return reflect.DeepEqual(&cs.zip, (&zip.ZipFile{}))
}

func (skin *CustomSkin) Zip() *zip.ZipFile {
	return &skin.zip
}

// Processing

// TODO: This should use the fs.DirEntires to generate a zip file for the individual skin
func (cs *CustomSkin) GenerateZipFile() {
	path := dirs.AssetsFolder() + "zips/" + cs.name
	cs.zip = *zip.NewZipFile(path)

	broadcast("Generating ZIP: ", cs.Name())
	body, fA, drone := cs.getBody_FA_Drone()

	if body != missingBody {
		cs.zip.AddZipFile(body, cs.Body())
	}

	if fA != missingForceArmour {
		cs.zip.AddZipFile(fA, cs.ForceArmour())
	}

	if drone != missingDrone {
		cs.zip.AddZipFile(drone, cs.Drone())
	}

	cs.zip.WriteAndClose()
	log.Debug("Wrote zip file", "skin", cs.name)
}

func (c *CustomSkin) FormatCredits(fmt formatter.Formatter) string {
	if len(c.credit) == 0 {
		return missingCredits
	}

	var credits string
	for _, credit := range c.credit {
		credits += fmt.Link(credit.ConstructName(), credit.ConstructLink())
	}

	return credits
}

func (cs *CustomSkin) ToTable(fmt formatter.Formatter) string {
	t := table.NewTable(2)
	bodyRow := table.NewRow(2)
	bodyRow.Set(0, "Body:")
	bodyRow.Set(1, cs.Body().Name())
	t.AddRecord(bodyRow)
	faRow := table.NewRow(2)
	faRow.Set(0, "Force Armour:")
	faRow.Set(1, cs.ForceArmour().Name())
	t.AddRecord(faRow)
	droneRow := table.NewRow(2)
	droneRow.Set(0, "Drone:")
	droneRow.Set(1, cs.Drone().Name())
	t.AddRecord(droneRow)
	angleRow := table.NewRow(2)
	angleRow.Set(0, "Angle:")
	angleRow.Set(1, cs.Angle())
	t.AddRecord(angleRow)
	distanceRow := table.NewRow(2)
	distanceRow.Set(0, "Distance:")
	distanceRow.Set(1, cs.Distance())
	t.AddRecord(distanceRow)

	// return format("| -- | --- | \n| Body:| %s| \n| ForceArmour:| %s| \n| Drone:| %s| \n| Angle:| %s| \n| Distance:| %s| \n", body, fA, drone, cs.getAngle(), cs.getDistance())

	return fmt.Table(*t)
}

// Not class related

func EmptyCustomSkin() *CustomSkin {
	return &CustomSkin{}
}
