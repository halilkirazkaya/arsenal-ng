// Package ui provides the terminal user interface for arsenal-ng.
//
// This file defines all UI component styles including text styles, list styles,
// input styles, container styles, header styles, and status styles used
// throughout the application.
package ui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// UI Component Styles
// =============================================================================

var (
	// Text styles
	titleStyle = lipgloss.NewStyle().
			Foreground(brightColor).
			Bold(true)

	descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(config.ColorDesc)).
			Italic(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(dimColor)

	// List styles
	toolStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Bold(true)

	tagStyle = lipgloss.NewStyle().
			Foreground(accentColor)

	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#4A4A4A")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true)

	cursorStyle = lipgloss.NewStyle().
		Foreground(primaryColor).
		Bold(true)

	// Input styles
	promptStyle = lipgloss.NewStyle().
		Foreground(accentColor).
		Bold(true)

	argNameStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Bold(true)

	argValueStyle = lipgloss.NewStyle().
			Foreground(brightColor)

	// Container styles
	infoBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(secondaryColor).
			Padding(0, 1)

	// Header styles
	logoStyle = lipgloss.NewStyle().
			Foreground(logoColor).
			Bold(true)

	versionStyle = lipgloss.NewStyle().
			Foreground(dimColor)

	// Status styles
	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#98D8C8")).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Bold(true)

	varCountStyle = lipgloss.NewStyle().
			Foreground(accentColor)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true)

	// Unused but kept for compatibility
	commandStyle = lipgloss.NewStyle().
			Foreground(dimColor)
)

