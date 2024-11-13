package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// ~~~~~~~~~~~~~~~~ File

type File struct {
	filename string
	suffix   string
	relPath  string
	contents []string
	buffer   []string // considering writing to the file with a buffer
	lines    int
}

func NewFileWithSuffix(fn string, suff string) *File {
	return &File{filename: fn, suffix: suff}
}

func NewFile(fn string) *File {
	fn, suff := splitFileName(fn)
	return &File{filename: fn, suffix: suff}
}

func splitFileName(filename string) (name, suffix string) {
	stringSections := strings.Split(filename, ".")
	// print(stringSections)

	if len(stringSections) > 1 {
		suffix = stringSections[len(stringSections)-1]
	}

	for i := 0; i < len(stringSections)-1; i++ {
		name += stringSections[i]
	}

	return
}

func (f *File) GetFileName() string {
	return (fmt.Sprintf("%s.%s", f.filename, f.suffix))
}

func (f *File) readFile() []string {
	data, err := os.ReadFile(f.GetFileName()) // For read access.
	checkError(err)

	oneLine := strings.ReplaceAll(string(data), "\r", "")
	f.contents = strings.Split(oneLine, "\n")
	f.lines = len(f.contents)

	return f.contents
}

func (f *File) isEmpty() bool {
	return len(f.contents) == 0
}

func (f *File) readLine(lineNum int) (output string, err error) {

	if f.isEmpty() {
		f.readFile()
	}

	if lineNum > f.lines {
		return "", err
	}

	output = f.contents[lineNum]
	print(output)

	return
}

func (f *File) writeFile(text string) {
	if err := os.WriteFile(f.GetFileName(), []byte(text), 0666); err != nil {
		log.Fatal(err)
	}
}

func (f *File) writeLines(arr []Stringable) {
	for _, v := range arr {
		f.writeFile(v.toString())
	}
}

func (f *File) appendToFile() { // TODO
	// is supposed to add lines to a file at the end
}

func (f *File) clearFile() {
	f.writeFile("")
}

func (f *File) toString() string {
	return f.GetFileName()
}

// ~~~~~~~~~~~~~~~~~~~~ CSVFile

type CSVFile struct {
	File
	headings []string
	contents [][]string
}

// returns an array of headings and a 2d array of
func readCSV(filePreffix string) (csv CSVFile) {
	file := File{filename: filePreffix, suffix: "csv"}
	// read fileName into CSVFile

	// file := makeFile(fileName)
	fileContents := file.readFile()

	// go through each line in CSV and
	for i, csvCell := range fileContents {
		// print("csv:", csvCell)
		if i == 0 { // adds headings to headings attribute
			csv.headings = strings.Split(csvCell, ",")
		} else { // ads csv items to contents attribute
			csv.contents = append(csv.contents, strings.Split(csvCell, ","))
		}
	}

	return
}

func (c *CSVFile) getIndexOfColumn(header string) (index int) {
	for i, heading := range c.headings {
		if reflect.DeepEqual(heading, header) {
			index = i
		}
	}

	return
}

func (c *CSVFile) printHeaders() {
	print(c.headings)
}

func (c *CSVFile) printContents() {
	for _, v := range c.contents {
		print(v)
	}
}

func (c *CSVFile) Rows() int {
	return len(c.contents)
}

// ~~~~~~~~~~~~~~~~~~~ AssetPage

type AssetsPage struct {
	File
	pageNumber int

	CustomSkin
}

func NewAssetsPage(f File, pageNum int, c CustomSkin) *AssetsPage {
	return &AssetsPage{File: f, pageNumber: pageNum, CustomSkin: c}
}

func (a *AssetsPage) writePagePreffix(pageNumber int) error {
	// write to file:
	// Page #
	// prev next
	a.writeFile(fmt.Sprintf("Page %d", pageNumber))
	err := a.writePrevNextPage(pageNumber)

	return err
}

func (a *AssetsPage) writePrevNextPage(pageNumber int) error {
	path := "../pages/"
	links := ""

	if pageNumber > 1 {

		links += constructMarkdownLink(false, "Page 1", (path + format("Page%d.md", (pageNumber-1))))
	}

	a.writeFile(links)

	return nil
}

// ~~~~~~~~~~~~~~~~~ Asset File

type AssetFile struct {
	CSVFile
}

func (a *AssetFile) determineAssets() (assets []string) {
	// splits out the column in CSV file that refers to assets

	// determines column of asset column
	iOfAssets := a.getIndexOfColumn("assetName")
	iOfFileType := a.getIndexOfColumn("fileType")

	for i, row := range a.contents {

		if i > 0 {
			item := row[iOfAssets]

			if reflect.DeepEqual(row[iOfFileType], "folder") {
			} else {
				item += "." + row[iOfFileType]
			}
			assets = append(assets, item)
		}
	}

	print("assets", assets)

	return
}
