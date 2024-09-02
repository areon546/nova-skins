package main

import (
	// "bytes"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func main() {
	// fmt.Println("Hello World")

	// reads what files are in the assets folder
	directoryContents := readAssets()
	count := len(directoryContents)
	isFolder := make([]bool, count)

	for i, dirContent := range directoryContents {
		print(reflect.TypeOf(dirContent))
		print(dirContent.IsDir())
		isFolder[i] = dirContent.IsDir()
	}

	writePagePreffix("file.md", 0)

	// load csv
	

	// if there are new files (files not present in the CSV file)

	createNewPage()

	// TODO reading: file, err := os.Open("../assets/file.txt") // For read access.

}

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

func readCSV(fileName string) {
	// read fileName into CSVFile
	// fileContents := make([][]byte, 1)

	return
}

func readFile(fileN os.File, b []byte, fileName string) (lengthOfLine int, err error) {
	// read contents of a file
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	lengthOfLine, err = file.Read(b)
	if err == io.EOF {
		fmt.Println("End of file")
	} else if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", lengthOfLine, b[:lengthOfLine])
	return
}

func writeFile(fileName, text string) {
	if err := os.WriteFile(fileName, []byte(text), 0666); err != nil {
		log.Fatal(err)
	}
}

func constructMarkdownLink(embed bool, displayText, path string) (link string) {
	if embed {
		link += "!"
	}
	return fmt.Sprintf("[%s](%s)", displayText, path)
}

func appendToFile(file os.File) {
	// is supposed to add a 
}

func createNewPage() {}

func writePagePreffix(fileName string, pageNumber int) {
	// write to file:
	// Page #
	// prev next
	writePageNumber(fileName, pageNumber)
	writePrevNextPage(fileName)
}

func writePageNumber(fileName string, n int) error { // TODO add error handling

	writeFile(fileName, fmt.Sprintf("Page %d", n))
	
	return nil
}

func writePrevNextPage(fileName string) error {
	link := constructMarkdownLink(false, "Page 1", "pages/Page1.md")

	writeFile(fileName, link)

	return nil
}

type CSVFile struct {
	file     os.File
	contents [][]byte
}

// tell program page or have it have a csv of pages
// use csv and have like 10 max per page
// have it then use assets to place into page

// helper functions

func print(a ...any) { fmt.Println(a...) }
