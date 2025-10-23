package main

import (
	"fmt"
	"os"
	"log"
	"usamaqaisrani/gorep/src/core"
)

func main() {

	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := core.ReadLines(f)

	for line := range lines {
		fmt.Println(line)
	}
}	

