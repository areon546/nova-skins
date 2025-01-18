package formatter

import "github.com/areon546/NovaDriftCustomSkins/goPageMaker/fileIO"

type Formatter interface {
	FormatLink(displayText, path string) string
	FormatEmbed(path string) string
	FormatHeading(tier int, heading string) string
	FormatTable(t Table, headers bool) string
	FormatBold(s string) string
	FormatItalic(s string) string
}

// ~~~~~~~~~~~~~~~~~~~~ FormattedFile
type FormattedFile struct {
	fileIO.TextFile
	Fmt Formatter
}

func NewHTMLFile(path, name string) *FormattedFile {
	return newFormattedFile(NewMarkdownFormatter(), path, name, "html")
}

func NewMarkdownFile(name, path string) *FormattedFile {
	return newFormattedFile(NewMarkdownFormatter(), path, name, "md")
}

func newFormattedFile(fmt Formatter, path, name, suffix string) *FormattedFile {
	return &FormattedFile{TextFile: *fileIO.NewTextFileWithSuffix(path, name, "md"), Fmt: fmt}
}

func (m *FormattedFile) AppendLink(displayText, path string) {
	m.Append(m.Fmt.FormatLink(displayText, path), false)
}

func (m *FormattedFile) AppendEmbed(path string) {
	m.Append(m.Fmt.FormatEmbed(path), false)
}

func (m *FormattedFile) AppendHeading(tier int, heading string) {
	m.Append(m.Fmt.FormatHeading(tier, heading), false)
}

func (m *FormattedFile) AppendItalics(heading string) {
	m.Append(m.Fmt.FormatItalic(heading), false)
}

func (m *FormattedFile) AppendBold(heading string) {
	m.Append(m.Fmt.FormatBold(heading), false)
}
