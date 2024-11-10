package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

// helper functions

func print(a ...any) { fmt.Println(a...) }

func printf(format string, a ...any) { fmt.Printf("%s\n", (fmt.Sprintf(format, a...))) }

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func assetsCSVPath() string {
	return "assets.csv"
}

func constructMarkdownLink(embed bool, displayText, path string) string {
	if embed {
		return fmt.Sprintf("![%s](%s)", displayText, path)
	}
	return fmt.Sprintf("[%s](%s)", displayText, path)
}

func search(item string, arr []string) (index int) {
	index = -1
	for i, v := range arr {
		if reflect.DeepEqual(v, item) {
			index = i
		}
	}
	return index
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

func skinFolder() string {
	return "../custom_skins/"
}

type Stringable interface {
	toString() string
}
