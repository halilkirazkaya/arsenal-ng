// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
)

// =============================================================================
// Update Navigation
// =============================================================================

// TestUpdateNavigationKeys verifies arrow and ctrl keys move the cursor.
func TestUpdateNavigationKeys(t *testing.T) {
	m := New(testCheats())
	m.Cursor = 1
	nm, _, oc, _ := m.Update(keyMsg(keys.Down), testCheats())
	if oc != Stay || nm.Cursor != 2 {
		t.Fatalf("down: oc=%v cursor=%d", oc, nm.Cursor)
	}
	nm, _, oc, _ = m.Update(keyMsg(keys.Up), testCheats())
	if oc != Stay || nm.Cursor != 0 {
		t.Fatalf("up: oc=%v cursor=%d", oc, nm.Cursor)
	}
	nm, _, oc, _ = m.Update(keyMsg(keys.CtrlN), testCheats())
	if oc != Stay || nm.Cursor != 2 {
		t.Fatalf("ctrl+n: oc=%v cursor=%d", oc, nm.Cursor)
	}
	nm, _, oc, _ = m.Update(keyMsg(keys.CtrlP), testCheats())
	if oc != Stay || nm.Cursor != 0 {
		t.Fatalf("ctrl+p: oc=%v cursor=%d", oc, nm.Cursor)
	}
}

// TestUpdateQuitKeys verifies cancel keys return the Cancel outcome.
func TestUpdateQuitKeys(t *testing.T) {
	m := New(testCheats())
	if _, _, oc, _ := m.Update(keyMsg(keys.CtrlC), testCheats()); oc != Cancel {
		t.Errorf("ctrl+c = %v, want Cancel", oc)
	}
	if _, _, oc, _ := m.Update(keyMsg(keys.Esc), testCheats()); oc != Cancel {
		t.Errorf("esc = %v, want Cancel", oc)
	}
}

// TestUpdateHelpKeyNeedsEmptyQuery verifies help only opens on empty query.
func TestUpdateHelpKeyNeedsEmptyQuery(t *testing.T) {
	m := New(testCheats()).SetQuery("")
	if _, _, oc, _ := m.Update(keyMsg("?"), testCheats()); oc != ShowHelp {
		t.Errorf("? on empty = %v, want ShowHelp", oc)
	}
	m = New(testCheats()).SetQuery("nmap")
	if _, _, oc, _ := m.Update(keyMsg("?"), testCheats()); oc != Stay {
		t.Errorf("? on query = %v, want Stay", oc)
	}
}
