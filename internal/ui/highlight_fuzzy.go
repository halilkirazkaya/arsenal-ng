// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains fuzzy highlighting functionality that highlights matching
// characters in text based on a search query. It provides visual feedback
// for search results in the UI.
package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// =============================================================================
// Fuzzy Highlighting
// =============================================================================

// fuzzyHighlight highlights matching characters in text based on query.
func fuzzyHighlight(text, query string, normalStyle, matchStyle lipgloss.Style) string {
	if query == "" {
		return normalStyle.Render(text)
	}

	query = strings.ToLower(query)
	textLower := strings.ToLower(text)

	var result strings.Builder
	queryIdx := 0

	for _, char := range text {
		if queryIdx < len(query) && strings.ToLower(string(char)) == string(query[queryIdx]) {
			result.WriteString(matchStyle.Render(string(char)))
			queryIdx++
		} else {
			result.WriteString(normalStyle.Render(string(char)))
		}
	}

	// Fallback to substring match if fuzzy didn't complete
	if queryIdx < len(query) && strings.Contains(textLower, query) {
		idx := strings.Index(textLower, query)
		if idx >= 0 {
			return normalStyle.Render(text[:idx]) +
				matchStyle.Render(text[idx:idx+len(query)]) +
				normalStyle.Render(text[idx+len(query):])
		}
	}

	return result.String()
}

