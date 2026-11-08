// Package textutil provides plain string utilities for terminal output:
// truncation with ellipsis and word-boundary wrapping.
package textutil

import (
	"strings"
)

// =============================================================================
// Truncation Helpers
// =============================================================================

// Truncate shortens s to max width, using ellipsis for overflow.
// Cutting is rune-based so multi-byte characters are never split.
func Truncate(s string, max int) string {
	// Newlines would break single-line rows, so flatten them first.
	s = strings.ReplaceAll(s, "\n", " ")

	if max <= 0 {
		return ""
	}
	runes := []rune(s)
	if max <= 3 {
		return string(runes[:min(len(runes), max)])
	}
	if len(runes) <= max {
		return s
	}
	return string(runes[:max-3]) + "..."
}

// =============================================================================
// Wrapping Helpers
// =============================================================================

// Wrap breaks text at word boundaries so lines fit width.
// Widths are measured in runes so multi-byte text wraps correctly.
func Wrap(text string, width int) string {
	if width <= 0 {
		return text
	}

	// Treat embedded newlines as ordinary word separators.
	text = strings.ReplaceAll(text, "\n", " ")

	// Short input needs no work, preserving spacing exactly.
	if len([]rune(text)) <= width {
		return text
	}

	var result strings.Builder
	words := strings.Fields(text)
	currentLine := ""

	for _, word := range words {
		// Tentatively extend the line before deciding to wrap.
		testLine := currentLine
		if testLine != "" {
			testLine += " "
		}
		testLine += word

		if len([]rune(testLine)) > width && currentLine != "" {
			// Current line is full, so flush it and start fresh.
			result.WriteString(currentLine)
			result.WriteString("\n")
			currentLine = word
		} else {
			// Word fits, so keep building the current line.
			currentLine = testLine
		}
	}

	// Flush the trailing partial line when words existed.
	if currentLine != "" {
		result.WriteString(currentLine)
	}

	return result.String()
}
