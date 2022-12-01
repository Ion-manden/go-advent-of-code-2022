package main

import (
	"advent-of-code/day1"
	"advent-of-code/reader"
	"log"
)

func main() {
	log.Println(day1.Day1Part1(reader.CreateReaderFromFile("./day1/input.txt")))
	log.Println(day1.Day1Part2(reader.CreateReaderFromFile("./day1/input.txt")))
}

