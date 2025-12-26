// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the main Update() router function that routes messages
// to appropriate state handlers based on application state. It also handles
// global window resize events.
package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update implements tea.Model. Routes messages to appropriate state handler.
func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle window resize globally
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = msg.Width
		m.height = msg.Height
		m.searchInput.Width = msg.Width - 10
		// Reinitialize tools table if in tools view
		if m.state == stateShowTools && m.toolsPerPage > 0 {
			m = m.initToolsTable()
		}
		// Reinitialize vars table if in vars view
		if m.state == stateShowVars {
			m = m.initVarsTable()
		}
	}

	switch m.state {
	case stateSearch:
		return m.updateSearch(msg)
	case stateArgs:
		return m.updateArgs(msg)
	case stateShowVars:
		return m.updateShowVars(msg)
	case stateShowTools:
		return m.updateShowTools(msg)
	case stateHelp:
		return m.updateHelp(msg)
	}
	return m, nil
}
