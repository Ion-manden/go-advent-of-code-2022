package reader

import (
	"bufio"
	"log"
	"os"
)

func CreateReaderFromFile(filePath string) *bufio.Reader {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewReader(f)
}
