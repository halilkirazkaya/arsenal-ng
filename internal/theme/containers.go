// Package theme provides the shared lipgloss styles for the TUI: base
// palette, component styles, command syntax styles, and tag rendering.
package theme

import (
	"charm.land/lipgloss/v2"
)

// =============================================================================
// Container Styles
// =============================================================================

// InfoBox renders bordered panels with a cyan rounded frame.
var InfoBox = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(Secondary).
	Padding(0, 1)

// =============================================================================
// Header Styles
// =============================================================================

// Logo renders the ASCII banner in the startup logo color.
var Logo = lipgloss.NewStyle().
	Foreground(logoColor).
	Bold(true)

// Version renders the version string in dim gray.
var Version = lipgloss.NewStyle().
	Foreground(Dim)

// =============================================================================
// Status Styles
// =============================================================================

// Success renders confirmations in bold light green.
var Success = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#98D8C8")).
	Bold(true)

// Error renders failures in bold red.
var Error = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF6B6B")).
	Bold(true)

// VarCount renders the active variable badge in yellow.
var VarCount = lipgloss.NewStyle().
	Foreground(Accent)

// StatusBar renders footer status text in bold white.
var StatusBar = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Bold(true)
