package matrix

import "math"

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
