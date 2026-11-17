// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	"io"
	"log"
	"os"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// TestMain silences logs and pins the logo color for deterministic output.
func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	theme.SetLogoColor(lipgloss.Color("#FFFFFF"))
	os.Exit(m.Run())
}

// keyMsg builds a key message for the named test key.
func keyMsg(s string) tea.KeyPressMsg {
	switch s {
	case "enter":
		return tea.KeyPressMsg{Code: tea.KeyEnter}
	case "esc":
		return tea.KeyPressMsg{Code: tea.KeyEsc}
	case "down":
		return tea.KeyPressMsg{Code: tea.KeyDown}
	case "q":
		return tea.KeyPressMsg{Text: "q", Code: 'q'}
	}
	return tea.KeyPressMsg{Text: string(s)}
}

// =============================================================================
// New And Rebuild
// =============================================================================

// TestNewSortsKeysAlphabetically verifies table rows follow sorted key order.
func TestNewSortsKeysAlphabetically(t *testing.T) {
	variables := map[string]string{"zebra": "z", "apple": "a", "mango": "m"}
	m := New(variables, 40)
	rows := m.table.Rows()
	want := []string{"apple", "mango", "zebra"}
	if len(rows) != len(want) {
		t.Fatalf("rows = %d, want %d", len(rows), len(want))
	}
	for i, w := range want {
		if rows[i][0] != w {
			t.Errorf("row %d = %q, want %q", i, rows[i][0], w)
		}
	}
}

// TestRebuildRefreshesRows verifies rebuild swaps the row set completely.
func TestRebuildRefreshesRows(t *testing.T) {
	m := New(map[string]string{"a": "1"}, 40)
	m = m.Rebuild(map[string]string{"b": "2", "a": "1"}, 40)
	if len(m.table.Rows()) != 2 {
		t.Fatalf("rows = %d, want 2", len(m.table.Rows()))
	}
	m = m.Rebuild(nil, 40)
	if !m.Empty() {
		t.Fatal("Rebuild(nil) must empty the table")
	}
}

// TestEmptyReportsState verifies Empty on nil and populated models.
func TestEmptyReportsState(t *testing.T) {
	if !New(nil, 40).Empty() {
		t.Fatal("nil variables must be empty")
	}
	if New(map[string]string{"a": "1"}, 40).Empty() {
		t.Fatal("one var must not be empty")
	}
}
