// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

// =============================================================================
// Non Key Messages
// =============================================================================

// TestUpdateNonKeyPassesToTable verifies window resizes stay in the view.
func TestUpdateNonKeyPassesToTable(t *testing.T) {
	m := New(map[string]string{"a": "1"}, 30)
	_, cmd, oc := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	if oc != Stay {
		t.Fatalf("outcome = %v, want Stay", oc)
	}
	_ = cmd
}
