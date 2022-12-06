package day6

import (
	"bufio"
	"io"
)

func shift[T comparable](slice []T, item T) []T {
	return append([]T{item}, slice[:len(slice)-1]...)
}

func isUniq[T comparable](slice []T) bool {
	for i, v := range slice {
		for _, c := range slice[i+1:] {
			if v == c {
				return false
			}
		}
	}

	return true
}

func Day6Part1(r *bufio.Reader) int {
	charsProcessed := 0
	recentRunes := []string{}
	for {
		charsProcessed++

		r, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		c := string(r)

		if len(recentRunes) < 4 {
			// Add to beging of slice
			recentRunes = append([]string{c}, recentRunes...)
		} else {
			recentRunes = shift(recentRunes, c)
		}

		if len(recentRunes) == 4 && isUniq(recentRunes) {
			return charsProcessed
		}
	}

	return 0
}

func Day6Part2(r *bufio.Reader) int {
	charsProcessed := 0
	recentRunes := []string{}
	for {
		charsProcessed++

		r, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		c := string(r)

		if len(recentRunes) < 14 {
			// Add to beging of slice
			recentRunes = append([]string{c}, recentRunes...)
		} else {
			recentRunes = shift(recentRunes, c)
		}

		if len(recentRunes) == 14 && isUniq(recentRunes) {
			return charsProcessed
		}
	}

	return 0
}
