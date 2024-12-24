package fileIO

import (
	"archive/zip"
	"bytes"
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

// ~~~~~~~~~~~~~~~~~~~~ ZipFile
type ZipFile struct {
	writer zip.Writer
	name   string
}

func NewZipFile(name string) *ZipFile {
	return &ZipFile{writer: *zip.NewWriter(new(bytes.Buffer)), name: name}
}

func (z *ZipFile) GetName() string {
	if helpers.Search("zip", strings.Split(z.name, ".")) > -1 {
		return z.name
	}
	return helpers.Format("%s.zip", z.name)
}

func (z *ZipFile) AddZipFile(filename string, contents []byte) {
	fileWriter, err := z.writer.Create(filename)

	handle(err)

	_, err = fileWriter.Write(contents)

	handle(err)

}

func (z *ZipFile) WriteToZipFile() {
	file, err := os.Create(z.GetName())
	handle(err)

	defer file.Close()
}

func (z *ZipFile) Close() {
	z.writer.Close()
}

func ZipFolder(path, output string) {
	file, err := os.Create(helpers.Format("%s.zip", output))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// helpers.Print(file.Name())

	w := zip.NewWriter(file)
	defer w.Close()

	// action performed at each file
	walker := func(path string, info os.FileInfo, err error) error {
		// helpers.Print("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		if path[0] == "/"[0] { // TODO not really proper, im just following the advice above and should do the job
			err = errors.New("not allowed to have an absolute path")
			return err
		}

		f, err := w.Create(path)
		if err != nil {
			return err
		}

		// copy file contents from file to f, the virtual zip file
		_, err = io.Copy(f, file)
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
