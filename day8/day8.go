package day8

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func Day8Part1(r *bufio.Reader) int {
	forest := [][]int{}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		treeLine := []int{}

		line := string(b)
		trees := strings.Split(line, "")
		for _, t := range trees {
			tHight, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal("Could not convert to int: ", tHight)
			}
			treeLine = append(treeLine, tHight)
		}
		forest = append(forest, treeLine)
	}

	visibleTrees := 0
	forestLength := len(forest)
	forestWidth := len(forest[0])

	for iRow, treeLine := range forest {
		for iCol, treeHeight := range treeLine {
			// Check if on edge of forest
			if iRow == 0 || iCol == 0 || iRow == forestLength-1 || iCol == forestWidth-1 {
				visibleTrees++
				continue
			}

			hasBlockingTreeTop := false
			hasBlockingTreeBottom := false
			hasBlockingTreeLeft := false
			hasBlockingTreeRight := false
			for r := 0; r < forestLength; r++ {
				for c := 0; c < forestWidth; c++ {
					if r != iRow && c != iCol {
						continue
					}
					th := forest[r][c]

					if th >= treeHeight {
						if r < iRow {
							hasBlockingTreeLeft = true
						}
						if r > iRow {
							hasBlockingTreeRight = true
						}
						if c < iCol {
							hasBlockingTreeTop = true
						}
						if c > iCol {
							hasBlockingTreeBottom = true
						}
					}
				}
			}

			if !hasBlockingTreeTop || !hasBlockingTreeBottom || !hasBlockingTreeLeft || !hasBlockingTreeRight {
				visibleTrees++
			}
		}
	}

	return visibleTrees
}

func Day8Part2(r *bufio.Reader) int {
	forest := [][]int{}

	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		treeLine := []int{}

		line := string(b)
		trees := strings.Split(line, "")
		for _, t := range trees {
			tHight, err := strconv.Atoi(t)
			if err != nil {
				log.Fatal("Could not convert to int: ", tHight)
			}
			treeLine = append(treeLine, tHight)
		}
		forest = append(forest, treeLine)
	}

	visibleTreesFromBestTreeHouse := 0
	forestLength := len(forest)
	forestWidth := len(forest[0])

	for iRow, treeLine := range forest {
		for iCol, treeHeight := range treeLine {
			treesViewableTop := 0
			for ic := iCol - 1; ic >= 0; ic-- {
				th := forest[iRow][ic]
				if th >= treeHeight {
					treesViewableTop++
					break
				}
				treesViewableTop++
			}

			treesViewableBottom := 0
			for ic := iCol + 1; ic < forestLength; ic++ {
				th := forest[iRow][ic]
				if th >= treeHeight {
					treesViewableBottom++
					break
				}
				treesViewableBottom++
			}

			treesViewableLeft := 0
			for ir := iRow - 1; ir >= 0; ir-- {
				th := forest[ir][iCol]
				if th >= treeHeight {
					treesViewableLeft++
					break
				}
				treesViewableLeft++
			}

			treesViewableRight := 0
			for ir := iRow + 1; ir < forestWidth; ir++ {
				th := forest[ir][iCol]
				if th >= treeHeight {
					treesViewableRight++
					break
				}
				treesViewableRight++
			}

			viewScore := treesViewableTop * treesViewableBottom * treesViewableLeft * treesViewableRight

			if viewScore > visibleTreesFromBestTreeHouse {
				visibleTreesFromBestTreeHouse = viewScore
			}
		}
	}

	return visibleTreesFromBestTreeHouse
}
