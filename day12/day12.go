package day12

import (
	"bufio"
)

func Day12Part1(r *bufio.Reader) int {
	matrix := readFileToMatrix(r)

	startingNode := findStartInMatrix(matrix)
	return traverseMatrix(startingNode,"E", matrix, true, false)
}

func Day12Part2(r *bufio.Reader) int {
	matrix := readFileToMatrix(r)

  // Start at the top
	startingNode := findEndInMatrix(matrix)
	return traverseMatrix(startingNode, "a", matrix, false, true)
}
