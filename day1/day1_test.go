package day1

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	expect := 24000
  got := Day1Part1(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 69912
  got = Day1Part1(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}

func TestDay1Part2(t *testing.T) {
	expect := 45000
  got := Day1Part2(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 208180
  got = Day1Part2(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}

