package day2

import "log"

func getHandScore(h Hand) int {
	switch h {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}

	log.Println("Invalid hand", h)
	return 0
}

func getWinScore(l Hand, r Hand) int {
	if l == r {
		return 3
	}

	if l == Rock {
		if r == Paper {
      return 6
		}
		if r == Scissors {
			return 0
		}
	}

	if l == Paper {
		if r == Rock {
			return 0
		}
		if r == Scissors {
      return 6
		}
	}

	if l == Scissors {
		if r == Rock {
			return 6
		}
		if r == Paper {
			return 0
		}
	}

  log.Println("Could not find winner", l, r)
  return 0
}
