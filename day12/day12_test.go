package day12

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay12Part1(t *testing.T) {
	expect := 31
	got := Day12Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 352
	got = Day12Part1(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}

func TestDay12Part2(t *testing.T) {
	expect := 29
	got := Day12Part2(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 345
	got = Day12Part2(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}
