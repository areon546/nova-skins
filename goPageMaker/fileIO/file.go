package fileIO

import (
	"errors"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

// ~~~~~~~~~~~~~~~~ File

type File struct {
	filename string
	suffix   string
	relPath  string

	contentBuffer []string
	lines         int
	hasBeenRead   bool

	osFile os.File
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
	return helpers.Format("%s.%s", f.filename, f.suffix)
}

func (f *File) GetContents() []string {
	return f.contentBuffer
}

func (f *File) readFile() []string {
	if !f.hasBeenRead {
		data, err := os.ReadFile(f.GetFileName()) // For read access.
		handle(err)

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
		return "", errors.New("Index out of bounds for File length")
	}

	output = f.contentBuffer[lineNum]
	print(output)

	return
}

func (f *File) WriteFile() {
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

func (f *File) Append(s string) {
	f.AppendLine(s, len(f.contentBuffer), true)
}

func (f *File) AppendNewLine() {
	f.Append("")
}

func (f *File) bufferLines(arr []string) {

	if f.isEmpty() {
		f.contentBuffer = make([]string, len(arr))
	}

	f.contentBuffer = append(f.contentBuffer, arr...)

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

func (f *File) AppendLine(s string, i int, nl bool) {

	for i >= len(f.contentBuffer) {
		f.contentBuffer = append(f.contentBuffer, "")
	}

	if nl {
		s += "\n"
	}

	f.contentBuffer[i] = s
}

func ConstructPath(preffix, directory, fileName string) (s string) {
	if !reflect.DeepEqual(preffix, "") {
		s += preffix + "/"
	}

	s += directory

	if !reflect.DeepEqual(fileName, "") {
		s += "/" + fileName
	}
	return s
}
