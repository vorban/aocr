package readme

import (
	"fmt"
	"os"
	"strings"
)

const MARKER_START = "<!-- AOCR_BADGES_START -->"
const MARKER_END = "<!-- AOCR_BADGES_END -->"

func FormatBadges(totalBadge string, yearBadges []string) string {
	lines := []string{totalBadge}
	for i := 0; i < len(yearBadges); i += 5 {
		chunk := i + 5
		if chunk > len(yearBadges) {
			chunk = len(yearBadges)
		}

		lines = append(lines, strings.Join(yearBadges[i:chunk], "\n"))
	}

	return fmt.Sprintf("<div>\n%s\n</div>\n", strings.Join(lines, "\n<br>\n"))
}

func UpdateReadme(serialized string) error {
	filename := "README.md"

	var content string
	exists := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exists = false
	}

	if !exists {
		// Create README with badges at top and # readme after
		content = fmt.Sprintf("%s\n%s\n%s\n\n# README\n", MARKER_START, serialized, MARKER_END)
	} else {
		data, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		content = string(data)

		startIdx := strings.Index(content, MARKER_START)
		endIdx := strings.Index(content, MARKER_END)

		if startIdx == -1 || endIdx == -1 || endIdx < startIdx {
			// Markers not found: prepend badges at the top
			content = fmt.Sprintf("%s\n%s\n%s\n\n%s", MARKER_START, serialized, MARKER_END, content)
		} else {
			// Replace existing badge block
			before := content[:startIdx+len(MARKER_START)]
			after := content[endIdx:]
			content = fmt.Sprintf("%s\n%s\n%s", before, serialized, after)
		}
	}

	return os.WriteFile(filename, []byte(content), 0644)
}
