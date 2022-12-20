package day16

import (
	"bufio"
	"log"

	"golang.org/x/exp/slices"
)

func Day16Part1(r *bufio.Reader) int {
	sections := createSectionsFromReader(r)
	foundI := slices.IndexFunc(sections, func(s section) bool { return s.name == "AA" })
	currenctSection := sections[foundI]

  scores := runRound(1,0, sections, currenctSection)

  topScore := 0
  for _, score := range scores {
    if score > topScore {
      topScore = score
    }
  }

  return topScore
}

func Day16Part1Smart(r *bufio.Reader) int {
	sections := createSectionsFromReader(r)

	presureReleased := 0

	foundI := slices.IndexFunc(sections, func(s section) bool { return s.name == "AA" })
	currenctSection := sections[foundI]
	for minute := 1; minute <= 30; minute++ {
		log.Println("")
		log.Println("Minute: ", minute)
		for _, s := range sections {
			if s.isOpen {
				log.Printf("%v is open, releasing %v \n", s.name, s.rate)
				presureReleased += s.rate
			}
		}

		sectionToGoTo := ""
		topSectionScore := 0
		for _, s := range sections {
			if s.isOpen {
				continue
			}
			d, path := distanceToSection(
				sections,
				currenctSection.name,
				s.name,
			)
			// add one as it take 1 turn to open a valve
			d += 1

			sectionScore := s.rate / d
			if sectionScore > topSectionScore {
				topSectionScore = sectionScore
				if len(path) == 1 {
					sectionToGoTo = path[0]
				} else {
					sectionToGoTo = path[1]
				}
			}
		}

		foundI := slices.IndexFunc(sections, func(s section) bool { return s.name == sectionToGoTo })
		if foundI >= 0 {
			if sectionToGoTo == currenctSection.name {
				log.Println("Opening: ", sections[foundI].name)
				sections[foundI].isOpen = true
			} else {
				log.Println("Moving to: ", sections[foundI].name)
				currenctSection = sections[foundI]
			}
		}

	}

	return presureReleased
}

func Day16Part2(r *bufio.Reader) int {
	return 0
}
