package day11

import (
	"bufio"
	"math"
	"sort"
)

func simulateMonkies(monkies []monkey, rounds int, divider int, useDividerAsMod bool) []monkey {
	for r := 0; r < rounds; r++ {
		for mi, m := range monkies {
			for _, item := range m.items {
				monkies[mi].inspectedItemLog = append(monkies[mi].inspectedItemLog, item)
				worry := m.getWorryWhenHolding(item)

				if useDividerAsMod {
					worry = worry % divider
				} else {
					worry = int(math.Floor(float64(worry) / float64(divider)))
				}

				if m.doesItemWorryPassTest(worry) {
					monkeyNumber := m.getTestTrueMonkey()
					monkies[monkeyNumber].items = append(monkies[monkeyNumber].items, worry)
				} else {
					monkeyNumber := m.getTestFalseMonkey()
					monkies[monkeyNumber].items = append(monkies[monkeyNumber].items, worry)
				}
			}
			monkies[mi].items = []int{}
		}
	}

	return monkies
}

func Day11Part1(r *bufio.Reader) int {
	monkies := getMonkiesFromFile(r)

	// go through 20 rounds
	// Divide by 3 to reduce
	monkies = simulateMonkies(monkies, 20, 3, false)

	itemsInspectedPrMonkey := []int{}
	for _, m := range monkies {
		itemsInspectedPrMonkey = append(itemsInspectedPrMonkey, len(m.inspectedItemLog))
	}

	sort.Ints(itemsInspectedPrMonkey)

	l := len(itemsInspectedPrMonkey)
	return itemsInspectedPrMonkey[l-1] * itemsInspectedPrMonkey[l-2]
}

func Day11Part2(r *bufio.Reader) int {
	monkies := getMonkiesFromFile(r)

	divider := 1
	for _, m := range monkies {
		divider = divider * m.getTestDivideNumber()
	}

	// go through 10_000 rounds
	// Divide by divider to reduce number
	monkies = simulateMonkies(monkies, 10_000, divider, true)

	itemsInspectedPrMonkey := []int{}
	for _, m := range monkies {
		itemsInspectedPrMonkey = append(itemsInspectedPrMonkey, len(m.inspectedItemLog))
	}

	sort.Ints(itemsInspectedPrMonkey)

	l := len(itemsInspectedPrMonkey)
	return itemsInspectedPrMonkey[l-1] * itemsInspectedPrMonkey[l-2]
}
