// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"fmt"
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

// TestMain silences logs and pins the logo color for deterministic output.
func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	theme.SetLogoColor(lipgloss.Color("#FFFFFF"))
	os.Exit(m.Run())
}

// manyCheats builds n cheats with distinct zero-padded tool names.
func manyCheats(n int) []*model.Cheat {
	cheats := make([]*model.Cheat, 0, n)
	for i := range n {
		cheats = append(cheats, &model.Cheat{
			Tool:    fmt.Sprintf("tool-%03d", i),
			Title:   fmt.Sprintf("cmd %d", i),
			Command: "run it",
		})
	}
	return cheats
}

// keyMsg builds a key message for the named test key.
func keyMsg(s string) tea.KeyPressMsg {
	switch s {
	case "enter":
		return tea.KeyPressMsg{Code: tea.KeyEnter}
	case "esc":
		return tea.KeyPressMsg{Code: tea.KeyEsc}
	case "q":
		return tea.KeyPressMsg{Text: "q", Code: 'q'}
	case "h":
		return tea.KeyPressMsg{Text: "h", Code: 'h'}
	case "l":
		return tea.KeyPressMsg{Text: "l", Code: 'l'}
	case "left":
		return tea.KeyPressMsg{Code: tea.KeyLeft}
	case "right":
		return tea.KeyPressMsg{Code: tea.KeyRight}
	case "down":
		return tea.KeyPressMsg{Code: tea.KeyDown}
	}
	return tea.KeyPressMsg{Text: string(s)}
}

// =============================================================================
// New And Rebuild
// =============================================================================

// TestNewEmptyCheats verifies an empty model selects nothing with zero rows.
func TestNewEmptyCheats(t *testing.T) {
	m := New(nil, 120, 40)
	if m.Selected() != "" {
		t.Fatalf("Selected = %q, want empty", m.Selected())
	}
	if len(m.Table.Rows()) != 0 {
		t.Fatalf("rows = %d, want 0", len(m.Table.Rows()))
	}
}

// TestNewShowsFirstToolSelected verifies the first tool is pre-selected.
func TestNewShowsFirstToolSelected(t *testing.T) {
	m := New(manyCheats(3), 120, 40)
	if got := m.Selected(); got != "tool-000" {
		t.Fatalf("Selected = %q, want tool-000", got)
	}
}
