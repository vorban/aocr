package badges

import (
	"fmt"

	"github.com/vorban/aocr/internal/aoc"
)

func BadgeColor(y aoc.AocYear) string {
	switch {
	case y.Obtained == 0:
		return "a8a29e"
	case y.Obtained == y.Stars:
		return "fcd34d"
	default:
		return "e5e5e5"
	}
}

func YearBadge(aocYear aoc.AocYear) string {
	color := BadgeColor(aocYear)
	return fmt.Sprintf(`<img src="https://img.shields.io/badge/%d%%20⭐-%02d%%2F%02d-%s">`,
		aocYear.Year, aocYear.Obtained, aocYear.Stars, color)
}

func TotalStarsBadge(list []aoc.AocYear) string {
	total := 0
	totalObtained := 0

	for _, aocYear := range list {
		totalObtained += aocYear.Obtained
		total += aocYear.Stars
	}

	return fmt.Sprintf(`<img src="https://img.shields.io/badge/total_stars%%20⭐-%02d%%2F%02d-fcd34d?style=for-the-badge">`,
		totalObtained, total)
}
