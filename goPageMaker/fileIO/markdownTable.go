package fileIO

type markdownTable struct {
	headers []string
	table   [][]string
}

func CreateNewTable() *markdownTable {
	return CreateNewTableWith(0, 0)
}

func CreateNewTableWith(rows, cols int) *markdownTable {
	return &markdownTable{headers: make([]string, cols), table: make([][]string, rows-1)}
}

func (m *markdownTable) Rows() int {
	return len(m.table) + 1
}

func (m *markdownTable) Cols() int {
	return len(m.headers)
}

func (m *markdownTable) String() string {
	s := ""

	headers := constructRow(m.headers)
	headerDecleration := markdownHeaderDeclarationRow(len(m.headers))

	s += headers + headerDecleration

	return s
}

func (m *markdownTable) AddHeader(newHeader string) {
	m.headers = append(m.headers, newHeader)
	print(len(m.headers))
}

func constructRow(s []string) (row string) {
	for _, v := range s {
		row += constructCell(v)
	}

	row += " |\n" // since a row needs to start and end with a | on a new line to be valid

	return
}

func constructCell(cell string) string {
	return " | " + cell
}

func markdownHeaderDeclarationRow(length int) string {
	headerDecleration := make([]string, 0, length)

	for _ = range length {
		headerDecleration = append(headerDecleration, "~~~")
	}

	return constructRow(headerDecleration)
}
