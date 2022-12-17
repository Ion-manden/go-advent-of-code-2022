package day14

import (
	"bufio"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

func createRockMapFromFile(r *bufio.Reader, trim bool) [][]string {
	rockMap := [][]string{}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		points := strings.Split(line, " -> ")

		for _, point := range points {
			parts := strings.Split(point, ",")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal("Failed to convert to int: ", parts[0])
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("Failed to convert to int: ", parts[1])
			}

			logestLine := 0
			for _, line := range rockMap {
				l := len(line)
				if l > logestLine {
					logestLine = l
				}
			}

			for yi := 0; yi <= int(math.Max(float64(y), float64(len(rockMap)-1))); yi++ {
				if len(rockMap) == yi {
					rockMap = append(rockMap, []string{})
				}

				for xi := 0; xi <= int(math.Max(float64(x), float64(logestLine-1))); xi++ {
					if len(rockMap[yi]) == xi {
						rockMap[yi] = append(rockMap[yi], ".")
					}
				}
			}
		}

		for pi := range points {
			// Skip first as we look back
			if pi == 0 {
				continue
			}

			parts := strings.Split(points[pi-1], ",")
			px, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal("Failed to convert to int: ", parts[0])
			}
			py, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("Failed to convert to int: ", parts[1])
			}

			parts = strings.Split(points[pi], ",")
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatal("Failed to convert to int: ", parts[0])
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("Failed to convert to int: ", parts[1])
			}

			if px == x {
				distance := int(math.Abs(float64(py - y)))
				direction := "UP"
				if py > y {
					direction = "DOWN"
				}

				for i := 0; i < distance+1; i++ {
					rockX := px
					rockY := py
					if direction == "UP" {
						rockY = rockY + i
					}
					if direction == "DOWN" {
						rockY -= i
					}

					rockMap[rockY][rockX] = "#"
				}
			}

			if py == y {
				distance := int(math.Abs(float64(px - x)))
				direction := "RIGHT"
				if px > x {
					direction = "LEFT"
				}

				for i := 0; i < distance+1; i++ {
					rockX := px
					rockY := py
					if direction == "RIGHT" {
						rockX = rockX + i
					}
					if direction == "LEFT" {
						rockX -= i
					}
					rockMap[rockY][rockX] = "#"
				}
			}
		}
	}

	// Place start sand source
	rockMap[0][500] = "+"

	// Cleanup map
	if trim {
		for xi := range rockMap[0] {
			hasRock := false
			for _, line := range rockMap {
				if line[xi] == "#" {
					hasRock = true
					break
				}
			}
			if !hasRock {
				for yi := range rockMap {
					rockMap[yi][xi] = rockMap[yi][xi][1:]
				}
			}
		}
	}

	for yi, line := range rockMap {
		rockMap[yi] = filterValueOut(line, "")
	}

	return rockMap
}

func filterValueOut[T comparable](s []T, filterValue T) []T {
	filtered := []T{}
	for _, item := range s {
		if item != filterValue {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
