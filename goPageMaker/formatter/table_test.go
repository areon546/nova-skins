package formatter

import (
	"testing"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

func TestNewTable(t *testing.T) {
	cols, rows := 1, 5
	want := &Table{table{headers: *NewRow(cols), rows: makeRows(rows, cols)}}
	get := NewTable(cols, rows)

	helpers.AssertObjectEquals(t, want, get)
}

func TestNewRow(t *testing.T) {
	l := 5
	want := &row{cells: make([]cell, l), maxLen: l}
	get := NewRow(l)

	helpers.AssertObjectEquals(t, want, get)
}
