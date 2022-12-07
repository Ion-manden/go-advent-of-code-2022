package day7

import (
	"bufio"
	"regexp"
)

type file struct {
	name       string
	size       int
	folderPath string
	path       string
}

type folder struct {
	name         string
	parentFolder string
	path         string
}

var removeLastSection *regexp.Regexp = regexp.MustCompile(`/\w+$`)

var folders = []folder{}
var files = []file{}

func Day7Part1(r *bufio.Reader) int {
	ohLord(r)

	totalSize := 0

	for _, f := range folders {
		size := getFolderSize(f.path)
		if size > 100_000 {
			continue
		}

		totalSize += size
	}

	return totalSize
}

func Day7Part2(r *bufio.Reader) int {
	ohLord(r)

	totalSpaceUsed := getFolderSize("/")
	freeSpace := 70000000 - totalSpaceUsed
	spaceNeeded := 30000000 - freeSpace

	smallestSeen := totalSpaceUsed

	for _, f := range folders {
		size := getFolderSize(f.path)

		if size < spaceNeeded || size > smallestSeen {
			continue
		}

		smallestSeen = size
	}

	return smallestSeen
}
