package day2

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func Day2Part1(r *bufio.Reader) int {
	score := 0

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}

		line := string(b)

		p := strings.Split(line, " ")
		left := getLeftHand(p[0])
		right := getRightHand(p[1])

		score += getWinScore(left, right) + getHandScore(right)
	}

	return score
}

func Day2Part2(r *bufio.Reader) int {
	score := 0

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}

		line := string(b)

		p := strings.Split(line, " ")
		left := getLeftHand(p[0])
		var right Hand = ""

		howRoundShouldEnd := p[1]
		// Lose
		if howRoundShouldEnd == "X" {
			switch left {
			case Rock:
				right = Scissors
			case Paper:
				right = Rock
			case Scissors:
				right = Paper
			}

		}
		// Tie
		if howRoundShouldEnd == "Y" {
			right = left
		}
		// Win
		if howRoundShouldEnd == "Z" {
			switch left {
			case Rock:
        right = Paper
			case Paper:
        right = Scissors
			case Scissors:
        right = Rock
			}
		}

		score += getWinScore(left, right) + getHandScore(right)
	}

	return score
}
