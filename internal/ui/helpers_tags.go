// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains tag rendering functions that assign consistent colors
// to tags and tools based on their hash values. It provides visual
// distinction for different tags in the UI.
package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Tag Rendering
// =============================================================================

// renderColoredTags renders each tag with a different color based on its hash.
func renderColoredTags(tags []string) string {
	if len(tags) == 0 {
		return ""
	}

	parts := make([]string, len(tags))
	for i, tag := range tags {
		color := getTagColor(tag)
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))
		parts[i] = style.Render(tag)
	}

	sep := lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorDim)).Render(", ")
	return strings.Join(parts, sep)
}

// getTagColor returns a consistent color for a tag based on its hash.
func getTagColor(tag string) string {
	hash := 0
	for _, c := range tag {
		hash = hash*31 + int(c)
	}
	if hash < 0 {
		hash = -hash
	}
	return config.TagColors[hash%len(config.TagColors)]
}

// getToolColor returns a consistent color for a tool based on its hash.
func getToolColor(tool string) string {
	hash := 0
	for _, c := range tool {
		hash = hash*31 + int(c)
	}
	if hash < 0 {
		hash = -hash
	}
	return config.TagColors[hash%len(config.TagColors)]
}

