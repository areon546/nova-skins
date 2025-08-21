package nova

import (
	"testing"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

func TestNewCustomSkin(t *testing.T) {
	want := &CustomSkin{name: "", angle: "0", distance: "0"}
	get := NewCustomSkin("")
	get.AddAngle("0")
	get.AddDistance("0")

	helpers.AssertObjectEquals(t, want, get)
}
