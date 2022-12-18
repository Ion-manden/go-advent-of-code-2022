package day15

import (
	"advent-of-code/pkg/matrix"
	"bufio"
)

func Day15Part1(r *bufio.Reader, y int) int {
	sbPairs := createSensorBeaconPairsFromReader(r)

	xMin, xMax := getXMinMax(sbPairs)

	pointsCovered := 0
	// TODO: reduce scan surface
	for xi := xMin - 10000000; xi <= xMax+10000000; xi++ {
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

func Day15Part2(r *bufio.Reader, maxSearchIndex int) int {
	sbPairs := createSensorBeaconPairsFromReader(r)
	distressBeacon := matrix.Point{}

	potentialPositions := []matrix.Point{}

	for _, sbPair := range sbPairs {
		points := matrix.GetBorderPoints(sbPair.sensor, sbPair.distance())

		potentialPositions = append(potentialPositions, points...)
	}


  ppLoop:
  for _, p := range potentialPositions {
    if p.X < 0 || p.Y < 0 {
      continue
    }
    if p.X > maxSearchIndex || p.Y > maxSearchIndex {
      continue
    }

		for _, sbPair := range sbPairs {
			if matrix.Distance(p, sbPair.sensor) <= sbPair.distance() {
				continue ppLoop
			}
		}
		distressBeacon = p
    break
	}

	return (distressBeacon.X * 4000000) + distressBeacon.Y
}
