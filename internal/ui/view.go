// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the main View() router function that routes to appropriate
// view renderers based on application state, and common render functions like
// renderHeader() and renderFooter() used across multiple views.
package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// View implements tea.Model. Routes to appropriate view based on state.
func (m App) View() string {
	switch m.state {
	case stateSearch:
		return m.viewSearch()
	case stateArgs:
		return m.viewArgs()
	case stateShowVars:
		return m.viewShowVars()
	case stateShowTools:
		return m.viewShowTools()
	case stateHelp:
		return m.viewHelp()
	}
	return ""
}

// =============================================================================
// Common Render Functions
// =============================================================================

// renderHeader renders the application header with logo and version.
func (m App) renderHeader() string {
	logo := logoStyle.Render(config.Logo)
	version := versionStyle.Render(" " + config.GetVersionInfo())
	return lipgloss.JoinHorizontal(lipgloss.Bottom, logo, version)
}

// renderFooter renders the footer with status information and help text.
func (m App) renderFooter() string {
	// Status count
	cursorPos := m.cursor + 1
	if m.cursor < 0 {
		cursorPos = 0
	}
	status := statusBarStyle.Render(fmt.Sprintf(" %d/%d ", cursorPos, len(m.filtered)))

	// Variable count (if any)
	varInfo := ""
	if count := m.globals.Count(); count > 0 {
		varInfo = varCountStyle.Render(fmt.Sprintf("│ active variable count: %d ", count))
	}

	// Help text
	help := statusBarStyle.Render("│ ↑/↓: nav │ set key=value │ unset key │ variables │ tools │ help │ esc: quit")

	return status + varInfo + help
}
