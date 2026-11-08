// Package theme provides the shared lipgloss styles for the TUI: base
// palette, component styles, command syntax styles, and tag rendering.
package theme

import (
	"charm.land/lipgloss/v2"
)

// =============================================================================
// Text Styles
// =============================================================================

// Title renders headings in bold bright text.
var Title = lipgloss.NewStyle().
	Foreground(Bright).
	Bold(true)

// Desc renders descriptions in italic light gray.
var Desc = lipgloss.NewStyle().
	Foreground(lipgloss.Color(ColorDesc)).
	Italic(true)

// Help renders help hints in dim gray.
var Help = lipgloss.NewStyle().
	Foreground(Dim)

// =============================================================================
// List Styles
// =============================================================================

// Tool renders tool names in bold cyan.
var Tool = lipgloss.NewStyle().
	Foreground(Secondary).
	Bold(true)

// Selected renders the highlighted row with a gray background.
var Selected = lipgloss.NewStyle().
	Background(lipgloss.Color("#4A4A4A")).
	Foreground(lipgloss.Color("#FFFFFF")).
	Bold(true)

// Cursor renders the row marker in bold red.
var Cursor = lipgloss.NewStyle().
	Foreground(Primary).
	Bold(true)

// =============================================================================
// Input Styles
// =============================================================================

// Prompt renders input prompts in bold yellow.
var Prompt = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(true)

// ArgName renders argument names in bold cyan.
var ArgName = lipgloss.NewStyle().
	Foreground(Secondary).
	Bold(true)
