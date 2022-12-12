package day12

import (
	"bufio"
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

func Day12Part1(r *bufio.Reader) int {
	matrix := readFileToMatrix(r)

	startingNode := findStartInMatrix(matrix)
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

	nodesTried := []string{}
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		nKey := fmt.Sprint(item.value.li, ":", item.value.char, ":", item.value.ni)
		nodesTried = append(nodesTried, nKey)
		if item.value == endNode {
			return item.pathCost
		}

		possibleNodes := getPosibleTouchingNodes(item.value, matrix)

	nodeLoop:
		for _, n := range possibleNodes {
			nKey := fmt.Sprint(n.li, ":", n.char, ":", n.ni)
			for _, seenNode := range nodesTried {
				if nKey == seenNode {
					continue nodeLoop
				}
			}

			lengthToEnd := getLengthToEnd(n, endNode)

			pathCost := item.pathCost + 1
			// Insert a new item and then modify its priority.
			item := &Item{
				value:       n,
				pathCost:    pathCost,
				lengthToEnd: lengthToEnd,
				priority:    pathCost + lengthToEnd,
			}
			heap.Push(&pq, item)
		}
	}

	searchedMatrix := [][]string{}
  for lines := 0; lines < 41; lines++ {
    line := []string{}
    for n := 0; n < 70; n++ {
      line = append(line, " ")
    }
    searchedMatrix = append(searchedMatrix, line)
  }
	for _, nk := range nodesTried {
		parts := strings.Split(nk, ":")
		li, _ := strconv.Atoi(parts[0])
		n := parts[1]
		ni, _ := strconv.Atoi(parts[2])
		searchedMatrix[li][ni] = n
	}

	for _, l := range searchedMatrix {
		for _, n := range l {
			print(n)
		}
		print("\n")
	}

	return 0
}

func Day12Part2(r *bufio.Reader) int {
	return 0
}
