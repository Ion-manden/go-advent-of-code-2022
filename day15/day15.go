package day15

import (
	"advent-of-code/pkg/matrix"
	"bufio"
)

func Day15Part1(r *bufio.Reader, y int) int {
	sbPairs := createSensorBeaconPairsFromReader(r)

	xMin, xMax := getXMinMax(sbPairs)

	pointsCovered := 0
	for xi := xMin; xi <= xMax; xi++ {
		p := matrix.Point{X: xi, Y: y}

		isCovered := false
		for _, sbPair := range sbPairs {
			// Don't count if there is already a beacon
			if matrix.Distance(p, sbPair.beacon) == 0 {
				break
			}
			if matrix.Distance(p, sbPair.sensor) <= sbPair.distance() {
				isCovered = true
				break
			}
		}

		if isCovered {
			pointsCovered += 1
		}
	}

	return pointsCovered
}

func Day15Part2(r *bufio.Reader) int {
	return 0
}
