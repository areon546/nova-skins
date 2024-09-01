package main

import (
	// "bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// fmt.Println("Hello World")

	// reads what files are in the assets folder
	readAssets()

	// if there are new files (files not present in the CSV file)

	createNewPage()

	// TODO reading: file, err := os.Open("../assets/file.txt") // For read access.

}

func readAssets() {
	readDirectory("../assets")
}

func readDirectory(dirPath string) {
		// read contents of a directory
		entries, err := os.ReadDir(dirPath)
		if err != nil {
			log.Fatal(err)
		}
	
		for _, e := range entries {
			fmt.Println(e.Name())
		}
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

}

func createNewPage() {}

func writePagePreffix(file os.File) {
	// write to file:
	// Page #
	// prev next
}

func writePrevNextPage() {
	constructMarkdownLink(false, "Page 1", "pages/Page1.md")
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
