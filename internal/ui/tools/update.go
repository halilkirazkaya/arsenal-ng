// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Update Outcome
// =============================================================================

// Outcome describes what the parent must do after a tools key press.
type Outcome int

const (
	// Stay keeps the tools view active.
	Stay Outcome = iota
	// Pick searches for the selected tool (check Tool; empty means back).
	Pick
	// Back returns to the search view.
	Back
)

// =============================================================================
// Message Handling
// =============================================================================

// Update handles a message and reports the outcome to the parent.
func (m Model) Update(msg tea.Msg, cheats []*model.Cheat, width, height int) (Model, tea.Cmd, Outcome, string) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case keys.Esc, keys.Quit:
			return m, nil, Back, ""

		case keys.Enter:
			// If a tool is selected, search for it; otherwise go back.
			return m, nil, Pick, m.Selected()

		case keys.Left, keys.VimLeft:
			return m.PrevPage(cheats, width, height), nil, Stay, ""

		case keys.Right, keys.VimRight:
			return m.NextPage(cheats, width, height), nil, Stay, ""

		default:
			// Pass other keys to table.
			m.Table, cmd = m.Table.Update(msg)
			return m, cmd, Stay, ""
		}
	}

	// Pass other messages to table.
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd, Stay, ""
}
