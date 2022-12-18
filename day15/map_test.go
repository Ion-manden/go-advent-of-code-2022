package day15

import (
	"advent-of-code/pkg/matrix"
	"log"
	"testing"
)

func TestGetSensorAndBeacon(t *testing.T) {
	expectSensor := matrix.Point{X: 2, Y: 18}
	expectBeacon := matrix.Point{X: -2, Y: 15}
	gotSensor, gotBeacon := getSensorAndBeacon("Sensor at x=2, y=18: closest beacon is at x=-2, y=15")

	if expectSensor != gotSensor {
		log.Println("Sensor")
		log.Println("Expected: ", expectSensor)
		log.Println("Got: ", gotSensor)
		t.Fail()
	}
	if expectBeacon != gotBeacon {
		log.Println("Beacon")
		log.Println("Expected: ", expectBeacon)
		log.Println("Got: ", gotBeacon)
		t.Fail()
	}
}
