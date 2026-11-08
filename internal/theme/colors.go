// Package theme provides the shared lipgloss styles for the TUI: base
// palette, component styles, command syntax styles, and tag rendering.
package theme

import (
	"image/color"
	"math/rand/v2"

	"charm.land/lipgloss/v2"
)

// =============================================================================
// Base Palette
// =============================================================================

// Primary marks titles, cursors, and prompts in red.
var Primary = lipgloss.Color(ColorPrimary)

// Secondary marks tools and arguments in cyan.
var Secondary = lipgloss.Color(ColorSecondary)

// Accent marks tags, highlights, and matches in yellow.
var Accent = lipgloss.Color(ColorAccent)

// Dim marks help text and separators in gray.
var Dim = lipgloss.Color(ColorDim)

// Bright marks selected and important text in white.
var Bright = lipgloss.Color(ColorBright)

// =============================================================================
// Logo Color
// =============================================================================

// logoColor holds the random logo color chosen once at startup.
var logoColor color.Color

// SetLogoColor overrides the logo color for deterministic output in tests.
func SetLogoColor(c color.Color) {
	logoColor = c
	Logo = lipgloss.NewStyle().Foreground(logoColor).Bold(true)
}

// =============================================================================
// Startup Setup
// =============================================================================

// init picks a random tag color for the logo on first load.
func init() {
	logoColor = lipgloss.Color(TagColors[rand.IntN(len(TagColors))])
	Logo = lipgloss.NewStyle().Foreground(logoColor).Bold(true)
}
