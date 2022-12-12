package day12

import (
	"bufio"
	"log"
)

func Day12Part1(r *bufio.Reader) int {
	matrix := readFileToMatrix(r)

	nodesTried := []string{}

	startingNode := findStartInMatrix(matrix)
	nodesSeenCounts := traverseNode(startingNode, matrix, nodesTried)
	log.Println(nodesSeenCounts)

	minSteps := nodesSeenCounts[0]
	for _, nodesSeenCount := range nodesSeenCounts {
		if minSteps > nodesSeenCount {
			minSteps = nodesSeenCount
		}
	}

	return minSteps
}

func Day12Part2(r *bufio.Reader) int {
	return 0
}
