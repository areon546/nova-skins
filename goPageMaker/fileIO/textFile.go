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

func (t *TextFile) WriteLine(s string, i int, newline bool) {

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

	t.WriteLine(s, lastLine, true)
}

func (t *TextFile) AppendLines(arr []string, newline bool) {
	for _, v := range arr {
		t.Append(v, newline)
	}
}

func (f *TextFile) Append(s string, newline bool) {
	f.WriteLine(s, len(f.contentBuffer), newline)
}

func (f *TextFile) AppendNewLine(s string) {
	f.Append(s, true)
}

func (t *TextFile) AppendEmptyLine() {
	t.Append("", true)
}
