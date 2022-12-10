package day10

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay10Part1(t *testing.T) {
	expect := 13140
	got := Day10Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = 13740
	got = Day10Part1(reader.CreateReaderFromFile("./input.txt"))
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}

func TestDay10Part2(t *testing.T) {
	expect := []string{
		"#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".", "#", "#", ".", ".",
		"#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", ".",
		"#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", "#", ".", ".", ".", ".",
		"#", "#", "#", "#", "#", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".",
		"#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#",
		"#", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".", ".", ".", "#", "#", "#", "#", "#", "#", "#", ".", ".", ".", ".", ".",
	}
	got := Day10Part2(reader.CreateReaderFromFile("./example.txt"))
	if len(got) != len(expect) {
		log.Println("got: ", len(got))
		log.Println("expect: ", len(expect))
		t.Fail()
	}
	for i := range expect {
		if got[i] != expect[i] {
			log.Println("i: ", i)
			log.Println("got: ", got[i])
			log.Println("expect: ", expect[i])
			t.Fail()
      return
		}
	}

	// expect = 208180
	//  got = Day10Part2(reader.CreateReaderFromFile("./input.txt"))
	//  if got != expect {
	// log.Println("got: ", got)
	// log.Println("expect: ", expect)
	//    t.Fail()
	//  }
}

func TestGetPixelForCycle(t *testing.T) {
	expect := "#"
	got := getPixelForCycle(10, 8)
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}


	expect = "#"
	got = getPixelForCycle(41, 1)
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}

	expect = "#"
	got = getPixelForCycle(200, 38)
	if got != expect {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}
