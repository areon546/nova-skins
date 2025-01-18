package formatter

type markdownFormatter struct {
}

func NewMarkdownFormatter() markdownFormatter {
	return markdownFormatter{}
}

func (m markdownFormatter) FormatLink(displayText, link string) string {
	return markdownLink(false, displayText, link)
}

func (m markdownFormatter) FormatEmbed(link string) string {
	return markdownLink(true, "", link)
}

func (m markdownFormatter) FormatHeading(tier int, heading string) string {
	s := ""

	for range tier {
		s += "#"
	}

	return s + " " + heading
}

func (m markdownFormatter) FormatTable(t Table, heading bool) string {
	s := ""

	headers := constructRow(t.headers)
	headerDecleration := markdownHeaderDeclarationRow(t.headers.Size())

	if heading {
		s += headers
	}

	s += headerDecleration

	for i := 0; i < t.Rows(); i++ {
		r, _ := t.GetRow(i)
		s += constructRow(r)
	}

	return s
}

func (m markdownFormatter) FormatBold(s string) string {

	return format("**%s**", s)
}

func (m markdownFormatter) FormatItalic(s string) string {
	return format("*%s*", s)
}

// Helper functions.
func markdownLink(embed bool, displayText, path string) (s string) {
	if embed {
		s += "!"
	}
	s += format("[%s](%s)", displayText, path)
	return
}

func constructRow(r row) (row string) {
	row += "|"
	for i := 0; i < r.Size(); i++ {
		v := r.Get(i)
		row += constructCell(v)
	}

	row += "\n" // since a row needs to start and end with a | on a new line to be valid

	return
}

func constructCell(cell string) string {
	return cell + " | "
}

func markdownHeaderDeclarationRow(length int) string {

	headerDecleration := NewRow(length)

	for i := range length {
		headerDecleration.Set(i, "---")
	}

	return constructRow(*headerDecleration)
}
