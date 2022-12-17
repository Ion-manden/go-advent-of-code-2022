package day14

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay14Part1(t *testing.T) {
	expect := 24
	got := Day14Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 1072
	got = Day14Part1(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}

func TestDay14Part2(t *testing.T) {
	expect := 93
	got := Day14Part2(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 24659
	got = Day14Part2(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}
