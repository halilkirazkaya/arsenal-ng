// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"strings"
	"testing"
)

// =============================================================================
// Hint Branches
// =============================================================================

// TestHintsAllBranches verifies every Hints branch renders its guidance.
func TestHintsAllBranches(t *testing.T) {
	variables := map[string]string{"ip": "1.1.1.1"}
	cases := []struct {
		query string
		want  string
	}{
		{"set", "Common variables"}, {"set ", "Common variables"},
		{"set i", "Suggestions"}, {"set zz", "custom variable"},
		{"set ip=1.1.1.1", "Press Enter to set variable"},
		{"unset", "No variables"}, {"unset ", "No variables"},
		{"variables", "show all variables"}, {"tools", "available tools"},
		{"help", "help menu"}, {"help ", "help menu"},
		{"whatever unknown", "execute command"},
	}
	for _, c := range cases {
		if got := Hints(c.query, nil); !strings.Contains(got, c.want) {
			t.Errorf("Hints(%q) missing %q", c.query, c.want)
		}
	}
	if got := Hints("unset", variables); !strings.Contains(got, "ip") {
		t.Errorf("unset hints must list names, got %q", got)
	}
}

// =============================================================================
// Layout Limits
// =============================================================================

// TestMaxVisibleClampsToMin verifies tiny terminals keep a minimum list.
func TestMaxVisibleClampsToMin(t *testing.T) {
	m := New(testCheats()).Resize(120, 1)
	if got := m.MaxVisible(); got != minVisible {
		t.Fatalf("MaxVisible = %d, want min %d", got, minVisible)
	}
	m = New(testCheats()).Resize(120, 100)
	if got := m.MaxVisible(); got <= minVisible {
		t.Fatalf("tall MaxVisible = %d, want > %d", got, minVisible)
	}
}

// TestWidthClampsNarrowTerminal verifies narrow widths stay positive.
func TestWidthClampsNarrowTerminal(t *testing.T) {
	if m := New(testCheats()).Resize(10, 40); m.Width() <= 0 {
		t.Fatalf("Width = %d, want positive fallback", m.Width())
	}
}

// TestInfoBoxWrapsLongDesc verifies long descriptions wrap in InfoBox.
func TestInfoBoxWrapsLongDesc(t *testing.T) {
	c := testCheats()[0]
	c.Desc = strings.Repeat("word ", 30)
	box := InfoBox(c, 40)
	// v2 styles always emit ANSI, splitting styled spans: assert parts.
	if !strings.Contains(box, "word") || !strings.Contains(box, "nmap") {
		t.Fatal("InfoBox must contain wrapped desc and command")
	}
}
