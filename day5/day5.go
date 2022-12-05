package day5

import (
	"bufio"
	"io"
)

func Day5Part1(r *bufio.Reader) string {
	stack := [][]string{}

  // Read until empty line
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		if line == "" {
			break
		}

		sections := parseLineAsSections(line)

		// Remove last line
		if sections[0] == "1" {
			break
		}

		for i, section := range sections {
			if len(stack) <= i {
				if len(stack) <= i {
					stack = append(stack, []string{})
				}
			}
      if section == "" {
        continue
      }

			stack[i] = append([]string{section}, stack[i]...)
		}
	}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		if line == "" {
			continue
		}

		stack = moveStack(line, stack)
	}

  topContainers := ""
  for _, container := range stack {
    topContainers += container[len(container)-1]
  }

  return topContainers
}

func Day5Part2(r *bufio.Reader) string {
	// Read until empty line

	stack := [][]string{}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		if line == "" {
			break
		}

		sections := parseLineAsSections(line)

		// Remove last line
		if sections[0] == "1" {
			break
		}

		for i, section := range sections {
			if len(stack) <= i {
				if len(stack) <= i {
					stack = append(stack, []string{})
				}
			}
      if section == "" {
        continue
      }

			stack[i] = append([]string{section}, stack[i]...)
		}
	}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		line := string(b)
		if line == "" {
			continue
		}

		stack = moveStackBatch(line, stack)
	}

  topContainers := ""
  for _, container := range stack {
    topContainers += container[len(container)-1]
  }

  return topContainers
}
