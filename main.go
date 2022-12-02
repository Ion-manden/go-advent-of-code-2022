package main

import (
	"advent-of-code/day1"
	"advent-of-code/day2"
	"advent-of-code/reader"
	"log"
)

func main() {
  log.Println("Day 1 Part 1:", day1.Day1Part1(reader.CreateReaderFromFile("./day1/input.txt")))
  log.Println("Day 1 Part 2:", day1.Day1Part2(reader.CreateReaderFromFile("./day1/input.txt")))
  log.Println("Day 2 Part 1:", day2.Day2Part1(reader.CreateReaderFromFile("./day2/input.txt")))
  log.Println("Day 2 Part 2:", day2.Day2Part2(reader.CreateReaderFromFile("./day2/input.txt")))
}

