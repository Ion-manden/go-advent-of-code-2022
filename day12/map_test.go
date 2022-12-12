package day12

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestReadFileToMatrix(t *testing.T) {
	expect := [][]string{
		{"a", "b", "c"},
		{"a", "S", "E"},
		{"a", "b", "c"},
	}

	got := readFileToMatrix(reader.CreateReaderFromString(
		`abc
aSE
abc`,
	))

	for li, line := range got {
		for ci := range line {
			if got[li][ci] != expect[li][ci] {
				log.Println("got: ", got)
				log.Println("expect: ", expect)
				t.Fail()
			}
		}
	}
}
