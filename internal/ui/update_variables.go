// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the variables view update handlers. It processes
// keyboard input for navigating the variables table and handles exit
// back to the search view.
package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// =============================================================================
// Variables View Update
// =============================================================================

// updateShowVars handles input in the show variables view.
func (m App) updateShowVars(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Initialize table if needed (check if table is empty or uninitialized)
	if len(m.varsTable.Rows()) == 0 {
		m = m.initVarsTable()
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keyEsc, keyQuit, keyEnter:
			m.state = stateSearch
			m.searchInput.Focus()
			return m, nil
		default:
			// Pass all keys to table for navigation
			m.varsTable, cmd = m.varsTable.Update(msg)
			return m, cmd
		}
	default:
		// Pass all other messages to table
		m.varsTable, cmd = m.varsTable.Update(msg)
		return m, cmd
	}
}

