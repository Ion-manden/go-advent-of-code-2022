package day2

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	expect := 15
  got := Day2Part1(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 10404
  got = Day2Part1(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}

func TestDay2Part2(t *testing.T) {
	expect := 12
  got := Day2Part2(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 10334
  got = Day2Part2(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}

