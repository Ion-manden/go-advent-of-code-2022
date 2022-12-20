package day16

import (
	"bufio"
	"container/heap"
	"io"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type section struct {
	name      string
	rate      int
	isOpen    bool
	tunnelsTo []string
}

func lineToSection(line string) section {
	parts := strings.Split(line, " ")

	name := parts[1]
	cleanedNumber := strings.Replace(strings.Replace(parts[4], "rate=", "", 1), ";", "", 1)
	rate, err := strconv.Atoi(cleanedNumber)
	if err != nil {
		log.Fatal("Could not convert to integer: ", cleanedNumber)
	}

	tunnelsTo := []string{}
	for _, part := range parts[9:] {
		sectionName := strings.Replace(part, ",", "", 1)
		tunnelsTo = append(tunnelsTo, sectionName)
	}

	return section{name: name, rate: rate, isOpen: false, tunnelsTo: tunnelsTo}
}

func createSectionsFromReader(r *bufio.Reader) []section {
	sections := []section{}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		s := lineToSection(line)
		sections = append(sections, s)
	}

	return sections
}

// returns distance and path
func distanceToSection(
	sections []section,
	startSectionName string,
	endSectionName string,
) (int, []string) {
	foundI := slices.IndexFunc(sections, func(s section) bool { return s.name == startSectionName })
	start := sections[foundI]

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		value:    start,
		priority: 0,
		index:    0,
    path: []string{startSectionName},
	}
	heap.Init(&pq)

	sectionsTried := map[string]bool{}
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		k := item.value.name
		_, ok := sectionsTried[k]
		if ok {
			continue
		}
		sectionsTried[k] = true
		if item.value.name == endSectionName {
			return item.pathCost, item.path
		}

		possibleNodes := item.value.tunnelsTo

	sectionLoop:
		for _, n := range possibleNodes {
			nKey := n
			_, ok := sectionsTried[nKey]
			if ok {
				continue sectionLoop
			}
			foundI := slices.IndexFunc(sections, func(s section) bool { return s.name == nKey })
			s := sections[foundI]

			pathCost := item.pathCost + 1
			priority := pathCost
			// Insert a new item and then modify its priority.
			it := &Item{
				value:    s,
				pathCost: pathCost,
				priority: priority,
        path: append(item.path, s.name),
			}
			heap.Push(&pq, it)
		}
	}

	return 0, []string{}
}
