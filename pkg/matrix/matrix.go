package matrix

import (
	"math"
)

type Point struct {
	X int
	Y int
}

func Find(matrix [][]string, value string) Point {
	for yi, line := range matrix {
		for xi, v := range line {
			if v == value {
				return Point{X: xi, Y: yi}
			}
		}
	}

	return Point{}
}

func CountOccurrences(matrix [][]string, value string) int {
	count := 0
	for _, line := range matrix {
		for _, v := range line {
			if v == value {
				count += 1
			}
		}
	}

	return count
}

func Print(matrix [][]string) {
	for _, line := range matrix {
		for _, v := range line {
			print(v)
		}
		print("\n")
	}
}

func Distance(p1 Point, p2 Point) int {
	xDiff := int(math.Abs(float64(p1.X - p2.X)))
	yDiff := int(math.Abs(float64(p1.Y - p2.Y)))

	return xDiff + yDiff
}

// Gets all points just out of distance
func GetBorderPoints(p Point, distance int) []Point {
	points := []Point{}

	// Starting at mosth top point
	x := p.X
	xStart := p.X
	// Get just out of reach from p
	y := p.Y + distance + 1
	yStart := p.Y + distance + 1

	// Control movement of next point
	xIncreasing := true
	yIncreasing := false

	// Working around clock wise
	for {
		newPoint := Point{X: x, Y: y}
		points = append(points, newPoint)

		if xIncreasing {
			x = x + 1
		} else {
			x = x - 1
		}
		if yIncreasing {
			y = y + 1
		} else {
			y = y - 1
		}

		if x == p.X {
      yIncreasing = !yIncreasing
		}

		if y == p.Y {
      xIncreasing = !xIncreasing
		}

    if x == xStart && y == yStart {
      break
    }
	}

	return points
}
