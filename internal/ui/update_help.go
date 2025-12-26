// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the help view update handlers. It processes keyboard
// input to exit the help screen and return to the search view.
package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// =============================================================================
// Help View Update
// =============================================================================

// updateHelp handles input in the help view.
func (m App) updateHelp(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case keyEsc, keyQuit, keyEnter, keyHelp:
			m.state = stateSearch
			m.searchInput.Focus()
			return m, nil
		}
	}
	return m, nil
}

