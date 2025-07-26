package nova

import (
	"reflect"
	"strconv"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-files/files/zip"
	"github.com/areon546/go-files/formatter"
	"github.com/areon546/go-files/table"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/cred"
)

var (
	emptySkin = files.EmptyFile().Name()

	ErrMalformedRow CustomSkinError = CustomSkinError{"malformed row"}

	defaultFileNames   map[string]string = map[string]string{"body": emptySkin, "forceArmour": emptySkin, "drone": emptySkin, "angle": "", "distance": ""}
	missingBody        string            = defaultFileNames["body"]
	missingForceArmour string            = defaultFileNames["forceArmour"]
	missingDrone       string            = defaultFileNames["drone"]
	missingAngle       string            = defaultFileNames["angle"]

	missingDistance string = defaultFileNames["distance"]

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
	credit   cred.CreditType

	name        string
	body        files.File
	forceArmour files.File
	drone       files.File
	angle       string
	distance    string

	zip zip.ZipFile
}

func NewCustomSkin(name, angle, distance string) (cs *CustomSkin) {
	cs = &CustomSkin{name: name, angle: angle, distance: distance}
	return
}

// ~~~ Setters
func (cs *CustomSkin) AddBody(f files.File) *CustomSkin {
	cs.body = f
	return cs
}

func (cs *CustomSkin) AddForceA(s files.File) *CustomSkin {
	cs.forceArmour = s
	return cs
}

func (cs *CustomSkin) AddDrone(f files.File) *CustomSkin {
	cs.drone = f
	return cs
}

func (cs *CustomSkin) AddCredits(c cred.CreditType) {
	cs.credit = c
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
	path := "../assets/zips/" + cs.name
	cs.zip = *zip.NewZipFile(path)

	print("Generating ZIP: ", cs.Name())
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
}

func (c *CustomSkin) FormatCredits(fmt formatter.Formatter) string {
	if c.credit == nil {
		return missingCredits
	}
	return fmt.FormatLink(c.credit.ConstructName(), c.credit.ConstructLink())
}

func (cs *CustomSkin) ToTable(fmt formatter.Formatter) string {
	t := table.NewTable(2, 0, true)
	bodyRow := table.NewRow(2)
	bodyRow.Set(0, "Body:")
	bodyRow.Set(1, cs.Body().Name())
	t.AddRow(*bodyRow)
	faRow := table.NewRow(2)
	faRow.Set(0, "Fource Armour:")
	faRow.Set(1, cs.ForceArmour().Name())
	t.AddRow(*faRow)
	droneRow := table.NewRow(2)
	droneRow.Set(0, "Drone:")
	droneRow.Set(1, cs.Drone().Name())
	t.AddRow(*droneRow)
	angleRow := table.NewRow(2)
	angleRow.Set(0, "Angle:")
	angleRow.Set(1, cs.Angle())
	t.AddRow(*angleRow)
	distanceRow := table.NewRow(2)
	distanceRow.Set(0, "Distance:")
	distanceRow.Set(1, cs.Distance())
	t.AddRow(*distanceRow)

	// return format("| -- | --- | \n| Body:| %s| \n| ForceArmour:| %s| \n| Drone:| %s| \n| Angle:| %s| \n| Distance:| %s| \n", body, fA, drone, cs.getAngle(), cs.getDistance())

	return fmt.FormatTable(*t, false)
}

// Not class related

func EmptyCustomSkin() *CustomSkin {
	return &CustomSkin{}
}
