package fileIO

import (
	"reflect"
	"strings"

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

func FilesEqual(a, b File) bool {
	return reflect.DeepEqual(a, b)
}
