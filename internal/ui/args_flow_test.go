// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/finder"
)

// =============================================================================
// Args Flow Tests
// =============================================================================

// argsApp builds an app already inside the args view with seeded globals.
func argsApp(t *testing.T, cheat *model.Cheat, seed map[string]string) App {
	t.Helper()
	a := newTestApp(t, nil)
	for k, v := range seed {
		if err := a.globals.Set(k, v); err != nil {
			t.Fatal(err)
		}
	}
	m, _ := a.pickCheat(finder.Pick{Cheat: cheat})
	return m.(App)
}

// TestArgsSubmitQuits verifies submitting filled args quits with a command.
func TestArgsSubmitQuits(t *testing.T) {
	a := argsApp(t, argCheat(), map[string]string{"ip": "9.9.9.9"})
	m, cmd := a.Update(keyEnter())
	got := m.(App)
	if got.FinalCommand != "nmap 9.9.9.9" {
		t.Fatalf("FinalCommand = %q", got.FinalCommand)
	}
	if !isQuit(cmd) {
		t.Fatal("submit must quit")
	}
}

// TestArgsSubmitBlockedWhenEmpty verifies empty args block submission.
func TestArgsSubmitBlockedWhenEmpty(t *testing.T) {
	a := argsApp(t, argCheat(), nil)
	m, cmd := a.Update(keyEnter())
	got := m.(App)
	if got.FinalCommand != "" {
		t.Fatal("empty args must not produce a command")
	}
	if isQuit(cmd) {
		t.Fatal("blocked submit must not quit")
	}
	if got.view != viewArgs {
		t.Fatal("blocked submit must stay in args")
	}
}

// TestArgsBackAndCancel verifies esc returns home and ctrl+c cancels.
func TestArgsBackAndCancel(t *testing.T) {
	a := argsApp(t, argCheat(), nil)
	m, _ := a.Update(keyEsc())
	if m.(App).view != viewSearch {
		t.Fatal("esc must return to search")
	}
	b := argsApp(t, argCheat(), nil)
	m2, cmd := b.Update(keyCtrlC())
	got := m2.(App)
	if !got.Cancelled || !isQuit(cmd) {
		t.Fatal("ctrl+c must cancel and quit")
	}
}
