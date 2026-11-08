// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

// =============================================================================
// Enter Fixtures
// =============================================================================

// keyMsg builds a key message for the given key name.
func keyMsg(s string) tea.KeyPressMsg {
	switch s {
	case "enter":
		return tea.KeyPressMsg{Code: tea.KeyEnter}
	case "esc":
		return tea.KeyPressMsg{Code: tea.KeyEsc}
	case "up":
		return tea.KeyPressMsg{Code: tea.KeyUp}
	case "down":
		return tea.KeyPressMsg{Code: tea.KeyDown}
	case "?":
		return tea.KeyPressMsg{Text: "?", Code: '?'}
	}
	return tea.KeyPressMsg{Text: string(s)}
}

// enterWith runs enter with the given query and cursor position.
func enterWith(t *testing.T, query string, cursor int) (Outcome, Pick) {
	t.Helper()
	m := New(testCheats()).SetQuery(query)
	m.Cursor = cursor
	_, _, oc, pick := m.enter(testCheats())
	return oc, pick
}

// =============================================================================
// Enter Outcomes
// =============================================================================

// TestEnterSpecialCommands verifies special commands map to their outcomes.
func TestEnterSpecialCommands(t *testing.T) {
	cases := []struct {
		query string
		want  Outcome
	}{
		{"variables", ShowVariables}, {"  VARIABLES ", ShowVariables},
		{"help", ShowHelp}, {"tools", ShowTools},
		{"set ip=1", SetVar}, {"unset ip", UnsetVar},
	}
	for _, c := range cases {
		oc, _ := enterWith(t, c.query, 0)
		if oc != c.want {
			t.Errorf("enter(%q) = %v, want %v", c.query, oc, c.want)
		}
	}
}

// TestEnterBareSetUnsetWithoutSelection verifies bare commands stay put.
func TestEnterBareSetUnsetWithoutSelection(t *testing.T) {
	// Bare "set"/"unset" match no enter branch, so an empty selection stays.
	for _, q := range []string{"set", "unset", "  set  "} {
		oc, pick := enterWith(t, q, -1)
		if oc != Stay || pick.Cheat != nil {
			t.Errorf("enter(%q) = %v/%v, want Stay/nil", q, oc, pick)
		}
	}
}

// TestEnterBareSetUnsetWithSelection verifies bare commands pick cheats.
func TestEnterBareSetUnsetWithSelection(t *testing.T) {
	// Bare "set"/"unset" match no branch, so a selection picks the cheat.
	for _, q := range []string{"set", "unset"} {
		oc, pick := enterWith(t, q, 0)
		if oc != Picked || pick.Cheat == nil {
			t.Errorf("enter(%q) = %v/%v, want Picked/cheat", q, oc, pick)
		}
	}
}
