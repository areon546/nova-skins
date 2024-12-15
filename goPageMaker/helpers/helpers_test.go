package helpers

import (
	"strconv"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{0, 1, 2}

	t.Run("Search when value missing", func(t *testing.T) {
		want := -1
		get := search(35, arr)

		AssertEqualsInt(t, want, get)
	})

	t.Run("Search when value in arr", func(t *testing.T) {
		AssertEqualsInt(t, 1, search(1, arr))
	})
}

func TestConvertToInteger(t *testing.T) {
	t.Run("Convert Valid String to Integer", func(t *testing.T) {
		want := 1
		get, err := convertToInteger("1")

		AssertNoError(t, err)
		AssertEqualsInt(t, want, get)
	})

	t.Run("Convert Float String to Integer", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, got := convertToInteger("1.0")

		AssertError(t, got, want)
	})
	t.Run("Convert Invalid String to Integer", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, got := convertToInteger("Abba")

		AssertError(t, got, want)
	})
}
