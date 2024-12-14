package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{0, 1, 2}

	want := -1
	get := search(35, arr)

	assertEqualsInt(t, want, get)
}
