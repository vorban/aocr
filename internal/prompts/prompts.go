package prompts

import (
	"fmt"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/vorban/aocr/internal/aoc"
)

func getAocYears() []int {
	now := time.Now()
	lastYear := now.Year()

	if now.Month() < time.December {
		lastYear--
	}

	years := []int{}
	for y := aoc.FIRST_YEAR; y <= lastYear; y++ {
		years = append(years, y)
	}

	return years
}

func makeStarValidator(maxStars int) func(string) error {
	return func(input string) error {
		n, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("must be a number")
		}
		if n < 0 || n > maxStars {
			return fmt.Errorf("must be between 0 and %d", maxStars)
		}

		return nil
	}
}

func promptStarsForYear(year int) int {
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("Stars solved in %d", year),
		Validate: makeStarValidator(aoc.GetMaxStarsPerYear(year)),
		Default:  "0",
	}

	for {
		input, err := prompt.Run()
		if err == nil {
			i, _ := strconv.Atoi(input)
			return i
		}
	}
}

func PromptStars() []aoc.AocYear {
	var starsPerYear []aoc.AocYear
	years := getAocYears()

	for _, year := range years {
		stars := promptStarsForYear(year)
		starsPerYear = append(starsPerYear, aoc.AocYear{Year: year, Stars: aoc.GetMaxStarsPerYear(year), Obtained: stars})
	}

	return starsPerYear
}
