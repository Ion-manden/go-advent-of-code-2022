package day10

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func getSignalStrength(clockCycle int, x int) int {
	if (clockCycle+20)%40 != 0 {
		return 0
	}
	return clockCycle * x
}

func Day10Part1(r *bufio.Reader) int {
	clockCycle := 0
	x := 1
	sumOfSignalStrength := 0

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		if line == "noop" {
			clockCycle++
			sumOfSignalStrength += getSignalStrength(clockCycle, x)
			continue
		}

		if strings.HasPrefix(line, "addx") {
			changeString := strings.Replace(line, "addx ", "", 1)
			change, err := strconv.Atoi(changeString)
			if err != nil {
				log.Fatal("Could not convert string to int: ", changeString)
			}

			// Check signal strength during execution
			clockCycle++
			sumOfSignalStrength += getSignalStrength(clockCycle, x)

			// Check signal strength at the end of execution
			clockCycle++
			sumOfSignalStrength += getSignalStrength(clockCycle, x)

			x = x + change
		}
	}

	return sumOfSignalStrength
}

func getPixelForCycle(clockCycle int, x int) string {
	log.Println("clockCycle: ", clockCycle)
	log.Println("x: ", x)
	if (clockCycle%40)-2 <= x && clockCycle%40 >= x {
		return "#"
	}
	return "."
}

func Day10Part2(r *bufio.Reader) []string {
	clockCycle := 0
	x := 1

	pixels := []string{}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		if line == "noop" {
			clockCycle++
			// Add pixel
			pixels = append(pixels, getPixelForCycle(clockCycle, x))
			continue
		}

		if strings.HasPrefix(line, "addx") {
			changeString := strings.Replace(line, "addx ", "", 1)
			change, err := strconv.Atoi(changeString)
			if err != nil {
				log.Fatal("Could not convert string to int: ", changeString)
			}

			clockCycle++
			// Add pixel during execution
			pixels = append(pixels, getPixelForCycle(clockCycle, x))

			// Add pixel at the end of execution
			clockCycle++
			pixels = append(pixels, getPixelForCycle(clockCycle, x))

			x = x + change
		}
	}

  for _, chunk := range chunkSlice(pixels, 40) {
    row  := strings.Join(chunk, "")
    println(row)
  }

	return pixels
}

func chunkSlice[T comparable](s []T, size int) [][]T {
	chunks := [][]T{}
	chunk := []T{}

	for _, item := range s {
		chunk = append(chunk, item)
		if len(chunk) == size {
			chunks = append(chunks, chunk)
			chunk = []T{}
		}
	}
	// add last chunk
	chunks = append(chunks, chunk)

	return chunks
}
