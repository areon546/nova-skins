package helpers

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/areon546/go-helpers"
)

func HandleExcept(err, allowed error) {
	errorAllowed := errors.Is(err, allowed)
	if err != nil {
		if !errorAllowed {
			log.Fatal(err)
		}
	}
}

// already in helpers
// a
//
// a
//
// a
// ~~~~~~~~~~~~~~~~~~~~~~~~~~

// helper functions

func Print(a ...any) { helpers.Print(a...) }

func Printf(s string, a ...any) { helpers.Printf(s, a...); helpers.Print("") }

func Format(s string, a ...any) string { return helpers.Format(s, a...) }

func Search[T any](item T, arr []T) (index int) {
	return helpers.Search(item, arr)
}

func ConvertToInteger(s string) (int, error) {
	return helpers.ConvertToInteger(s)
}

func Handle(err error) {
	helpers.Handle(err)
}

func AssertEquals(t testing.TB, expected, result string) {
	t.Helper()
	if reflect.DeepEqual(expected, result) {
		return
	}

	t.Log(expected, result)

	t.Errorf("Variables are not equal, \nexpected: %s \nresult: %s", expected, result)
}

func AssertObjectEquals(t testing.TB, expected, result fmt.Stringer) {
	t.Helper()
	if reflect.DeepEqual(expected, result) {
		return
	}

	t.Log(expected.String(), result.String())

	t.Errorf("Variables are not equal, \nexpected: %s \nresult: %s", expected, result)
}

func AssertIntEquals(t testing.TB, expected, result int) {
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

// dont think they are necessary honestly but here still
func BytesToString(b []byte) (s string) {
	return string(b)
}
func StringToBytes(s string) (b []byte) {
	return []byte(s)
}
