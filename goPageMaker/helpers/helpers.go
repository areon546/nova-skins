package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/areon546/go-helpers"
)

// helper functions

func Print(a ...any) { helpers.Print(a...) }

func Printf(s string, a ...any) { helpers.Printf(s, a...) }

func Format(s string, a ...any) string { return helpers.Format(s, a...) }

func AssetsCSVPath() string {
	return "assets.csv"
}

func search[T any](item T, arr []T) (index int) {
	index = -1
	for i, v := range arr {
		if reflect.DeepEqual(v, item) {
			index = i
		}
	}
	return index
}

func convertToInteger(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return i, err
	}
	return i, err
}

func Handle(err error) {
	// helpers.Handle(err)
	helpers.CheckError(err)
}
func AssertEquals(t testing.TB, expected, result fmt.Stringer) {
	t.Helper()
	if reflect.DeepEqual(expected, result) {
		return
	}

	t.Log(expected.String(), result.String())

	t.Errorf("Variables are not equal, \nexpected: %s \nresult: %s", expected, result)
}

func AssertEqualsInt(t testing.TB, expected, result int) {
	t.Helper()

	if expected == result {
		return
	}

	t.Log(expected, result)

	t.Errorf("Integers are not equal. \nexpected: %d \nresult: %d", expected, result)

}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	// log.Fatalf("\nexpected %q \ngot %q", want, got)
	if !errors.Is(got, want) {
		t.Fatalf("got error %q want %q", got, want)
	}

}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	AssertError(t, err, nil)
}
