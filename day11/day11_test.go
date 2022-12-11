package day11

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay11Part1(t *testing.T) {
	expect := 10605
	got := Day11Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 62491
	got = Day11Part1(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}

func TestDay11Part2(t *testing.T) {
	expect := 2713310158
	got := Day11Part2(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 17408399184
	 got = Day11Part2(reader.CreateReaderFromFile("./input.txt"))
	 if got != expect {
	log.Println("got: ", got)
	log.Println("expect: ", expect)
	   t.Fail()
	 }
}
