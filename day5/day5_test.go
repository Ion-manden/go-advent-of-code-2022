package day5

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	expect := "CMZ"
	got := Day5Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = "VPCDMSLWJ"
	got = Day5Part1(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}
}

func TestDay1Part2(t *testing.T) {
	expect := "MCD"
	got := Day5Part2(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = "TPWCGNCCG"
	got = Day5Part2(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}
}
