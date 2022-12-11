package day11

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type monkeyTest struct {
	condition string
	truthy    string
	falsy     string
}

type monkey struct {
	items            []int
	operation        string
	test             monkeyTest
	inspectedItemLog []int
}

func (m *monkey) getWorryWhenHolding(item int) int {
	operationWithoutAssignment := strings.Replace(m.operation, "new = ", "", 1)
	parts := strings.Split(operationWithoutAssignment, " ")
	leftNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		if parts[0] == "old" {
			leftNumber = item
		} else {
			log.Fatal("Could not convert to int: ", parts[0])
		}
	}
	operation := parts[1]
	rightNumber, err := strconv.Atoi(parts[2])
	if err != nil {
		if parts[2] == "old" {
			rightNumber = item
		} else {
			log.Fatal("Could not convert to int: ", parts[2])
		}
	}

	switch operation {
	case "+":
		return leftNumber + rightNumber
	case "-":
		return leftNumber + rightNumber
	case "*":
		return leftNumber * rightNumber
	case "/":
		return leftNumber * rightNumber
	default:
		log.Fatal("operation not handled: ", operation)
		return 0
	}
}

func (m *monkey) getTestDivideNumber() int {
	div := strings.Replace(m.test.condition, "divisible by ", "", 1)

	n, err := strconv.Atoi(div)
	if err != nil {
		log.Fatal("Could not convert to int: ", div)
	}

	return n
}

func (m *monkey) doesItemWorryPassTest(worry int) bool {
  n := m.getTestDivideNumber()

	return worry%n == 0
}

func (m *monkey) getTestTrueMonkey() int {
	mon := strings.Replace(m.test.truthy, "throw to monkey ", "", 1)
	monkeyNumber, err := strconv.Atoi(mon)
	if err != nil {
		log.Println("getTestTrueMonkey")
		log.Fatal("Could not convert to int: ", mon)
	}
	return monkeyNumber
}

func (m *monkey) getTestFalseMonkey() int {
	mon := strings.Replace(m.test.falsy, "throw to monkey ", "", 1)
	monkeyNumber, err := strconv.Atoi(mon)
	if err != nil {
		log.Println("getTestFalseMonkey")
		log.Fatal("Could not convert to int: ", mon)
	}
	return monkeyNumber
}

func getMonkiesFromFile(r *bufio.Reader) []monkey {
	monkies := []monkey{}

	newMonkey := monkey{}
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			// Add last monkey to list
			monkies = append(monkies, newMonkey)
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		// on empty line, go to next monkey
		if line == "" {
			monkies = append(monkies, newMonkey)
			newMonkey = monkey{}
			continue
		}

		// get Starting items
		if strings.HasPrefix(line, "  Starting items:") {
			items := strings.Replace(line, "  Starting items: ", "", 1)
			for _, i := range strings.Split(items, ", ") {
				itemWorryScore, err := strconv.Atoi(i)
				if err != nil {
					log.Fatal("Could not convert to int: ", itemWorryScore)
				}
				newMonkey.items = append(newMonkey.items, itemWorryScore)
			}

			continue
		}

		if strings.HasPrefix(line, "  Operation: ") {
			newMonkey.operation = strings.Replace(line, "  Operation: ", "", 1)
			continue
		}

		if strings.HasPrefix(line, "  Test: ") {
			newMonkey.test.condition = strings.Replace(line, "  Test: ", "", 1)
			continue
		}

		if strings.HasPrefix(line, "    If true: ") {
			newMonkey.test.truthy = strings.Replace(line, "    If true: ", "", 1)
			continue
		}

		if strings.HasPrefix(line, "    If false: ") {
			newMonkey.test.falsy = strings.Replace(line, "    If false: ", "", 1)
			continue
		}
	}

  return monkies
}
