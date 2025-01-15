package formatter

import (
	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

// helper functions

func handle(err error) {
	// helpers.Handle(err)
	helpers.Handle(err)
}

func print(a ...any) {
	helpers.Print(a...)
}

func format(s string, a ...any) string { return helpers.Format(s, a...) }
