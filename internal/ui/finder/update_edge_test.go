// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
)

// =============================================================================
// Update Edge Cases
// =============================================================================

// TestUpdateTypingFiltersAndResetsCursor verifies typing refilters the list.
func TestUpdateTypingFiltersAndResetsCursor(t *testing.T) {
	m := New(testCheats()).SetQuery("")
	// Stale selection must reset once the list narrows.
	m.Cursor = 2
	nm, _, oc, _ := m.Update(rk("f"), testCheats())
	if oc != Stay {
		t.Fatalf("typing = %v, want Stay", oc)
	}
	if len(nm.Filtered) != 1 || nm.Cursor != -1 {
		t.Fatalf("filtered=%d cursor=%d, want 1/-1", len(nm.Filtered), nm.Cursor)
	}
}

// TestUpdateClearsStatusMessage verifies keys clear stale status output.
func TestUpdateClearsStatusMessage(t *testing.T) {
	m := New(testCheats())
	m.StatusMsg = "stale"
	m.StatusIsError = true
	nm, _, _, _ := m.Update(keyMsg(keys.Down), testCheats())
	if nm.StatusMsg != "" || nm.StatusIsError {
		t.Fatal("status message must be cleared on any key")
	}
}

// TestUpdateEnterPicksPlainCommand verifies enter picks arg-less cheats.
func TestUpdateEnterPicksPlainCommand(t *testing.T) {
	m := New(testCheats())
	// First cheat has no arguments so picking quits immediately.
	m.Cursor = 0
	_, _, oc, pick := m.Update(keyMsg(keys.Enter), testCheats())
	if oc != Picked || !pick.Quit || pick.Cheat.Title != "Basic scan" {
		t.Fatalf("oc=%v quit=%v cheat=%v", oc, pick.Quit, pick.Cheat)
	}
}
