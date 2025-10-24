package main

import (
	"fmt"
	"os"
	"log"
	"usamaqaisrani/gorep/src/core/read"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := read.ReadLines(f)

	for line := range lines {
		fmt.Println(line)
	}
}	

