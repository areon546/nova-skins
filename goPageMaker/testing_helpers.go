package main

import (
	"fmt"
	"reflect"
	"testing"
)

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
