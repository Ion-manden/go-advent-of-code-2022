package main

import (
	"advent-of-code/day1"
	"advent-of-code/day10"
	"advent-of-code/day11"
	"advent-of-code/day12"

	"advent-of-code/day14"
	"advent-of-code/day15"
	"advent-of-code/day16"
	"advent-of-code/day2"
	"advent-of-code/day3"
	"advent-of-code/day4"
	"advent-of-code/day5"
	"advent-of-code/day6"
	"advent-of-code/day7"
	"advent-of-code/day8"
	"advent-of-code/day9"
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
	log.Println("Day 4 Part 1:", day4.Day4Part1(reader.CreateReaderFromFile("./day4/input.txt")))
	log.Println("Day 4 Part 2:", day4.Day4Part2(reader.CreateReaderFromFile("./day4/input.txt")))
	log.Println("Day 5 Part 1:", day5.Day5Part1(reader.CreateReaderFromFile("./day5/input.txt")))
	log.Println("Day 5 Part 2:", day5.Day5Part2(reader.CreateReaderFromFile("./day5/input.txt")))
	log.Println("Day 6 Part 1:", day6.Day6Part1(reader.CreateReaderFromFile("./day6/input.txt")))
	log.Println("Day 6 Part 2:", day6.Day6Part2(reader.CreateReaderFromFile("./day6/input.txt")))
	log.Println("Day 7 Part 1:", day7.Day7Part1(reader.CreateReaderFromFile("./day7/input.txt")))
	log.Println("Day 7 Part 2:", day7.Day7Part2(reader.CreateReaderFromFile("./day7/input.txt")))
	log.Println("Day 8 Part 1:", day8.Day8Part1(reader.CreateReaderFromFile("./day8/input.txt")))
	log.Println("Day 8 Part 2:", day8.Day8Part2(reader.CreateReaderFromFile("./day8/input.txt")))
	log.Println("Day 9 Part 1:", day9.Day9Part1(reader.CreateReaderFromFile("./day9/input.txt")))
	log.Println("Day 9 Part 2:", day9.Day9Part2(reader.CreateReaderFromFile("./day9/input.txt")))
	log.Println("Day 10 Part 1:", day10.Day10Part1(reader.CreateReaderFromFile("./day10/input.txt")))
	log.Println("Day 10 Part 2:")
	day10.Day10Part2(reader.CreateReaderFromFile("./day10/input.txt"))
	log.Println("Day 11 Part 1:", day11.Day11Part1(reader.CreateReaderFromFile("./day11/input.txt")))
	log.Println("Day 11 Part 2:", day11.Day11Part2(reader.CreateReaderFromFile("./day11/input.txt")))
	log.Println("Day 12 Part 1:", day12.Day12Part1(reader.CreateReaderFromFile("./day12/input.txt")))
	log.Println("Day 12 Part 2:", day12.Day12Part2(reader.CreateReaderFromFile("./day12/input.txt")))

	log.Println("Day 14 Part 1:", day14.Day14Part1(reader.CreateReaderFromFile("./day14/input.txt")))
	log.Println("Day 14 Part 2:", day14.Day14Part2(reader.CreateReaderFromFile("./day14/input.txt")))
	log.Println("Day 15 Part 1:", day15.Day15Part1(reader.CreateReaderFromFile("./day15/input.txt"), 2000000))
	log.Println("Day 15 Part 2:", day15.Day15Part2(reader.CreateReaderFromFile("./day15/input.txt"), 4000000))
	log.Println("Day 16 Part 1:", day16.Day16Part1(reader.CreateReaderFromFile("./day16/input.txt")))
	log.Println("Day 16 Part 2:", day16.Day16Part2(reader.CreateReaderFromFile("./day16/input.txt")))
}
