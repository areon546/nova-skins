package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// ~~~~~~~~~~~~~~~~ File

type File struct { // TODO UPDATE TO USING THE BUFFER
	filename      string
	suffix        string
	relPath       string
	contentBuffer []string
	lines         int
	hasBeenRead   bool
}

func NewFileWithSuffix(fn string, suff string, path string) *File {
	return &File{filename: fn, suffix: suff, relPath: path, hasBeenRead: false}
}

func NewFile(fn string) *File {
	fn, suff := splitFileName(fn)
	return &File{filename: fn, suffix: suff, hasBeenRead: false}
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
	if !f.hasBeenRead {
		data, err := os.ReadFile(f.GetFileName()) // For read access.
		checkError(err)

		oneLine := strings.ReplaceAll(string(data), "\r", "")
		f.contentBuffer = strings.Split(oneLine, "\n")
		f.lines = len(f.contentBuffer)
	}
	return f.contentBuffer
}

func (f *File) isEmpty() bool {
	return len(f.contentBuffer) == 0
}

func (f *File) readLine(lineNum int) (output string, err error) {
	lineNum -= 1 // converted to index notation

	if f.isEmpty() {
		f.readFile()
	}

	if lineNum > f.lines {
		return "", err
	}

	output = f.contentBuffer[lineNum]
	print(output)

	return
}

func (f *File) writeFile() {
	if err := os.WriteFile(f.GetFileName(), []byte(f.bufferToString()), 0666); err != nil {
		log.Fatal(err)
	}
}

func (f *File) appendLines(arr []string) {

	// f.contentBuffer = append(f.contentBuffer, arr...)
	for _, v := range arr {
		f.contentBuffer = append(f.contentBuffer, v)
	}
}

func (f *File) append(s string) {
	f.appendLine(s, len(f.contentBuffer), true)
}

func (f *File) appendNewLine() {
	f.append("")
}

func (f *File) bufferLines(arr []string) {
	s := ""
	print(s)
	f.contentBuffer = make([]string, len(arr))
	for i, v := range arr {
		f.contentBuffer[i] = v
	}
}

func (f *File) clearFile() {
	if err := os.WriteFile(f.GetFileName(), make([]byte, 0), 0666); err != nil {
		log.Fatal(err)
	}
}

func (f *File) String() string {
	return f.GetFileName()
}

func (f *File) bufferToString() string {
	s := ""
	for _, v := range f.contentBuffer {
		s += v
	}

	return s
}

func (f *File) appendLine(s string, i int, nl bool) {

	for i >= len(f.contentBuffer) {
		f.contentBuffer = append(f.contentBuffer, "")
	}

	if nl {
		s += "\n"
	}

	f.contentBuffer[i] = s
}

func constructPath(preffix, directory, fileName string) (s string) {
	if !reflect.DeepEqual(preffix, "") {
		s += preffix + "/"
	}
	s += directory + "/" + fileName
	return s
}

// ~~~~~~~~~~~~~~~~~~~~ MarkdownFile

type MarkdownFile struct {
	File
}

func NewMarkdownFile(name, path string) *MarkdownFile {
	return &MarkdownFile{File: *NewFileWithSuffix(name, "md", path)}
}

func (m *MarkdownFile) appendMarkdownLink(displayText, path string) {
	m.append(constructMarkDownLink(false, displayText, path))
}

func (m *MarkdownFile) appendMarkdownEmbed(path string) {
	m.append(constructMarkDownLink(true, "", path))
}

func constructMarkDownLink(embed bool, displayText, path string) (s string) {
	if embed {
		s += "!"
	}
	s += fmt.Sprintf("[%s](%s)", displayText, path)
	return
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

			// check if the string is empty, if so skip
			if reflect.DeepEqual(csvCell, "") {
				continue
			}

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

func (c *CSVFile) numHeaders() int {
	return len(c.headings)
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
