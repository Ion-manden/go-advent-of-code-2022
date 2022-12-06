package day6

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestShift(t *testing.T) {
	expect := []string{"E", "D", "C", "B"}
	got := shift([]string{"D", "C", "B", "A"}, "E")
	if got[0] != expect[0] || len(got) != len(expect) {
		log.Println(got)
		t.Fail()
	}

	expect = []string{"p", "j", "q", "j"}
	got = shift([]string{"j", "q", "j", "m"}, "p")
	if got[0] != expect[0] && len(got) != len(expect) {
		t.Log(got)
		t.Fail()
	}
}

func TestIsUniq(t *testing.T) {
	expect := true
	got := isUniq([]string{"E", "D", "C", "B"})
	if got != expect {
		log.Println(got)
		t.Fail()
	}
}

func TestDay6Part1(t *testing.T) {
	expect := 7
	got := Day6Part1(reader.CreateReaderFromFile("./example.txt"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 5
	got = Day6Part1(reader.CreateReaderFromString("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 6
	got = Day6Part1(reader.CreateReaderFromString("nppdvjthqldpwncqszvftbrmjlhg"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 10
	got = Day6Part1(reader.CreateReaderFromString("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 11
	got = Day6Part1(reader.CreateReaderFromString("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 1766
	 got = Day6Part1(reader.CreateReaderFromFile("./input.txt"))
	 if got != expect {
	   log.Println(got)
	   t.Fail()
	 }
}

func TestDay6Part2(t *testing.T) {
	expect := 19
  got := Day6Part2(reader.CreateReaderFromFile("./example.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }

	expect = 23
	got = Day6Part2(reader.CreateReaderFromString("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 23
	got = Day6Part2(reader.CreateReaderFromString("nppdvjthqldpwncqszvftbrmjlhg"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 29
	got = Day6Part2(reader.CreateReaderFromString("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 26
	got = Day6Part2(reader.CreateReaderFromString("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	if got != expect {
		log.Println(got)
		t.Fail()
	}

	expect = 2383
  got = Day6Part2(reader.CreateReaderFromFile("./input.txt"))
  if got != expect {
    log.Println(got)
    t.Fail()
  }
}
