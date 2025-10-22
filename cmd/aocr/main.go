package main

import (
	"github.com/vorban/aocr/internal/badges"
	"github.com/vorban/aocr/internal/prompts"
	"github.com/vorban/aocr/internal/readme"
)

func main() {
	stars := prompts.PromptStars()

	serialized := []string{}
	for _, s := range stars {
		serialized = append(serialized, badges.YearBadge(s))
	}

	serializedBlock := readme.FormatBadges(badges.TotalStarsBadge(stars), serialized)
	readme.UpdateReadme(serializedBlock)
}
