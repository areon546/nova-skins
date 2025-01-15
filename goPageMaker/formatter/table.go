package formatter

import (
	"errors"
)

var (
	ErrEndOfRow    = errors.New("fileIO: end of row")
	ErrOutOfBounds = errors.New("fileIO: index out of bounds")
)

type TableConverter func(t table) string

type Table struct{ table }
type table struct {
	headers row
	rows    []row
}

func NewTable(cols, rows int) *Table {
	return &Table{table{headers: *NewRow(cols), rows: makeRows(rows, cols)}}
}
func (t *table) Rows() int {
	return len(t.rows)
}

func (t *table) Cols() int {
	return t.headers.Size()
}

func (t *table) AddRow(r row) {
	t.rows = append(t.rows, r)
}

func (t *table) GetRow(i int) (row, error) {
	if i < 0 || i > t.Rows() {
		return *NewRow(0), ErrOutOfBounds
	}
	return t.rows[i], nil
}

func (t *table) AddCol() {
	t.headers.Lengthen(1)
	for _, row := range t.rows {
		row.Lengthen(1)
	}
}

func (t *table) AddHeader(index int, newHeader string) (err error) {
	err = t.headers.Set(index, newHeader)
	return errors.Join(err, errors.New(":end of headers"))
}
