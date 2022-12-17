package day14

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestCreateMatrixFromFile(t *testing.T) {
	expect := [][]string{
		{".", ".", ".", ".", ".", ".", "+", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "#", ".", ".", ".", "#", "#"},
		{".", ".", ".", ".", "#", ".", ".", ".", "#", "."},
		{".", ".", "#", "#", "#", ".", ".", ".", "#", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "#", "."},
		{"#", "#", "#", "#", "#", "#", "#", "#", "#", "."},
	}

	got := createRockMapFromFile(reader.CreateReaderFromFile("./example.txt"), true)

	if len(expect) != len(got) {
		log.Println("Invalid lenth of lines")
		log.Println("Expected: ", len(expect))
		log.Println("Got: ", len(got))
		t.FailNow()
	}

	for yi := range got {
		if len(expect[yi]) != len(got[yi]) {
			log.Println("Invalid lenth of Row: ", yi)
			log.Println("Expected: ", len(expect[yi]))
			log.Println("Got: ", len(got[yi]))
			t.FailNow()
		}

		for xi := range got[yi] {
			if expect[yi][xi] != got[yi][xi] {
				log.Println("Invalid point: ", yi, xi)
				log.Println("Expected: ", expect[yi][xi])
				log.Println("Got: ", got[yi][xi])
				t.Fail()
			}
		}
	}
}
