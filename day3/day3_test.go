package day3

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay3Part1(t *testing.T) {
	expect := 157
  got := Day3Part1(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 7908
  got = Day3Part1(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}

func TestDay3Part2(t *testing.T) {
	expect := 70
  got := Day3Part2(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 2838
  got = Day3Part2(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}

