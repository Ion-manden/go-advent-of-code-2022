package day16

import (
	"advent-of-code/reader"
	"log"
	"testing"
)

func TestLineToSection(t *testing.T) {
	expect := section{name: "AA", rate: 0, tunnelsTo: []string{"DD", "II", "BB"}}
	got := lineToSection("Valve AA has flow rate=0; tunnels lead to valves DD, II, BB")

	if got.name != expect.name || got.rate != expect.rate {
		log.Println("got: ", got)
		log.Println("expect: ", expect)
		t.Fail()
	}
}

func TestDistanceToSection(t *testing.T) {
	expectDistance := 5
	expectPath := []string{"AA", "DD", "EE", "FF", "GG", "HH"}

	gotDistance, gotPath := distanceToSection(
		createSectionsFromReader(reader.CreateReaderFromFile("./example.txt")),
		"AA",
		"HH",
	)

	if gotDistance != expectDistance {
		log.Println("got: ", gotDistance)
		log.Println("expect: ", expectDistance)
		t.Fail()
	}

	for i := range gotPath {
		if gotPath[i] != expectPath[i] {
			log.Println("got: ", gotPath[i])
			log.Println("expect: ", expectPath[i])
			t.Fail()
		}
	}
}
