package day15

import (
	"advent-of-code/pkg/matrix"
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type sensorBeaconPair struct {
	sensor matrix.Point
	beacon matrix.Point
}

func (sbPair *sensorBeaconPair) distance() int {
	return matrix.Distance(sbPair.sensor, sbPair.beacon)
}

func createSensorBeaconPairsFromReader(r *bufio.Reader) []sensorBeaconPair {
	SBPairs := []sensorBeaconPair{}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		sensor, beacon := getSensorAndBeacon(line)
		SBPairs = append(SBPairs, sensorBeaconPair{sensor: sensor, beacon: beacon})
	}

	return SBPairs
}

func getSensorAndBeacon(line string) (matrix.Point, matrix.Point) {
	parts := strings.Split(line, ": ")
	sensorParts := strings.Split(strings.Replace(parts[0], "Sensor at ", "", 1), ", ")
	beaconParts := strings.Split(strings.Replace(parts[1], "closest beacon is at ", "", 1), ", ")

	sensorX, err := strconv.Atoi(strings.Replace(sensorParts[0], "x=", "", 1))
	if err != nil {
		log.Fatal("Could not convert to int: ", sensorParts[0])
	}
	sensorY, err := strconv.Atoi(strings.Replace(sensorParts[1], "y=", "", 1))
	if err != nil {
		log.Fatal("Could not convert to int: ", sensorParts[1])
	}

	beaconX, err := strconv.Atoi(strings.Replace(beaconParts[0], "x=", "", 1))
	if err != nil {
		log.Fatal("Could not convert to int: ", beaconParts[0])
	}
	beaconY, err := strconv.Atoi(strings.Replace(beaconParts[1], "y=", "", 1))
	if err != nil {
		log.Fatal("Could not convert to int: ", beaconParts[1])
	}

	return matrix.Point{X: sensorX, Y: sensorY}, matrix.Point{X: beaconX, Y: beaconY}
}

func getXMinMax(sbPairs []sensorBeaconPair) (int, int) {
	min := 0
	max := 0

	for _, sbPair := range sbPairs {
		if sbPair.sensor.X < min {
			min = sbPair.sensor.X
		}
		if sbPair.sensor.X > max {
			max = sbPair.sensor.X
		}

		if sbPair.beacon.X < min {
			min = sbPair.beacon.X
		}
		if sbPair.beacon.X > max {
			max = sbPair.beacon.X
		}
	}

	return min, max
}

func getXMinMaxCoverage(sbPairs []sensorBeaconPair) (int, int) {
	min := 0
	max := 0

	for _, sbPair := range sbPairs {
		if sbPair.sensor.X < min {
			min = sbPair.sensor.X
		}
		if sbPair.sensor.X > max {
			max = sbPair.sensor.X
		}

		if sbPair.beacon.X < min {
			min = sbPair.beacon.X
		}
		if sbPair.beacon.X > max {
			max = sbPair.beacon.X
		}
	}

	return min, max
}
