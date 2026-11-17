// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// =============================================================================
// Fuzzy Matching
// =============================================================================

// Fuzzy highlights matching characters in text based on query.
// Greedy left-to-right subsequence matching is used, so any substring match
// is found this way and no fallback pass is needed.
func Fuzzy(text, query string, normalStyle, matchStyle lipgloss.Style) string {
	if query == "" {
		return normalStyle.Render(text)
	}

	queryRunes := []rune(strings.ToLower(query))

	var result strings.Builder
	queryIdx := 0

	for _, char := range text {
		if queryIdx < len(queryRunes) && strings.ToLower(string(char)) == string(queryRunes[queryIdx]) {
			result.WriteString(matchStyle.Render(string(char)))
			queryIdx++
		} else {
			result.WriteString(normalStyle.Render(string(char)))
		}
	}

	return result.String()
}
