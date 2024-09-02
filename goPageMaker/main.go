package main

import (
	// "bytes"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

// tell program page or have it have a csv of pages
// use csv and have like 10 max per page
// have it then use assets to place into page

func main() {

	// reads what files are in the assets folder
	assets := readAssets()
	count := len(assets)
	isFolder := make([]bool, count)

	for i, asset := range assets {
		print(reflect.TypeOf(asset))
		print(asset.IsDir())
		isFolder[i] = asset.IsDir()
	}

	/*
	currently, assets is everything that is in the folder assets
	I want it to save a list of all assets in the folder in a file locally

	to do this i have to read the assets csv file and from it, compare against the assets slice
	 */


	// load csv
	
	
	

	// writePagePreffix("file.md", 0)


	// if there are new files (files not present in the CSV file)

	createNewPage()

	// TODO reading: file, err := os.Open("../assets/file.txt") // For read access.

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

func writePagePreffix(fileName string, pageNumber int) error {
	// write to file:
	// Page #
	// prev next
	writeFile(fileName, fmt.Sprintf("Page %d", pageNumber))
	err := writePrevNextPage(fileName, pageNumber)

	return err
}

func writePrevNextPage(fileName string, pageNumber int) error {
	path := "../pages/"
	links := ""

	if pageNumber > 1 {

		links += constructMarkdownLink(false, "Page 1", (path + fmt.Sprintf("Page%d.md", (pageNumber-1))))
	}

	writeFile(fileName, links)

	return nil
}

type CSVFile struct {
	file     os.File
	contents [][]byte
}
