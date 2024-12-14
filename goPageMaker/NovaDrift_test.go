package main

import (
	"testing"
)

func TestNewCustomSkin(t *testing.T) {
	want := &CustomSkin{name: "", angle: "0", distance: "0"}
	get := NewCustomSkin("", "0", "0")

	assertEquals(t, want, get)

}
