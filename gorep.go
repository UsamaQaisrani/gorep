package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"strings"
	"errors"
)

func main() {

	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := readLines(f)

	for line := range lines {
		fmt.Println(line)
	}
}	

func readLines(f io.ReadCloser) <- chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)
		buffer := ""

		for {
			text := make([]byte, 8)

			_, err := f.Read(text)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break	
				}
				log.Fatal(err)
			}

			buffer += string(text)
			parts := strings.Split(buffer, "\n")

			if len(parts) > 1 {
				lines <- fmt.Sprintf("%s", parts[0])
				buffer = parts[1]
			}
		}
	}()
	return lines
}
