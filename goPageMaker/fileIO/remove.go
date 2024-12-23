package fileIO

import "os"

// I want to remove all the files within a given directory, but keep the directory

func RemoveAllWithinDirectory(path string) {
	os.RemoveAll(path)

	os.Mkdir(path, os.ModePerm) // TODO too permissive, does 0777 when i dont need it to
}
