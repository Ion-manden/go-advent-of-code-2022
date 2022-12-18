package matrix

import (
	"log"
	"testing"
)


func TestDistance(t *testing.T) {
  expect := 10
  got := Distance(Point{X: 0, Y: 0}, Point{X: 5, Y: 5})

  if expect != got {
		log.Println("Expected: ", expect)
		log.Println("Got: ", got)
		t.Fail()
  }
}
