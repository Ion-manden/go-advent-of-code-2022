package day14

import (
	"advent-of-code/pkg/matrix"
	"fmt"
)

// Returns updated map and an error if there were no more space for sand
func dropSand(rockMap [][]string) ([][]string, error) {
	sandSource := matrix.Find(rockMap, "+")

	sand := sandSource

	if rockMap[sand.Y+1][sand.X] != "." &&
		rockMap[sand.Y+1][sand.X-1] != "." &&
		rockMap[sand.Y+1][sand.X+1] != "." {
		rockMap[sand.Y][sand.X] = "O"

		return rockMap, fmt.Errorf("No more space")
	}

	for {
		if sand.Y+1 == len(rockMap) {
			return rockMap, fmt.Errorf("No more space")
		}

		if rockMap[sand.Y+1][sand.X] == "." {
			sand.Y = sand.Y + 1
			continue
		}

		if sand.X == 0 {
			return rockMap, fmt.Errorf("No more space")
		}

		if rockMap[sand.Y+1][sand.X] != "." && rockMap[sand.Y+1][sand.X-1] == "." {
			sand.Y = sand.Y + 1
			sand.X = sand.X - 1
			continue
		}
		if rockMap[sand.Y+1][sand.X] != "." && rockMap[sand.Y+1][sand.X+1] == "." {
			sand.Y = sand.Y + 1
			sand.X = sand.X + 1
			continue
		}

		break
	}

	rockMap[sand.Y][sand.X] = "O"

	return rockMap, nil
}
