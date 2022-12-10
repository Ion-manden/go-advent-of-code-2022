package day8

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay8Part1(t *testing.T) {
	expect := 21
	got := Day8Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 1825
	got = Day8Part1(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}

func TestDay8Part2(t *testing.T) {
	expect := 8
	got := Day8Part2(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 235200
	got = Day8Part2(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}
