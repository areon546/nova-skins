package main

import (
	"os"
	"log"
	"fmt"
)

func readAssets() []os.DirEntry {
	return readDirectory("../assets", false)
}

func readDirectory(dirPath string, output bool) (entries []os.DirEntry) {
	// read contents of a directory
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	if output {
		for _, e := range entries {
			fmt.Println(e.Name())
		}
	}

	return
}
