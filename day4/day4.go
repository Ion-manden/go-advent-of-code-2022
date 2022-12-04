package day4

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type elfAssignment struct {
	assignment string
	lower      int
	upper      int
}

func createElfAssignment(assignment string) elfAssignment {
	edges := strings.Split(assignment, "-")
	lower, err := strconv.Atoi(edges[0])
	if err != nil {
		log.Fatal(err)
	}
	upper, err := strconv.Atoi(edges[1])
	if err != nil {
		log.Fatal(err)
	}

	return elfAssignment{
		assignment: assignment,
		lower:      lower,
		upper:      upper,
	}
}

func Day4Part1(r *bufio.Reader) int {
	fullyOverlappingPairs := 0

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)

		elfAssignments := strings.Split(line, ",")
		elf1 := createElfAssignment(elfAssignments[0])
		elf2 := createElfAssignment(elfAssignments[1])

		if elf1.lower <= elf2.lower && elf1.upper >= elf2.upper {
			fullyOverlappingPairs += 1
			continue
		}
		if elf2.lower <= elf1.lower && elf2.upper >= elf1.upper {
			fullyOverlappingPairs += 1
			continue
		}
	}

	return fullyOverlappingPairs
}

func Day4Part2(r *bufio.Reader) int {
	hasOverlappingPairs := 0

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)

		elfAssignments := strings.Split(line, ",")
		elf1 := createElfAssignment(elfAssignments[0])
		elf2 := createElfAssignment(elfAssignments[1])

		if elf1.lower <= elf2.upper && elf1.upper >= elf2.lower {
			hasOverlappingPairs += 1
			continue
		}
		if elf2.lower <= elf1.upper && elf2.upper >= elf1.lower {
			hasOverlappingPairs += 1
			continue
		}
	}

	return hasOverlappingPairs
}
