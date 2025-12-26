// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains string utility functions including truncate() for
// shortening strings with ellipsis, wordWrap() for wrapping text at word
// boundaries, and min() helper function.
package ui

import (
	"strings"
)

// =============================================================================
// String Utilities
// =============================================================================

// truncate shortens a string to max length with ellipsis.
func truncate(s string, max int) string {
	s = strings.ReplaceAll(s, "\n", " ")

	if max <= 0 {
		return ""
	}
	if max <= 3 {
		return s[:min(len(s), max)]
	}
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

// wordWrap wraps text to the specified width, breaking at word boundaries.
// It does not add ellipsis and preserves the full text across multiple lines.
func wordWrap(text string, width int) string {
	if width <= 0 {
		return text
	}

	// Replace newlines with spaces first
	text = strings.ReplaceAll(text, "\n", " ")

	// If text fits in one line, return as is
	if len(text) <= width {
		return text
	}

	var result strings.Builder
	words := strings.Fields(text)
	currentLine := ""

	for _, word := range words {
		// If adding this word would exceed width, start a new line
		testLine := currentLine
		if testLine != "" {
			testLine += " "
		}
		testLine += word

		if len(testLine) > width && currentLine != "" {
			// Current line is full, start new line
			result.WriteString(currentLine)
			result.WriteString("\n")
			currentLine = word
		} else {
			// Add word to current line
			currentLine = testLine
		}
	}

	// Add the last line
	if currentLine != "" {
		result.WriteString(currentLine)
	}

	return result.String()
}

// min returns the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

