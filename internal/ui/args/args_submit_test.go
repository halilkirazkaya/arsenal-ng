// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"os"
	"path/filepath"
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Form Submission
// =============================================================================

// TestSubmitBuildsCommand verifies submit assembles the final command.
func TestSubmitBuildsCommand(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	m.Inputs[0].SetValue("10.0.0.1")
	m.Inputs[1].SetValue("443")
	_, _, oc, cmd := m.submit()
	if oc != Submit || cmd != "nmap -sV 10.0.0.1 443" {
		t.Fatalf("submit = %v/%q", oc, cmd)
	}
}

// TestSubmitBlockedWhenRequiredEmpty verifies empty required args block.
func TestSubmitBlockedWhenRequiredEmpty(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	// Required ip stays empty so submission must return no command.
	m.Inputs[0].SetValue("")
	_, _, oc, cmd := m.submit()
	if oc != Submit || cmd != "" {
		t.Fatalf("blocked submit = %v/%q, want Submit/empty", oc, cmd)
	}
}

// TestSubmitQuotesCompletedPathWithSpace verifies a completed path with a
// space survives submit as one shell word: the completion writes the raw
// name into the field and BuildCommand quotes it at interpolation.
func TestSubmitQuotesCompletedPathWithSpace(t *testing.T) {
	g := testStore(t, nil)
	home, _ := os.UserHomeDir()
	if err := os.WriteFile(filepath.Join(home, "annual report.txt"), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	cheat := &model.Cheat{Tool: "cat", Title: model.NewText("Cat"), Command: "cat {{file}}"}
	m := New(cheat, g)
	m.Inputs[0].SetValue("~/annual")
	m.suggest()

	tab := tea.KeyPressMsg{Code: tea.KeyTab}
	m, _, _, _ = m.Update(tab)
	if got := m.Inputs[0].Value(); got != "~/annual report.txt" {
		t.Fatalf("completion = %q, want ~/annual report.txt", got)
	}
	_, _, oc, cmd := m.submit()
	if oc != Submit {
		t.Fatalf("submit outcome = %v, want Submit", oc)
	}
	if cmd != "cat ~/'annual report.txt'" {
		t.Fatalf("submit must quote spaced path as one word: %q", cmd)
	}
}
