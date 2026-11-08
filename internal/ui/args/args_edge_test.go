// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

// =============================================================================
// Update Key Fixtures
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
	case "tab":
		return tea.KeyPressMsg{Code: tea.KeyTab}
	case "shift+tab":
		return tea.KeyPressMsg{Code: tea.KeyTab, Mod: tea.ModShift}
	case "ctrl+c":
		return tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl}
	}
	return tea.KeyPressMsg{Text: string(s)}
}

// =============================================================================
// Update Key Mapping
// =============================================================================

// TestUpdateQuitBackKeys verifies cancel keys return Cancel and Back.
func TestUpdateQuitBackKeys(t *testing.T) {
	g := testStore(t, nil)
	newM := func() Model { return New(plainCheat(), g) }

	if _, _, oc, _ := newM().Update(keyMsg("ctrl+c")); oc != Cancel {
		t.Errorf("ctrl+c = %v, want Cancel", oc)
	}
	if _, _, oc, _ := newM().Update(keyMsg("esc")); oc != Back {
		t.Errorf("esc = %v, want Back", oc)
	}
}

// TestUpdateEnterSubmitsOrBlocks verifies enter submits or blocks on empty.
func TestUpdateEnterSubmitsOrBlocks(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	m.Inputs[0].SetValue("1.1.1.1")
	m.Inputs[1].SetValue("80")
	if _, _, oc, cmd := m.Update(keyMsg("enter")); oc != Submit || cmd == "" {
		t.Errorf("enter = %v/%q, want Submit/command", oc, cmd)
	}
	m2 := New(plainCheat(), g)
	if _, _, oc, cmd := m2.Update(keyMsg("enter")); oc != Submit || cmd != "" {
		t.Errorf("enter blocked = %v/%q", oc, cmd)
	}
}

// TestUpdateTabCyclesFocus verifies tab and arrow keys cycle focus.
func TestUpdateTabCyclesFocus(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	nm, _, oc, _ := m.Update(keyMsg("tab"))
	if oc != Stay || nm.Cursor != 1 {
		t.Fatalf("tab: %v/%d", oc, nm.Cursor)
	}
	nm, _, oc, _ = nm.Update(keyMsg("shift+tab"))
	if oc != Stay || nm.Cursor != 0 {
		t.Fatalf("shift+tab: %v/%d", oc, nm.Cursor)
	}
	nm, _, oc, _ = m.Update(keyMsg("down"))
	if oc != Stay || nm.Cursor != 1 {
		t.Fatalf("down: %v/%d", oc, nm.Cursor)
	}
	nm, _, oc, _ = nm.Update(keyMsg("up"))
	if oc != Stay || nm.Cursor != 0 {
		t.Fatalf("up: %v/%d", oc, nm.Cursor)
	}
}

// TestUpdateTypingReachesFocusedInput verifies typing edits focused input.
func TestUpdateTypingReachesFocusedInput(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	nm, _, oc, _ := m.Update(tea.KeyPressMsg{Text: "x", Code: 'x'})
	if oc != Stay || !strings.Contains(nm.Inputs[0].Value(), "x") {
		t.Errorf("typing = %v/%q, want Stay/x inside", oc, nm.Inputs[0].Value())
	}
}
