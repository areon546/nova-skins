package fileIO

import (
	"log"
	"os"
)

func filterFiles(arr []os.DirEntry) (fs []File, folders []os.DirEntry) {

	for _, v := range arr {

		if !v.IsDir() { // for any files, turn them into files
			vName, suffix := splitFileName(v.Name())
			fs = append(fs, *NewFileWithSuffix(vName, suffix, ""))
		} else {
			folders = append(folders, v)
		}
	}

	return
}

// This function
func ReadDirectory(dirPath string) (entries []os.DirEntry) {
	// read contents of a directory
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	return
}
