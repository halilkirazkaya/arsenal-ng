// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
)

// =============================================================================
// Update Outcome
// =============================================================================

// Outcome describes what the parent must do after a variables key press.
type Outcome int

const (
	// Stay keeps the variables view active.
	Stay Outcome = iota
	// Back returns to the search view.
	Back
)

// =============================================================================
// Message Handling
// =============================================================================

// Update handles a message and reports the outcome to the parent.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd, Outcome) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case keys.Esc, keys.Quit, keys.Enter:
			return m, nil, Back
		default:
			// Pass all keys to table for navigation.
			m.table, cmd = m.table.Update(msg)
			return m, cmd, Stay
		}
	default:
		// Pass all other messages to table.
		m.table, cmd = m.table.Update(msg)
		return m, cmd, Stay
	}
}
