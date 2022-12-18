package matrix

import (
	"log"
	"testing"
)

func TestDistance(t *testing.T) {
	expect := 10
	got := Distance(Point{X: 0, Y: 0}, Point{X: 5, Y: -5})

	if expect != got {
		log.Println("Expected: ", expect)
		log.Println("Got: ", got)
		t.Fail()
	}

	expect = 12
	got = Distance(Point{X: -3, Y: 5}, Point{X: -5, Y: -5})

	if expect != got {
		log.Println("Expected: ", expect)
		log.Println("Got: ", got)
		t.Fail()
	}
}

func TestGetBorderPoints(t *testing.T) {
	expect := []Point{
    {X: 0, Y: 2},
    {X: 1, Y: 1},
    {X: 2, Y: 0},
    {X: 1, Y: -1},
    {X: 0, Y: -2},
    {X: -1, Y: -1},
    {X: -2, Y: 0},
    {X: -1, Y: 1},
  }
	got := GetBorderPoints(Point{X: 0, Y: 0}, 1)

  for i := range got {
	if expect[i] != got[i] {
		log.Println("I: ", i)
		log.Println("Expected: ", expect[i])
		log.Println("Got: ", got[i])
		t.Fail()
	}
  }
}
