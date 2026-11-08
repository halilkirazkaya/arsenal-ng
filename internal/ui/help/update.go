// Package help provides the help view: keyboard shortcuts, usage
// examples, and workflow instructions.
package help

import (
	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
)

// =============================================================================
// Update Outcome
// =============================================================================

// Outcome describes what the parent must do after a help key press.
type Outcome int

const (
	// Stay keeps the help view active.
	Stay Outcome = iota
	// Back returns to the search view.
	Back
)

// =============================================================================
// Message Handling
// =============================================================================

// Update handles a message and reports the outcome to the parent.
func Update(msg tea.Msg) Outcome {
	if msg, ok := msg.(tea.KeyPressMsg); ok {
		switch msg.String() {
		case keys.Esc, keys.Quit, keys.Enter, keys.Help:
			return Back
		}
	}
	return Stay
}
