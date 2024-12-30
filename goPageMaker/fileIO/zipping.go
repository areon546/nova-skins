package fileIO

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/areon546/NovaDriftCustomSkins/goPageMaker/helpers"
)

func EmptyZip() *ZipFile {
	return &ZipFile{}
}

// Steps to Zip a File
// 1. Create the .zip file you want to write to
// 2. Create the writer to the .zip file
// 3. In the zip writer, make a virtual file
// 4. In the virtual file, add the contents to the related physical file
// 4b. Close the virtual file writer
// 5. Repeat steps 3 and 4 for each file you want to zip
// 6. Close

// ~~~~~~~~~~~~~~~~~~~~ ZipFile
type ZipFile struct {
	writer zip.Writer
	name   string
	file   *os.File
}

func NewZipFile(name string) *ZipFile {
	name = constructZipName(name)

	print(name)
	file, err := os.Create(name)
	handle(err)

	return &ZipFile{writer: *zip.NewWriter(file), name: name, file: file}
	// return &ZipFile{}
}

func constructZipName(name string) string {
	if helpers.Search("zip", strings.Split(name, ".")) > -1 {
		return name
	}
	return helpers.Format("%s.zip", name)
}

func (z *ZipFile) GetName() string { return z.name }

func (z *ZipFile) AddZipFile(filename string, contents io.Reader) {
	fileWriter, err := z.writer.Create(filename)
	handle(err)

	_, err = io.Copy(fileWriter, contents)
	handle(err)
}

func (z *ZipFile) WriteToZipFile() {
	z.writer.Close()
	z.file.Close()
}

func (z *ZipFile) Close() {
	z.file.Close()
}

func ZipFolder(path, output string) {
	// TODO tomorrow, we read this, make notes of how it works
	// TODO then we fix zip file

	// here we create the zip zipFile
	zipFile, err := os.Create(helpers.Format("%s.zip", output))
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	// helpers.Print(file.Name())

	// here we create the zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// action performed at each file
	walker := func(path string, info os.FileInfo, err error) error {
		helpers.Printf("Crawling: %v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// here we open the fileToZip that we want to zip
		fileToZip, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		if path[0] == "/"[0] { // TODO not really proper, im just following the advice above and should do the job
			err = errors.New("not allowed to have an absolute path")
			return err
		}

		// HERE is the actual file processing, above is error checking

		// here, in the zip Writer, we create the virtual file fileBeingZipped
		fileBeingZipped, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		// here we copy the contents in the physical file to the virtual file being zipped
		_, err = io.Copy(fileBeingZipped, fileToZip)
		if err != nil {
			return err
		}

		return nil
	}

	// performs function `walker` on each file within path, recursively
	err = filepath.Walk(path, walker)
	if err != nil {
		panic(err)
	}

}
