// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the search view update handlers including input processing,
// navigation helpers, filter updates, special command handling (set/unset),
// and state transitions to argument input mode.
package ui

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/loader"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Search View Update
// =============================================================================

// updateSearch handles input in the main search view.
func (m App) updateSearch(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Clear status message on any key
	m.statusMsg = ""
	m.statusIsError = false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case keyCtrlC, keyEsc:
			log.Printf("User cancelled application (key: %s)", msg.String())
			m.Cancelled = true
			return m, tea.Quit

		case keyHelp:
			if m.searchInput.Value() == "" {
				m.state = stateHelp
				return m, nil
			}
			m.searchInput, cmd = m.searchInput.Update(msg)
			return m, cmd

		case keyEnter:
			return m.handleEnter()

		case keyUp, keyCtrlP:
			m = m.moveCursorUp()

		case keyDown, keyCtrlN:
			m = m.moveCursorDown()

		case keyPgUp:
			m = m.pageUp()

		case keyPgDown:
			m = m.pageDown()

		default:
			m.searchInput, cmd = m.searchInput.Update(msg)
			m = m.updateFilter()
			return m, cmd
		}
	}

	m.searchInput, cmd = m.searchInput.Update(msg)
	return m, cmd
}

// handleEnter processes enter key in search view.
func (m App) handleEnter() (tea.Model, tea.Cmd) {
	query := strings.TrimSpace(m.searchInput.Value())

	// Handle special commands (set/unset/show)
	if handled, newModel := m.handleSpecialCommand(query); handled {
		return newModel, nil
	}

	// Normal cheat selection
	if len(m.filtered) == 0 || m.cursor < 0 || m.cursor >= len(m.filtered) {
		log.Printf("Invalid cheat selection: cursor=%d, filtered=%d", m.cursor, len(m.filtered))
		return m, nil
	}

	m.selectedCheat = m.filtered[m.cursor]
	log.Printf("Selected cheat: %s (tool: %s, file: %s)", m.selectedCheat.Title, m.selectedCheat.Tool, m.selectedCheat.Filename)
	m.args = model.ParseArguments(m.selectedCheat.Command)

	// Pre-fill arguments with global variables
	appliedVars := []string{}
	for i, arg := range m.args {
		if globalVal, ok := m.globals.Get(arg.Name); ok {
			m.args[i].Value = globalVal
			appliedVars = append(appliedVars, arg.Name)
		} else if arg.DefaultValue != "" {
			m.args[i].Value = arg.DefaultValue
		}
	}

	if len(appliedVars) > 0 {
		log.Printf("Pre-filled %d argument(s) with variables: %v", len(appliedVars), appliedVars)
	}

	// Execute directly only if there are no arguments at all
	// If arguments exist, always show the input form (even if pre-filled with variables)
	if len(m.args) == 0 {
		m.FinalCommand = model.BuildCommand(m.selectedCheat.Command, m.args)
		log.Printf("Command ready (no arguments): %s", m.FinalCommand)
		return m, tea.Quit
	}

	// Show argument input form (variables will be pre-filled but user can review/edit)
	log.Printf("Entering argument input mode (%d argument(s), %d pre-filled with variables)", len(m.args), len(appliedVars))
	return m.enterArgsState()
}

// enterArgsState transitions to argument input mode.
func (m App) enterArgsState() (tea.Model, tea.Cmd) {
	m.state = stateArgs
	m.argInputs = make([]textinput.Model, len(m.args))
	firstEmpty := -1

	for i, arg := range m.args {
		ti := textinput.New()
		ti.Placeholder = arg.Name
		ti.CharLimit = config.ArgCharLimit
		ti.Width = 40
		ti.SetValue(arg.Value)

		if arg.Value == "" && firstEmpty == -1 {
			firstEmpty = i
			ti.Focus()
		}
		m.argInputs[i] = ti
	}

	if firstEmpty == -1 {
		m.argInputs[0].Focus()
		m.argCursor = 0
	} else {
		m.argCursor = firstEmpty
	}

	return m, nil
}

// =============================================================================
// Navigation Helpers
// =============================================================================

// moveCursorUp moves the cursor up in the list.
func (m App) moveCursorUp() App {
	if m.cursor > 0 {
		m.cursor--
		if m.cursor < m.offset {
			m.offset = m.cursor
		}
	} else if m.cursor == 0 {
		// Move from first item to no selection
		m.cursor = -1
	}
	return m
}

// moveCursorDown moves the cursor down in the list.
func (m App) moveCursorDown() App {
	if m.cursor < 0 {
		// Move from no selection to first item
		if len(m.filtered) > 0 {
			m.cursor = 0
			m.offset = 0
		}
	} else if m.cursor < len(m.filtered)-1 {
		m.cursor++
		maxVisible := m.maxVisibleItems()
		if m.cursor >= m.offset+maxVisible {
			m.offset = m.cursor - maxVisible + 1
		}
	}
	return m
}

