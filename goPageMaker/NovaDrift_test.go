package main

import (
	"reflect"
	"testing"
)

func TestNewCustomSkin(t *testing.T) {
	want := &CustomSkin{name: "", angle: "0", distance: "0"}
	get := NewCustomSkin("", "0", "0")

	assertEquals(t, want, get)

}

func assertEquals(t testing.TB, want, get Stringable) {
	t.Helper()
	if reflect.DeepEqual(want, get) {
		return
	}

	t.Log(want.toString(), get.toString())

	t.Errorf("aaaaaaaaa")
}
