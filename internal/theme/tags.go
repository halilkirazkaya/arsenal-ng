// Package theme provides the shared lipgloss styles for the TUI: base
// palette, component styles, command syntax styles, and tag rendering.
package theme

import (
	"strings"

	"charm.land/lipgloss/v2"
)

// =============================================================================
// Tag Rendering
// =============================================================================

// ColoredTags renders tags in distinct hash-based colors joined by commas.
func ColoredTags(tags []string) string {
	if len(tags) == 0 {
		return ""
	}

	parts := make([]string, len(tags))
	for i, tag := range tags {
		color := tagColor(tag)
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		parts[i] = style.Render(tag)
	}

	sep := lipgloss.NewStyle().Foreground(lipgloss.Color(ColorDim)).Render(", ")
	return strings.Join(parts, sep)
}

// =============================================================================
// Tag Hashing
// =============================================================================

// tagColor maps a tag to a stable palette entry via its hash.
func tagColor(tag string) string {
	hash := 0
	for _, c := range tag {
		hash = hash*31 + int(c)
	}
	// Negate overflowed hashes so the modulo below stays valid.
	if hash < 0 {
		hash = -hash
	}
	return TagColors[hash%len(TagColors)]
}
