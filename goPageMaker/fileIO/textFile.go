package fileIO

import (
	"errors"
	"os"
	"strings"
)

type TextFile struct {
	File
	textBuffer []string
}

func NewTextFileWithSuffix(path, filename, suff string) *TextFile {
	return &TextFile{File: *NewFileWithSuffix(filename, suff, path)}
}

func NewTextFile(path, filename string) *TextFile {
	filename, suff := splitFileName(filename)
	return NewTextFileWithSuffix(filename, suff, path)
}

func (f *TextFile) ReadFile() []string {
	if !f.hasBeenRead {
		data, err := os.ReadFile(f.Name()) // For read access.
		handle(err)

		oneLine := strings.ReplaceAll(string(data), "\r", "")
		f.textBuffer = strings.Split(oneLine, "\n")
		f.lines = len(f.contentBuffer)
	}
	return f.textBuffer
}

func (f *TextFile) ReadLine(lineNum int) (output string, err error) {
	lineNum -= 1 // converted to index notation

	if f.IsEmpty() {
		f.ReadFile()
	}

	if lineNum > f.lines {
		return "", errors.New("Index out of bounds for File length")
	}

	output = string(f.textBuffer[lineNum])
	// print(output)

	return
}

func (t *TextFile) AppendLine(s string, i int, newline bool) {

	for i >= len(t.textBuffer) {
		t.textBuffer = append(t.textBuffer, "")
	}

	if newline {
		s += "\n"
	}

	t.textBuffer[i] = s
	t.File.AppendLine(s)
}

func (t *TextFile) AppendLastLine(s string) {
	lastLine := len(t.contentBuffer) - 1

	if t.IsEmpty() {
		lastLine = 0
	}

	t.AppendLine(s, lastLine, true)
}

func (t *TextFile) AppendLines(arr []string) {
	for _, v := range arr {
		t.Append(v)
	}
}

func (f *TextFile) Append(s string) {
	f.AppendLine(s, len(f.contentBuffer), true)
}

func (f *TextFile) AppendNewLine() {
	f.Append("")
}

// func (f *TextFile) bufferLines(arr []byte) {

// 	if f.IsEmpty() {
// 		f.contentBuffer = make([]byte, len(arr))
// 	}

// 	f.contentBuffer = append(f.contentBuffer, arr...)

// }

// func (f *TextFile) BufferToString() string {
// 	s := ""
// 	for _, v := range f.textBuffer {
// 		s += v
// 	}

// 	return s
// }

// func (f *TextFile) WriteFile() {
// 	if err := os.WriteFile(f.Name(), []byte(f.BufferToString()), 0664); err != nil {
// 		log.Fatal(err)
// 	}
// }
