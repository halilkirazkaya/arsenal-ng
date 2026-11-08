// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"
)

// =============================================================================
// Enter Pick Cases
// =============================================================================

// TestEnterPicksCheatWithAndWithoutArgs verifies quit flag by args presence.
func TestEnterPicksCheatWithAndWithoutArgs(t *testing.T) {
	oc, pick := enterWith(t, "", 0)
	if oc != Picked || !pick.Quit || pick.Cheat.Title != "Basic scan" {
		t.Fatalf("plain: oc=%v quit=%v cheat=%v", oc, pick.Quit, pick.Cheat)
	}
	oc, pick = enterWith(t, "", 1)
	if oc != Picked || pick.Quit {
		t.Fatalf("args: oc=%v quit=%v, want Picked/false", oc, pick.Quit)
	}
}

// TestEnterInvalidSelectionStays verifies invalid cursors stay in place.
func TestEnterInvalidSelectionStays(t *testing.T) {
	for _, cursor := range []int{-1, 99} {
		oc, pick := enterWith(t, "", cursor)
		if oc != Stay || pick.Cheat != nil {
			t.Errorf("cursor %d: got %v/%v, want Stay/nil", cursor, oc, pick.Cheat)
		}
	}
	m := New(nil)
	m.Cursor = 0
	if _, _, oc, _ := m.enter(nil); oc != Stay {
		t.Fatalf("empty list enter = %v, want Stay", oc)
	}
}
