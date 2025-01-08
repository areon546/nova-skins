package fileIO

import (
	"testing"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

func TestMarkdownHeaderDeclarationRow(t *testing.T) {
	want := " | ~~~ | ~~~ | ~~~ | ~~~ | ~~~ |\n"
	got := markdownHeaderDeclarationRow(5)

	helpers.AssertEquals(t, want, got)
}