// pageUp moves the cursor one page up.
func (m App) pageUp() App {
	if m.cursor < 0 {
		return m // Already at top (no selection)
	}
	m.cursor -= m.maxVisibleItems()
	if m.cursor < 0 {
		m.cursor = -1 // Move to no selection
		m.offset = 0
	} else {
		m.offset = m.cursor
	}
	return m
}

// pageDown moves the cursor one page down.
func (m App) pageDown() App {
	if m.cursor < 0 {
		// Move from no selection to first page
		if len(m.filtered) > 0 {
			m.cursor = 0
			m.offset = 0
		}
		return m
	}
	m.cursor += m.maxVisibleItems()
	if m.cursor >= len(m.filtered) {
		m.cursor = len(m.filtered) - 1
	}
	maxVisible := m.maxVisibleItems()
	if m.cursor >= m.offset+maxVisible {
		m.offset = m.cursor - maxVisible + 1
	}
	return m
}

// =============================================================================
// Filter and Special Commands
// =============================================================================

// updateFilter updates the filtered list based on search query.
func (m App) updateFilter() App {
	query := m.searchInput.Value()
	if !isSpecialCommand(query) {
		m.filtered = loader.Search(m.cheats, query)
		// Reset cursor to no selection when filtering
		m.cursor = -1
		m.offset = 0
	}
	return m
}

// isSpecialCommand checks if the query is a special command.
func isSpecialCommand(query string) bool {
	q := strings.TrimSpace(strings.ToLower(query))
	return q == cmdSet ||
		strings.HasPrefix(q, cmdSet+" ") ||
		q == cmdUnset ||
		strings.HasPrefix(q, cmdUnset+" ") ||
		q == "variables" ||
		q == "help" ||
		q == "tools"
}

// handleSpecialCommand processes set/unset/show commands.
func (m App) handleSpecialCommand(query string) (bool, App) {
	q := strings.TrimSpace(query)
	qLower := strings.ToLower(q)

	switch {
	case qLower == "variables":
		log.Printf("User navigated to variables view")
		m.state = stateShowVars
		m.searchInput.SetValue("")
		// Initialize vars table when entering variables view
		m = m.initVarsTable()
		return true, m

	case qLower == "help":
		log.Printf("User navigated to help view")
		m.state = stateHelp
		m.searchInput.SetValue("")
		return true, m

	case qLower == "tools":
		log.Printf("User navigated to tools view")
		m.state = stateShowTools
		m.searchInput.SetValue("")
		// Initialize tools table
		if m.toolsPerPage == 0 {
			m.toolsPerPage = 25
		}
		m = m.initToolsTable()
		return true, m

	case strings.HasPrefix(qLower, "set "):
		return m.handleSetCommand(q)

	case strings.HasPrefix(qLower, "unset "):
		return m.handleUnsetCommand(q)
	}

	return false, m
}

// handleSetCommand processes the "set" command.
func (m App) handleSetCommand(query string) (bool, App) {
	parts := strings.SplitN(query[4:], "=", 2)
	if len(parts) == 2 {
		name := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if name != "" && value != "" {
			if err := m.globals.Set(name, value); err != nil {
				log.Printf("ERROR: Failed to set variable %s: %v", name, err)
				m.statusMsg = fmt.Sprintf("✗ Failed to save variable: %v", err)
				m.statusIsError = true
			} else {
				m.statusMsg = fmt.Sprintf("✓ Set %s = %s", name, value)
				// Reinitialize vars table if in variables view
				if m.state == stateShowVars {
					m = m.initVarsTable()
				}
			}
			m.searchInput.SetValue("")
			m.filtered = m.cheats
			return true, m
		}
	}
	log.Printf("Invalid set command format: %s", query)
	m.statusMsg = "✗ Usage: set name=value"
	m.statusIsError = true
	return true, m
}

// handleUnsetCommand processes the "unset" command.
func (m App) handleUnsetCommand(query string) (bool, App) {
	name := strings.TrimSpace(query[6:])
	if name != "" {
		existed, err := m.globals.Unset(name)
		if err != nil {
			log.Printf("ERROR: Failed to unset variable %s: %v", name, err)
			m.statusMsg = fmt.Sprintf("✗ Failed to save changes: %v", err)
			m.statusIsError = true
		} else if existed {
			m.statusMsg = fmt.Sprintf("✓ Unset %s", name)
			// Reinitialize vars table if in variables view
			if m.state == stateShowVars {
				m = m.initVarsTable()
			}
		} else {
			log.Printf("Attempted to unset non-existent variable: %s", name)
			m.statusMsg = fmt.Sprintf("✗ Variable '%s' not found", name)
			m.statusIsError = true
		}
		m.searchInput.SetValue("")
		m.filtered = m.cheats
		return true, m
	}
	log.Printf("Invalid unset command format: %s", query)
	m.statusMsg = "✗ Usage: unset name"
	m.statusIsError = true
	return true, m
}

