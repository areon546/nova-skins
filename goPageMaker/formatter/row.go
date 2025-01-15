package formatter

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

type row struct {
	cells  []cell
	maxLen int
}

func NewRow(length int) *row {
	return &row{cells: make([]cell, length), maxLen: length}
}

func (r *row) Set(index int, value string) error {
	if index < r.maxLen {
		r.cells[index] = *NewCell(value)
		return nil
	}
	return ErrEndOfRow
}

func (r *row) Get(index int) string {
	return r.cells[index].String()
}

func (r *row) Lengthen(increase int) {
	r.maxLen += increase
}

func (r *row) Size() int {
	return len(r.cells)
}

func (r *row) String() string {
	return helpers.Format("maxLen: %d, cells: %s", r.maxLen, r.cells)
}

func makeRows(rowC, length int) []row {
	rows := make([]row, 0)
	for range rowC {
		rows = append(rows, *NewRow(length))
	}
	return rows
}

type cell struct {
	value string
}

func NewCell(value string) *cell {
	return &cell{value: value}
}

func (c cell) String() string {
	return c.value
}
