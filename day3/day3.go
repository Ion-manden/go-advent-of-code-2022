package day3

import (
	"bufio"
	"io"
	"strings"
)

func Day3Part1(r *bufio.Reader) int {
	sum := 0

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)

		chars := strings.Split(line, "")
		firstCompartment := chars[:len(chars)/2]
		secondCompartment := chars[len(chars)/2:]

		charsInBoth := make(map[string]bool)
		for _, fc := range firstCompartment {
			for _, sc := range secondCompartment {
				if fc == sc {
					charsInBoth[fc] = true
					break
				}
			}
		}

		for c, isIn := range charsInBoth {
			if isIn {
				sum += getItemPriority(c)
			}
		}
	}

	return sum
}

func Day3Part2(r *bufio.Reader) int {
	sum := 0

	for {
		b1, _, err := r.ReadLine()
		b2, _, err := r.ReadLine()
		b3, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line1 := string(b1)
		line2 := string(b2)
		line3 := string(b3)

		firstElf := strings.Split(line1, "")
		secondElf := strings.Split(line2, "")
		thirdElf := strings.Split(line3, "")

		charsInFirstAndSecond := make(map[string]bool)
		for _, fc := range firstElf {
			for _, sc := range secondElf {
				if fc == sc {
					charsInFirstAndSecond[fc] = true
				}
			}
		}

    outer:
		for c, isIn := range charsInFirstAndSecond {
			for _, tc := range thirdElf {
				if isIn && c == tc {
					sum += getItemPriority(c)
          break outer
				}
			}
		}
	}

	return sum
}
