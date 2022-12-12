package day12

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type node struct {
	char string
	// List index
	li int
	// Node index
	ni int
}

func getNodeHeight(node string) int {
	// Start node has hight equal to a
	if node == "S" {
		node = "a"
	}
	// End node has hight equal to z
	if node == "E" {
		node = "z"
	}

	chars := "abcdefghijklmnopqrstuvwxyz"
	return strings.Index(chars, node) + 1
}

func readFileToMatrix(r *bufio.Reader) [][]string {
	matrix := [][]string{}
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		matrix = append(matrix, strings.Split(line, ""))
	}

	return matrix
}

func findStartInMatrix(matrix [][]string) node {
	for li, line := range matrix {
		for ni, n := range line {
			if n == "S" {
				return node{char: n, li: li, ni: ni}
			}
		}
	}

	log.Fatal("Start was not found")
	return node{}
}

func findEndInMatrix(matrix [][]string) node {
	for li, line := range matrix {
		for ni, n := range line {
			if n == "E" {
				return node{char: n, li: li, ni: ni}
			}
		}
	}

	log.Fatal("End was not found")
	return node{}
}

// Returns steps it took to get to end
func traverseNode(currentNode node, matrix [][]string, nodesTried []string) []int {
	log.Printf("%+v\n", currentNode)
	nodesTried = append(nodesTried, fmt.Sprint(currentNode.li, ":", currentNode.ni))
	possibleNextNodes := []node{}

	if currentNode.li > 0 {
		char := matrix[currentNode.li-1][currentNode.ni]
		possibleNextNodes = append(possibleNextNodes, node{char: char, li: currentNode.li - 1, ni: currentNode.ni})
	}
	if currentNode.li < len(matrix)-1 {
		char := matrix[currentNode.li+1][currentNode.ni]
		possibleNextNodes = append(possibleNextNodes, node{char: char, li: currentNode.li + 1, ni: currentNode.ni})
	}

	if currentNode.ni > 0 {
		char := matrix[currentNode.li][currentNode.ni-1]
		possibleNextNodes = append(possibleNextNodes, node{char: char, li: currentNode.li, ni: currentNode.ni - 1})
	}
	if currentNode.ni < len(matrix[0])-1 {
		char := matrix[currentNode.li][currentNode.ni+1]
		possibleNextNodes = append(possibleNextNodes, node{char: char, li: currentNode.li, ni: currentNode.ni + 1})
	}

	currentnodeHeight := getNodeHeight(currentNode.char)
	nodesTriedCounts := []int{}
	endNode := findEndInMatrix(matrix)

nodeLoop:
	for _, n := range possibleNextNodes {
		nHeight := getNodeHeight(n.char)
		if currentnodeHeight+1 < nHeight || currentnodeHeight-1 > nHeight {
			continue
		}

		nKey := fmt.Sprint(n.li, ":", n.ni)
		for _, seenNode := range nodesTried {
			if nKey == seenNode {
				continue nodeLoop
			}
		}

		if n == endNode {
			nodesTriedCounts = append(nodesTriedCounts, len(nodesTried))
		}

		ntraversedCounts := traverseNode(n, matrix, nodesTried)
		nodesTriedCounts = append(nodesTriedCounts, ntraversedCounts...)
	}

	return nodesTriedCounts
}
