package day16

import (
	"golang.org/x/exp/slices"
)

func runRound(round int, presureReleased int, sectionsProvided []section, currenctSection section) []int {
	if round > 30 {
		return []int{presureReleased}
	}

	// Copy sections to not mutate original
	sections := make([]section, len(sectionsProvided))
	copy(sections, sectionsProvided)

	for _, s := range sections {
		if s.isOpen {
			presureReleased += s.rate
		}
	}

	scores := []int{}
	for _, p := range currenctSection.tunnelsTo {
		foundI := slices.IndexFunc(sections, func(s section) bool { return s.name == p })
		nextSection := sections[foundI]

		roundScores := runRound(round+1, presureReleased, sections, nextSection)
		scores = append(scores, roundScores...)
	}

	if !currenctSection.isOpen {
		foundI := slices.IndexFunc(sections, func(s section) bool { return s.name == currenctSection.name })
		sections[foundI].isOpen = true

		roundScores := runRound(round+1, presureReleased, sections, currenctSection)
		scores = append(scores, roundScores...)
	}

	return scores
}
