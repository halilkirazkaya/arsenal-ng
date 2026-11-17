// Package help provides the help view: keyboard shortcuts, usage
// examples, and workflow instructions.
package help

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
	case "?":
		return tea.KeyPressMsg{Text: "?", Code: '?'}
	case "q":
		return tea.KeyPressMsg{Text: "q", Code: 'q'}
	}
	return tea.KeyPressMsg{Text: string(s)}
}

// =============================================================================
// Update Outcomes
// =============================================================================

// TestUpdateBackKeys verifies leave keys return Back.
func TestUpdateBackKeys(t *testing.T) {
	for _, k := range []string{"esc", "q", "enter", "?"} {
		if got := Update(keyMsg(k)); got != Back {
			t.Errorf("%s = %v, want Back", k, got)
		}
	}
}

// TestUpdateStaysOnOtherKeys verifies navigation keys keep the help open.
func TestUpdateStaysOnOtherKeys(t *testing.T) {
	for _, k := range []string{"down", "up", "x", "tab"} {
		var msg tea.Msg = keyMsg(k)
		if k == "down" {
			msg = tea.KeyPressMsg{Code: tea.KeyDown}
		}
		if k == "up" {
			msg = tea.KeyPressMsg{Code: tea.KeyUp}
		}
		if k == "tab" {
			msg = tea.KeyPressMsg{Code: tea.KeyTab}
		}
		if got := Update(msg); got != Stay {
			t.Errorf("%s = %v, want Stay", k, got)
		}
	}
}

// TestUpdateIgnoresNonKeyMsg verifies non-key messages keep the help open.
func TestUpdateIgnoresNonKeyMsg(t *testing.T) {
	if got := Update(nil); got != Stay {
		t.Errorf("nil msg = %v, want Stay", got)
	}
	if got := Update(struct{}{}); got != Stay {
		t.Errorf("struct msg = %v, want Stay", got)
	}
}
