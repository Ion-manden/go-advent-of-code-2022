package reader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func CreateReaderFromFile(filePath string) *bufio.Reader {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewReader(f)
}

func CreateReaderFromString(input string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(input))
}
