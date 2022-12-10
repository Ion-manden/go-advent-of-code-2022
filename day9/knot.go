package day9

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type knot struct {
	yIndex int
	xIndex int
}

// Moves the head and moves the tail if needed
func moveHead(robe []knot, action string) []knot {
	parts := strings.SplitN(action, " ", 2)
	direction := parts[0]
	distance, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Println("action that failed: ", action)
		log.Fatal("Could not convert to int: ", parts[1])
	}

	for d := 0; d < distance; d++ {
		switch direction {
		case "U":
			robe[0].yIndex++
		case "D":
			robe[0].yIndex--
		case "L":
			robe[0].xIndex--
		case "R":
			robe[0].xIndex++
		}

		for i := range robe {
			// Skip first
			if i == 0 {
				continue
			}

			robe[i] = moveKnot(robe[i-1], robe[i])

      // if last knot
			if i == len(robe)-1 {
				// Add visit to log
				tileKey := fmt.Sprint(robe[i].xIndex, ":", robe[i].yIndex)
				tilesVisited[tileKey] = true
			}
		}
	}

	return robe
}

// Checks if the tail should be moved
func moveKnot(leadingKnot knot, followingKnot knot) knot {
	// If head is not far enought away, don't move tail
	xDiff := getIntDiff(leadingKnot.xIndex, followingKnot.xIndex)
	yDiff := getIntDiff(leadingKnot.yIndex, followingKnot.yIndex)
	if xDiff <= 1 && yDiff <= 1 {
		return followingKnot
	}

	if xDiff == 0 {
		if leadingKnot.yIndex > followingKnot.yIndex {
			followingKnot.yIndex++
		} else {
			followingKnot.yIndex--
		}

		return followingKnot
	}

	if yDiff == 0 {
		if leadingKnot.xIndex > followingKnot.xIndex {
			followingKnot.xIndex++
		} else {
			followingKnot.xIndex--
		}

		return followingKnot
	}

	if leadingKnot.yIndex > followingKnot.yIndex {
		followingKnot.yIndex++
	} else {
		followingKnot.yIndex--
	}
	if leadingKnot.xIndex > followingKnot.xIndex {
		followingKnot.xIndex++
	} else {
		followingKnot.xIndex--
	}

	return followingKnot
}

func getIntDiff(a int, b int) int {
	return int(math.Max(float64(a), float64(b)) - math.Min(float64(a), float64(b)))
}
