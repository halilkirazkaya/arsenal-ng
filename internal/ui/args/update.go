// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"log"
	"slices"

	"charm.land/bubbles/v2/textinput"
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

// staleSuggestion reports whether the selected completion no longer exists
// on disk. Bubbles caches matches from an earlier keystroke while the
// filesystem may change underneath, so Tab revalidates the selection
// through the same pathSuggestions mechanism before accepting it.
func staleSuggestion(in textinput.Model) bool {
	sel := in.CurrentSuggestion()
	if sel == "" || len(in.MatchedSuggestions()) == 0 {
		return false
	}
	return !slices.Contains(pathSuggestions(in.Value()), sel)
}

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

		case keys.Tab:
			// Tab completes a path when there is a suggestion left to
			// append; otherwise it moves on like ↓. The empty-match
			// guard matters: after Ctrl+P with no matches bubbles can
			// leave the suggestion index at -1, where
			// CurrentSuggestion would index out of bounds.
			in := m.Inputs[m.Cursor]
			if len(in.MatchedSuggestions()) == 0 || len(in.CurrentSuggestion()) <= len(in.Value()) {
				return m.Next(), nil, Stay, ""
			}
			// The suggestion list is a cached filesystem snapshot, so a
			// selected entry may have vanished since it was listed. Tab
			// must complete to something that still exists: refresh and
			// fall through only if the selection survives, otherwise
			// advance when nothing matches anymore.
			if staleSuggestion(in) {
				m.suggest()
				in = m.Inputs[m.Cursor]
				if len(in.MatchedSuggestions()) == 0 {
					return m.Next(), nil, Stay, ""
				}
			}

		case keys.Down:
			return m.Next(), nil, Stay, ""

		case keys.ShiftTab, keys.Up:
			return m.Prev(), nil, Stay, ""
		}
	}

	before := m.Inputs[m.Cursor].Value()
	m.Inputs[m.Cursor], cmd = m.Inputs[m.Cursor].Update(msg)
	if m.Inputs[m.Cursor].Value() != before {
		m.suggest()
	}
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
