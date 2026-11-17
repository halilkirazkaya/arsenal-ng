// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"
)

// =============================================================================
// Special Commands
// =============================================================================

// TestIsSpecialCommand verifies special command detection across cases.
func TestIsSpecialCommand(t *testing.T) {
	cases := []struct {
		query string
		want  bool
	}{
		{"set", true}, {"SET", true}, {"  set  ", true}, {"set ip=1", true},
		{"unset", true}, {"unset ip", true}, {"variables", true},
		{"help", true}, {"tools", true},
		{"", false}, {"nmap", false}, {"settings", false}, {"help me", false},
		{"unsetx", false}, {"tool", false}, {"my set", false},
	}
	for _, c := range cases {
		if got := IsSpecialCommand(c.query); got != c.want {
			t.Errorf("IsSpecialCommand(%q) = %v, want %v", c.query, got, c.want)
		}
	}
}

// =============================================================================
// Query Filtering
// =============================================================================

// TestFilterNarrowsAndResetsCursor verifies Filter narrows and resets cursor.
func TestFilterNarrowsAndResetsCursor(t *testing.T) {
	m := New(testCheats())
	m.Cursor = 2
	m.Offset = 1
	m = m.SetQuery("ffuf")
	m = m.Filter(testCheats())
	if len(m.Filtered) != 1 || m.Filtered[0].Tool != "ffuf" {
		t.Fatalf("Filtered = %+v, want only ffuf", m.Filtered)
	}
	if m.Cursor != -1 || m.Offset != 0 {
		t.Fatalf("cursor/offset not reset: %d/%d", m.Cursor, m.Offset)
	}
}

// TestFilterSkipsSpecialCommands verifies special queries keep the list.
func TestFilterSkipsSpecialCommands(t *testing.T) {
	m := New(testCheats()).SetQuery("set ip=1")
	m = m.Filter(testCheats())
	if len(m.Filtered) != 3 {
		t.Fatalf("special query must keep full list, got %d", len(m.Filtered))
	}
}
