package main

import (
	"advent-of-code/day1"
	"advent-of-code/day2"
	"advent-of-code/day3"
	"advent-of-code/reader"
	"log"
)

func main() {
	log.Println("Day 1 Part 1:", day1.Day1Part1(reader.CreateReaderFromFile("./day1/input.txt")))
	log.Println("Day 1 Part 2:", day1.Day1Part2(reader.CreateReaderFromFile("./day1/input.txt")))
	log.Println("Day 2 Part 1:", day2.Day2Part1(reader.CreateReaderFromFile("./day2/input.txt")))
	log.Println("Day 2 Part 2:", day2.Day2Part2(reader.CreateReaderFromFile("./day2/input.txt")))
	log.Println("Day 3 Part 1:", day3.Day3Part1(reader.CreateReaderFromFile("./day3/input.txt")))
	log.Println("Day 3 Part 2:", day3.Day3Part2(reader.CreateReaderFromFile("./day3/input.txt")))
}
