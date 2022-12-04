package day4

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay4Part1(t *testing.T) {
	expect := 2
  got := Day4Part1(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 450
  got = Day4Part1(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}

func TestDay4Part2(t *testing.T) {
	expect := 4
  got := Day4Part2(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 837
  got = Day4Part2(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}

