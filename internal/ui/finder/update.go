// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Update Outcomes
// =============================================================================

// Outcome describes what the parent must do after a search key press.
type Outcome int

const (
	// Stay keeps the search view active.
	Stay Outcome = iota
	// Picked selects the highlighted cheat (check Pick for details).
	Picked
	// ShowVariables opens the variables view.
	ShowVariables
	// ShowTools opens the tools view.
	ShowTools
	// ShowHelp opens the help view.
	ShowHelp
	// SetVar runs "set name=value" against the variable store.
	SetVar
	// UnsetVar runs "unset name" against the variable store.
	UnsetVar
	// Cancel quits the application without a command.
	Cancel
)

// Pick carries the selected cheat and whether the app should quit.
type Pick struct {
	// Selected cheat, nil when nothing was picked.
	Cheat *model.Cheat
	// Quit executes immediately without opening the args view.
	Quit bool
}

// =============================================================================
// Message Handling
// =============================================================================

// Update handles a message and reports the outcome to the parent.
func (m Model) Update(msg tea.Msg, cheats []*model.Cheat) (Model, tea.Cmd, Outcome, Pick) {
	var cmd tea.Cmd

	// A new keypress replaces any stale feedback.
	m.StatusMsg = ""
	m.StatusIsError = false

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case keys.CtrlC, keys.Esc:
			return m, nil, Cancel, Pick{}

		case keys.Help:
			if m.Input.Value() == "" {
				return m, nil, ShowHelp, Pick{}
			}

		case keys.Enter:
			return m.enter(cheats)

		case keys.Up, keys.CtrlP:
			m = m.MoveUp()

		case keys.Down, keys.CtrlN:
			m = m.MoveDown()

		case keys.PgUp:
			m = m.PageUp()

		case keys.PgDown:
			m = m.PageDown()

		default:
			m.Input, cmd = m.Input.Update(msg)
			m = m.Filter(cheats)
			return m, cmd, Stay, Pick{}
		}
	}

	m.Input, cmd = m.Input.Update(msg)
	return m, cmd, Stay, Pick{}
}
