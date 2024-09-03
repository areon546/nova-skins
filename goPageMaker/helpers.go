package main

import (
	"fmt"
	"log"
)

// helper functions

func print(a ...any) { fmt.Println(a...) }

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}