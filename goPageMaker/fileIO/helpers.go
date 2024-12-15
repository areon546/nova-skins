package fileIO

import (
	"github.com/areon546/go-helpers"
)

// helper functions

func Handle(err error) {
	// helpers.Handle(err)
	helpers.CheckError(err)
}
