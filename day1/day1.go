package day1

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strconv"
)

func Day1Part1(r *bufio.Reader) int {
	highest := 0
	sum := 0

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			if sum > highest {
				highest = sum
			}
			sum = 0
			break
		}
		if err != nil {
			break
		}

		line := string(b)

		if line == "" {
			if sum > highest {
				highest = sum
			}
			sum = 0
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}

	return highest
}

func Day1Part2(r *bufio.Reader) int {
	sums := []int{}
	sum := 0

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			sums = append(sums, sum)
			sum = 0
			break
		}
		if err != nil {
			break
		}

		line := string(b)

		if line == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}

	sort.Ints(sums)

	result := 0
	for _, v := range sums[len(sums)-3:] {
		result += v
	}

	return result
}
