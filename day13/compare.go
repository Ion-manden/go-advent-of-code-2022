package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getPartsFromString(line string) []string {
	// Clean Start and ending brackets
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")
	if line == "" {
		return []string{}
	}

	parts := []string{}

	chars := strings.Split(line, "")
	nestLevel := 0
	part := ""
	for _, char := range chars {
		if char == "[" {
			nestLevel += 1
		}
		if char == "]" {
			nestLevel -= 1
		}

		if char == "," && nestLevel == 0 {
			parts = append(parts, part)
			part = ""
			continue
		}

		part = fmt.Sprint(part, char)
	}

	parts = append(parts, part)

	return parts
}

func getNumber(part string, index int, nestingLevel int) (int, int, int) {
	parts := getPartsFromString(part)

	if len(parts) == 0 {
		return 0, index, nestingLevel
	}

	if len(parts) <= index {
		return 0, index, nestingLevel
	}

	n, err := strconv.Atoi(parts[index])
	if err != nil {
		log.Printf("%+v is not a number, looking deeper \n", parts[0])
		return getNumber(parts[0], index, nestingLevel+1)
	}

	return n, index, nestingLevel
}

func isPacketsInRightOrder(left string, right string) bool {
	leftParts := getPartsFromString(left)
	rightParts := getPartsFromString(right)

	for i := 0; i < len(left)+len(right); i++ {
		if len(leftParts) == i && len(rightParts) > i {
			return true
		}
		if len(leftParts) > i && len(rightParts) == i {
			return false
		}

		lp := leftParts[i]
		rp := rightParts[i]

		for i := 0; i < 20; i++ {
			l, _, lnl := getNumber(lp, i, 0)
			r, _, rnl := getNumber(rp, i, 0)

			if l == r && lnl > rnl {
				return false
			}

			if l == r && lnl < rnl {
				return true
			}

			if l < r {
				return true
			}

			if l > r {
				return false
			}
		}
	}

	return true
}
