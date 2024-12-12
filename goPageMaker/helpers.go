package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

// helper functions

func print(a ...any) { fmt.Println(a...) }

func printf(s string, a ...any) { print(format(s, a...)) }

func format(s string, a ...any) string { return fmt.Sprintf(s, a...) }

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func assetsCSVPath() string {
	return "assets.csv"
}

func search(item string, arr []string) (index int) {
	index = -1
	for i, v := range arr {
		if reflect.DeepEqual(v, item) {
			index = i
		}
	}
	return index
}

func convertToInteger(s string) (i int) {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return
}

func handle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
