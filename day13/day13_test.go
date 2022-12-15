package day13

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay13Part1(t *testing.T) {
	expect := 13
	got := Day13Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	// expect = 352
	// got = Day13Part1(reader.CreateReaderFromFile("./input.txt"))
	// if got != expect {
	// 	log.Println("got: ", got)
	// 	log.Println("expect: ", expect)
	// 	t.Fail()
	// }
}

func TestDay13Part2(t *testing.T) {
	// expect := 29
	// got := Day13Part2(reader.CreateReaderFromFile("./example.txt"))
	// if got != expect {
	// 	log.Println("got: ", got)
	// 	log.Println("expect: ", expect)
	// 	t.Fail()
	// }

	// expect = 345
	// got = Day13Part2(reader.CreateReaderFromFile("./input.txt"))
	// if got != expect {
	// 	log.Println("got: ", got)
	// 	log.Println("expect: ", expect)
	// 	t.Fail()
	// }
}
