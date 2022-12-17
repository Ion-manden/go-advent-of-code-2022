package day14

import (
	"advent-of-code/pkg/matrix"
	"bufio"
)

func Day14Part1(r *bufio.Reader) int {
	rockMap := createRockMapFromFile(r, true)

	var err error
	for {
		rockMap, err = dropSand(rockMap)

		if err != nil {
			break
		}
	}

	return matrix.CountOccurrences(rockMap, "O")
}

func Day14Part2(r *bufio.Reader) int {
	rockMap := createRockMapFromFile(r, false)

	// extend to ends
	for yi := range rockMap {
		for i := 0; i < 500; i++ {
			rockMap[yi] = append(rockMap[yi], ".")
		}
	}

	line := []string{}
	for i := 0; i < len(rockMap[0]); i++ {
		line = append(line, ".")
	}
  rockMap = append(rockMap, line)

	line = []string{}
	for i := 0; i < len(rockMap[0]); i++ {
		line = append(line, "#")
	}
  rockMap = append(rockMap, line)


	var err error
	for {
		rockMap, err = dropSand(rockMap)

		if err != nil {
			break
		}
	}

	return matrix.CountOccurrences(rockMap, "O")
}
