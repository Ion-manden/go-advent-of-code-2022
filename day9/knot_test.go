package day9

import (
	"log"
	"testing"
)

func TestMoveHead(t *testing.T) {
	expect := []knot{
		{xIndex: 1, yIndex: 3},
		{xIndex: 1, yIndex: 2},
	}
	robe := []knot{
		{xIndex: 1, yIndex: 1},
		{xIndex: 0, yIndex: 0},
	}

	got := moveHead(robe, "U 2")
	for i := range expect {
		if got[i].xIndex != expect[i].xIndex {
			log.Println("index: ", i)
			log.Println("got: ", got)
			log.Println("expect: ", expect)
			t.Fail()
		}
	}

	expect = []knot{
		{xIndex: 4, yIndex: 4},
		{xIndex: 4, yIndex: 3},
		{xIndex: 4, yIndex: 2},
		{xIndex: 3, yIndex: 2},
		{xIndex: 2, yIndex: 2},
		{xIndex: 1, yIndex: 1},
		{xIndex: 0, yIndex: 0},
	}
	robe = []knot{
		{xIndex: 4, yIndex: 1},
		{xIndex: 3, yIndex: 0},
		{xIndex: 2, yIndex: 0},
		{xIndex: 1, yIndex: 0},
		{xIndex: 0, yIndex: 0},
		{xIndex: 0, yIndex: 0},
		{xIndex: 0, yIndex: 0},
	}

	got = moveHead(robe, "U 4")
	for i := range expect {
		if got[i].xIndex != expect[i].xIndex {
			log.Println("index: ", i)
			log.Println("got: ", got)
			log.Println("expect: ", expect)
			t.Fail()
		}
	}
}
