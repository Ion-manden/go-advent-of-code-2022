package day16

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay16Part1(t *testing.T) {
	expect := 1651
	got := Day16Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// expect = 5335787
	// got = Day16Part1(reader.CreateReaderFromFile("./input.txt"))

	// if got != expect {
	// 	log.Println("got: ", got)
	// 	log.Println("expect: ", expect)
	// 	t.Fail()
	// }
}

func TestDay16Part2(t *testing.T) {
	// expect := 56000011
	// got := Day16Part2(reader.CreateReaderFromFile("./example.txt"))
	// if got != expect {
	// 	log.Println("got: ", got)
	// 	log.Println("expect: ", expect)
	// 	t.Fail()
	// }

	// expect = 13673971349056
	// got = Day16Part2(reader.CreateReaderFromFile("./input.txt"))
	// if got != expect {
	// 	log.Println("got: ", got)
	// 	log.Println("expect: ", expect)
	// 	t.Fail()
	// }
}
