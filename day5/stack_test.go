package day5

import (
	"log"
	"testing"
)

func TestParseLineAsSections(t *testing.T) {
	expect := []string{"", "D", ""}
	got := parseLineAsSections("    [D]    ")
	if got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
		log.Println("expect: ", expect)
		log.Println("got: ", got)
		t.Fail()
	}

	expect = []string{"N", "C", ""}
	got = parseLineAsSections("[N] [C]    ")
	if got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
		log.Println("expect: ", expect)
		log.Println("got: ", got)
		t.Fail()
	}

	expect = []string{"Z", "M", "P"}
	got = parseLineAsSections("[Z] [M] [P]")
	if got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
		log.Println("expect: ", expect)
		log.Println("got: ", got)
		t.Fail()
	}
}

func TestMoveStack(t *testing.T) {
	expect := [][]string{
		{"Z", "N", "D"},
		{"M", "C"},
		{"P"},
	}
	got := moveStack(
		"move 1 from 2 to 1",
		[][]string{
			{"Z", "N"},
			{"M", "C", "D"},
			{"P"},
		},
	)
	if got[0][2] != expect[0][2] {
		log.Println(got)
		t.Fail()
	}

	expect = [][]string{
		{},
		{"M", "C"},
		{"P", "D", "N", "Z"},
	}
	got = moveStack(
		"move 3 from 1 to 3",
		[][]string{
			{"Z", "N", "D"},
			{"M", "C"},
			{"P"},
		},
	)
	if got[2][3] != expect[2][3] {
		log.Println(got)
		t.Fail()
	}
}

func TestMoveStackBatch(t *testing.T) {
	expect := [][]string{
		{"Z", "N", "D"},
		{"M", "C"},
		{"P"},
	}
	got := moveStackBatch(
		"move 1 from 2 to 1",
		[][]string{
			{"Z", "N"},
			{"M", "C", "D"},
			{"P"},
		},
	)
	if got[0][2] != expect[0][2] {
		log.Println(got)
		t.Fail()
	}

	expect = [][]string{
		{},
		{"M", "C"},
		{"P", "Z", "N", "D"},
	}
	got = moveStackBatch(
		"move 3 from 1 to 3",
		[][]string{
			{"Z", "N", "D"},
			{"M", "C"},
			{"P"},
		},
	)
	if got[2][3] != expect[2][3] {
		log.Println(got)
		t.Fail()
	}
}
