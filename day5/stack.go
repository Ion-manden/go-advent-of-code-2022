package day5

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parseLineAsSections(line string) []string {
	sections := []string{}

	chars := strings.Split(line, "")

	for i, c := range chars {
		if i%4 != 1 {
			continue
		}

		val := c
		if val == " " {
			val = ""
		}

		sections = append(sections, val)
	}

	return sections
}

// move 1 from 2 to 1
func moveStack(operation string, stack [][]string) [][]string {
	re := regexp.MustCompile(` \d+`)

	matches := re.FindAll([]byte(operation), 3)
	nToMove, err := strconv.Atoi(strings.Replace(string(matches[0]), " ", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	fromStack, err := strconv.Atoi(strings.Replace(string(matches[1]), " ", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	fromStack -= 1
	toStack, err := strconv.Atoi(strings.Replace(string(matches[2]), " ", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	toStack -= 1

	for n := 0; n < nToMove; n++ {
		item := stack[fromStack][len(stack[fromStack])-1]
		stack[fromStack] = stack[fromStack][:len(stack[fromStack])-1]
		stack[toStack] = append(stack[toStack], item)
	}

	return stack
}

func moveStackBatch(operation string, stack [][]string) [][]string {
	re := regexp.MustCompile(` \d+`)

	matches := re.FindAll([]byte(operation), 3)
	nToMove, err := strconv.Atoi(strings.Replace(string(matches[0]), " ", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	fromStack, err := strconv.Atoi(strings.Replace(string(matches[1]), " ", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	fromStack -= 1
	toStack, err := strconv.Atoi(strings.Replace(string(matches[2]), " ", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	toStack -= 1

  items := stack[fromStack][len(stack[fromStack])-nToMove:]
	stack[fromStack] = stack[fromStack][:len(stack[fromStack])-nToMove]
	stack[toStack] = append(stack[toStack], items...)

	return stack
}
