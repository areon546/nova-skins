package helpers

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/areon546/go-helpers/helpers"
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
	return helpers.StringToInteger(s)
}

func Handle(err error) {
	helpers.Handle(err)
}

func AssertEquals(t testing.TB, expected, result fmt.Stringer) {
	t.Helper()
	helpers.AssertEqualsStringer(t, expected, result)
}

func AssertObjectEquals(t testing.TB, expected, result any) {
	t.Helper()
	helpers.AssertEqualsObject(t, expected, result)
}

func AssertIntEquals(t testing.TB, expected, result int) {
	helpers.AssertEqualsInt(t, expected, result)
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	helpers.AssertError(t, got, want)
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	AssertError(t, err, nil)
}
