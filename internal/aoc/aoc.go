package aoc

const FIRST_YEAR = 2015

type AocYear struct {
	Year     int
	Stars    int
	Obtained int
}

func GetMaxStarsPerYear(year int) int {
	if year < 2025 {
		return 50
	}

	return 24
}
