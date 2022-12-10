package day9

import (
	"bufio"
	"fmt"
	"io"
)

var tilesVisited = map[string]bool{}

func Day9Part1(r *bufio.Reader) int {
	// reset log
	tilesVisited = map[string]bool{}

	robe := []knot{}
	for i := 0; i < 2; i++ {
		robe = append(robe, knot{xIndex: 0, yIndex: 0})
	}
	// Log starting point
	tileKey := fmt.Sprint(0, ":", 0)
	tilesVisited[tileKey] = true

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)

		robe = moveHead(robe, line)
	}

	return len(tilesVisited)
}

func Day9Part2(r *bufio.Reader) int {
	// reset log
	tilesVisited = map[string]bool{}

	robe := []knot{}
	for i := 0; i < 10; i++ {
		robe = append(robe, knot{xIndex: 0, yIndex: 0})
	}

	// Log starting point
	tileKey := fmt.Sprint(0, ":", 0)
	tilesVisited[tileKey] = true

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)

		robe = moveHead(robe, line)
	}

	return len(tilesVisited)
}
