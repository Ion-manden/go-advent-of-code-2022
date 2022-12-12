package day12

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"math"
	"strings"
)

type node struct {
	char string
	// List index
	li int
	// Node index
	ni int
}

func getLengthToEnd(n node, end node) int {
	liCost := math.Abs(float64(n.li - end.li))
	niCost := math.Abs(float64(n.ni - end.ni))

	return int(liCost + niCost)
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

func getPosibleTouchingNodes(currentNode node, matrix [][]string, runningReverse bool) []node {
	adjacentNodes := []node{}
	if currentNode.li > 0 {
		char := matrix[currentNode.li-1][currentNode.ni]
		adjacentNodes = append(adjacentNodes, node{char: char, li: currentNode.li - 1, ni: currentNode.ni})
	}
	if currentNode.li < len(matrix)-1 {
		char := matrix[currentNode.li+1][currentNode.ni]
		adjacentNodes = append(adjacentNodes, node{char: char, li: currentNode.li + 1, ni: currentNode.ni})
	}

	if currentNode.ni > 0 {
		char := matrix[currentNode.li][currentNode.ni-1]
		adjacentNodes = append(adjacentNodes, node{char: char, li: currentNode.li, ni: currentNode.ni - 1})
	}
	if currentNode.ni < len(matrix[0])-1 {
		char := matrix[currentNode.li][currentNode.ni+1]
		adjacentNodes = append(adjacentNodes, node{char: char, li: currentNode.li, ni: currentNode.ni + 1})
	}

	currentnodeHeight := getNodeHeight(currentNode.char)
	heightFilteredNodes := []node{}
	for _, n := range adjacentNodes {
		nHeight := getNodeHeight(n.char)
		if runningReverse {
			if currentnodeHeight-1 <= nHeight {
				heightFilteredNodes = append(heightFilteredNodes, n)
			}
		} else {
			if currentnodeHeight+1 >= nHeight {
				heightFilteredNodes = append(heightFilteredNodes, n)
			}
		}
	}

	return heightFilteredNodes
}

// Returns steps it took to get to end
func traverseMatrix(startingNode node, endChar string, matrix [][]string, useLengthToEndInPriority bool, runningReverse bool) int {
	endNode := findEndInMatrix(matrix)

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		value:    startingNode,
		priority: 0,
		index:    0,
	}
	heap.Init(&pq)

	nodesTried := map[string]bool{}
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		k := fmt.Sprint(item.value.li, ":", item.value.char, ":", item.value.ni)
		_, ok := nodesTried[k]
		if ok {
			continue
		}
		nodesTried[k] = true
		if item.value.char == endChar {
			return item.pathCost
		}

		possibleNodes := getPosibleTouchingNodes(item.value, matrix, runningReverse)

	nodeLoop:
		for _, n := range possibleNodes {
			nKey := fmt.Sprint(n.li, ":", n.char, ":", n.ni)
			_, ok := nodesTried[nKey]
			if ok {
				continue nodeLoop
			}

			lengthToEnd := 0
			pathCost := item.pathCost + 1
			priority := pathCost
			if useLengthToEndInPriority {
				lengthToEnd = getLengthToEnd(n, endNode)
				priority += lengthToEnd
			}
			// Insert a new item and then modify its priority.
			it := &Item{
				value:       n,
				pathCost:    pathCost,
				lengthToEnd: lengthToEnd,
				priority:    priority,
			}
			heap.Push(&pq, it)
		}
	}

	return 0
}
