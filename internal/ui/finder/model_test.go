// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"io"
	"log"
	"os"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// TestMain silences logs and pins deterministic theme colors for tests.
func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	theme.SetLogoColor(lipgloss.Color("#FFFFFF"))
	os.Exit(m.Run())
}

// testCheats returns a small fixed cheat list for tests.
func testCheats() []*model.Cheat {
	return []*model.Cheat{
		{Tool: "nmap", Title: "Basic scan", Desc: "Scan a host", Command: "nmap -sV 10.0.0.1", Tags: []string{"scan"}},
		{Tool: "ffuf", Title: "Fuzz paths", Desc: "Fuzz web paths", Command: "ffuf -u {{url}}/FUZZ", Tags: []string{"web"}},
		{Tool: "nmap", Title: "Vuln scan", Desc: "Vulnerability scan", Command: "nmap --script vuln {{ip}}", Tags: []string{"scan"}},
	}
}

// rk builds a runes key message for typing simulation.
func rk(s string) tea.KeyPressMsg { return tea.KeyPressMsg{Text: string(s)} }

// =============================================================================
// Model Basics
// =============================================================================

// TestNewStartsUnselected verifies New shows all cheats with no selection.
func TestNewStartsUnselected(t *testing.T) {
	cheats := testCheats()
	m := New(cheats)
	if len(m.Filtered) != len(cheats) {
		t.Fatalf("Filtered = %d, want %d", len(m.Filtered), len(cheats))
	}
	if m.Cursor != -1 || m.Offset != 0 {
		t.Fatalf("Cursor/Offset = %d/%d, want -1/0", m.Cursor, m.Offset)
	}
	if !m.Input.Focused() {
		t.Fatal("input should be focused after New")
	}
}

// TestResizeFocusSetQueryResetList verifies basic model mutations.
func TestResizeFocusSetQueryResetList(t *testing.T) {
	m := New(testCheats()).Resize(100, 40)
	if m.Input.Width() != 90 {
		t.Fatalf("Input.Width = %d, want 90", m.Input.Width())
	}
	m.Input.Blur()
	m = m.Focus()
	if !m.Input.Focused() {
		t.Fatal("Focus should refocus the input")
	}
	m = m.SetQuery("nmap")
	if m.Input.Value() != "nmap" {
		t.Fatalf("Value = %q, want nmap", m.Input.Value())
	}
	m.Cursor = 2
	m.Offset = 1
	m = m.ResetList(testCheats())
	if m.Cursor != -1 || m.Offset != 0 || len(m.Filtered) != 3 {
		t.Fatalf("ResetList = cursor %d offset %d len %d", m.Cursor, m.Offset, len(m.Filtered))
	}
}
