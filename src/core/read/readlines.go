package read
import (
	"io"
	"fmt"
	"strings"
	"errors"
	"log"
	"usamaqaisrani/gorep/src/core/search"
	"usamaqaisrani/gorep/src/core/pattern"
)

func ReadLines(f io.ReadCloser) <- chan string {
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
				if search.Search(pattern.SingleDigit, parts[0]) {
					lines <- fmt.Sprintf("%s", parts[0])
				}
				buffer = parts[1]
			}
		}
		if buffer != "" && search.Search(pattern.SingleDigit, buffer) {
			fmt.Println(buffer)
			lines <- buffer
		}
	}()
	return lines
}
