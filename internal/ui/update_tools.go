// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the tools view update handlers. It processes keyboard
// input for navigating the tools table, handles pagination (previous/next page),
// and manages exit back to the search view.
package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// =============================================================================
// Tools View Update
// =============================================================================

// updateShowTools handles input in the show tools view.
func (m App) updateShowTools(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keyEsc, keyQuit:
			m.state = stateSearch
			m.searchInput.Focus()
			return m, nil

		case keyEnter:
			// If a tool is selected, search for it; otherwise go back
			selectedRow := m.toolsTable.SelectedRow()
			if len(selectedRow) > 0 {
				// Get tool name from first column
				toolName := selectedRow[0]
				// Switch to search view and set the tool name as search query
				m.state = stateSearch
				m.searchInput.SetValue(toolName)
				m.searchInput.Focus()
				// Update filter to show results for this tool
				m = m.updateFilter()
			} else {
				// No selection, just go back
				m.state = stateSearch
				m.searchInput.Focus()
			}
			return m, nil

		case keyLeft, "h":
			// Previous page
			if m.toolsPaginator.Page > 0 {
				m.toolsPaginator.PrevPage()
				m = m.initToolsTable()
			}
			return m, nil

		case keyRight, "l":
			// Next page
			if m.toolsPaginator.Page < m.toolsPaginator.TotalPages-1 {
				m.toolsPaginator.NextPage()
				m = m.initToolsTable()
			}
			return m, nil

		case keyUp, keyCtrlP:
			// Move up in table
			m.toolsTable, cmd = m.toolsTable.Update(msg)
			return m, cmd

		case keyDown, keyCtrlN:
			// Move down in table
			m.toolsTable, cmd = m.toolsTable.Update(msg)
			return m, cmd

		default:
			// Pass other keys to table
			m.toolsTable, cmd = m.toolsTable.Update(msg)
			return m, cmd
		}
	}

	// Pass other messages to table
	m.toolsTable, cmd = m.toolsTable.Update(msg)
	return m, cmd
}

