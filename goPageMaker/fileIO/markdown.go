package fileIO

import "fmt"

// ~~~~~~~~~~~~~~~~~~~~ MarkdownFile
type MarkdownFile struct {
	File
}

func NewMarkdownFile(name, path string) *MarkdownFile {
	return &MarkdownFile{File: *NewFileWithSuffix(name, "md", path)}
}

func (m *MarkdownFile) AppendMarkdownLink(displayText, path string) {
	m.Append(ConstructMarkDownLink(false, displayText, path))
}

func (m *MarkdownFile) AppendMarkdownEmbed(path string) {
	m.Append(ConstructMarkDownLink(true, "", path))
}

func ConstructMarkDownLink(embed bool, displayText, path string) (s string) {
	if embed {
		s += "!"
	}
	s += fmt.Sprintf("[%s](%s)", displayText, path)
	return
}
