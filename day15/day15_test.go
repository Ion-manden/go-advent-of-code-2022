package day15

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay15Part1(t *testing.T) {
	expect := 26
	got := Day15Part1(reader.CreateReaderFromFile("./example.txt"), 10)
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 5335787
	got = Day15Part1(reader.CreateReaderFromFile("./input.txt"), 2000000)

	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}

func TestDay15Part2(t *testing.T) {
	expect := 56000011
	got := Day15Part2(reader.CreateReaderFromFile("./example.txt"), 20)
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 13673971349056
	got = Day15Part2(reader.CreateReaderFromFile("./input.txt"), 20)
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}
