package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{0, 1, 2}

	t.Run("Search when value missing", func(t *testing.T) {
		want := -1
		get := search(35, arr)

		assertEqualsInt(t, want, get)
	})

	t.Run("Search when value in arr", func(t *testing.T) {
		assertEqualsInt(t, 1, search(1, arr))
	})
}

func TestConvertToInteger(t *testing.T) {
	t.Run("Convert Valid String to Integer", func(t *testing.T) {
		want := 1
		get, err := convertToInteger("1")

		assertNoError(t, err)
		assertEqualsInt(t, want, get)
	})

	t.Run("Convert Float String to Integer", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, got := convertToInteger("1.0")

		assertError(t, got, want)
	})
	t.Run("Convert Invalid String to Integer", func(t *testing.T) {
		want := strconv.ErrSyntax
		_, got := convertToInteger("Abba")

		assertError(t, got, want)
	})
}

func assertEquals(t testing.TB, expected, result fmt.Stringer) {
	t.Helper()
	if reflect.DeepEqual(expected, result) {
		return
	}

	t.Log(expected.String(), result.String())

	t.Errorf("Variables are not equal, \nexpected: %s \nresult: %s", expected, result)
}

func assertEqualsInt(t testing.TB, expected, result int) {
	t.Helper()

	if expected == result {
		return
	}

	t.Log(expected, result)

	t.Errorf("Integers are not equal. \nexpected: %d \nresult: %d", expected, result)

}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	// log.Fatalf("\nexpected %q \ngot %q", want, got)
	if !errors.Is(got, want) {
		t.Fatalf("got error %q want %q", got, want)
	}

}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	assertError(t, err, nil)
}
