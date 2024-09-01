package main

import (
	// "bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello World")

	if err := os.WriteFile("../file.txt", []byte("Hello GOSAMPLES!"), 0666); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("../file.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(file.Name())


	bytes := make([]byte, 5)
	count, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	
	fmt.Printf("read %d bytes: %q\n", count, bytes[:count])

}

// tell program page or have it have a csv of pages
// use csv and have like 10 max per page
// have it then use assets to place into page
