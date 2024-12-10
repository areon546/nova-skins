package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCustomSkin(t *testing.T) {
	want := &CustomSkin{name: "", angle: "0", distance: "0"}
	get := NewCustomSkin("", "0", "0")

	assertEquals(t, want, get)

}

func assertEquals(t testing.TB, want, get fmt.Stringer) {
	t.Helper()
	if reflect.DeepEqual(want, get) {
		return
	}

	t.Log(want.String(), get.String())

	t.Errorf("aaaaaaaaa")
}
