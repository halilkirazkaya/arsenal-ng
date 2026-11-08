// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"log"

	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Update Outcomes
// =============================================================================

// Outcome describes what the parent must do after an args key press.
type Outcome int

const (
	// Stay keeps the args view active.
	Stay Outcome = iota
	// Submit confirms the form (check Command; empty means blocked).
	Submit
	// Back returns to the search view.
	Back
	// Cancel quits the application without a command.
	Cancel
)

// =============================================================================
// Message Handling
// =============================================================================

// Update handles a message and reports the outcome to the parent.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd, Outcome, string) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case keys.CtrlC:
			return m, nil, Cancel, ""

		case keys.Esc:
			return m, nil, Back, ""

		case keys.Enter:
			return m.submit()

		case keys.Tab, keys.Down:
			return m.Next(), nil, Stay, ""

		case keys.ShiftTab, keys.Up:
			return m.Prev(), nil, Stay, ""
		}
	}

	m.Inputs[m.Cursor], cmd = m.Inputs[m.Cursor].Update(msg)
	return m, cmd, Stay, ""
}

// =============================================================================
// Form Submission
// =============================================================================

// submit reads the inputs and builds the final command.
// An empty command means a required argument is still missing.
func (m Model) submit() (Model, tea.Cmd, Outcome, string) {
	for i := range m.Args {
		m.Args[i].Value = m.Inputs[i].Value()
	}

	// Missing required args block submission with an empty command.
	if !model.HasEmptyArgs(m.Args) {
		command := model.BuildCommand(m.Cheat.Command, m.Args)
		// Never log the command: arguments often hold credentials.
		log.Printf("Command built from %d args (%d bytes)", len(m.Args), len(command))
		return m, nil, Submit, command
	}
	log.Printf("WARNING: Attempted to submit args with empty values")
	return m, nil, Submit, ""
}
