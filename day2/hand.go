package day2

import "log"

type Hand = string

const (
	Rock     Hand = "rock"
	Paper         = "paper"
	Scissors      = "scissors"
)

func getLeftHand(c string) Hand {
	switch c {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	}

	log.Println("Invalid left char", c)
	return ""
}

func getRightHand(c string) Hand {
	switch c {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	}

	log.Println("Invalid right char", c)
	return ""
}
