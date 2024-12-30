package fileIO

import (
	"reflect"
	"strings"
)

// ~~~~~~~~~~~~~~~~~~~~ CSVFile
type CSVFile struct {
	TextFile
	headings []string
	contents [][]string
}

// returns an array of headings and a 2d array of
func ReadCSV(fileName string) (csv CSVFile) {
	file := NewTextFileWithSuffix("", fileName, "csv")
	// read fileName into CSVFile

	// file := makeFile(fileName)
	fileContents := file.ReadFile()

	// go through each line in CSV and
	for i, csvCell := range fileContents {
		// print("csv:", csvCell)
		if i == 0 { // adds headings to headings attribute
			csv.headings = strings.Split(csvCell, ",")
		} else { // ads csv items to contents attribute

			// check if the string is empty, if so skip
			if reflect.DeepEqual(csvCell, "") {
				continue
			}

			csv.contents = append(csv.contents, strings.Split(csvCell, ","))
		}
	}

	return
}

func (c *CSVFile) GetIndexOfColumn(header string) (index int) {
	for i, heading := range c.headings {
		if reflect.DeepEqual(heading, header) {
			index = i
		}
	}

	return
}

func (c *CSVFile) GetRow(i int) string { // TODO make more efficient
	return strings.Join(c.contents[i], ",")
	// return c.contentBuffer[i+1] // this is buggy, fix
}

func (c *CSVFile) GetCell(row, col int) string {
	return c.contents[row][col]
}

func (c *CSVFile) NumHeaders() int {
	return len(c.headings)
}

func (c *CSVFile) Rows() int {
	return len(c.contents)
}

func (c *CSVFile) GetContents() [][]string {
	return c.contents
}
