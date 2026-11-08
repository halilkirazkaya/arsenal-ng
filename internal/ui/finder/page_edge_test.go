// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
)

// =============================================================================
// Page Edge Cases
// =============================================================================

// TestPageUpKeepsOffsetOnPartialPage verifies partial pages keep offset.
func TestPageUpKeepsOffsetOnPartialPage(t *testing.T) {
	// Short terminal clamps visible rows to minVisible.
	m := New(manyCheats(10)).Resize(120, 1)
	for range 6 {
		m = m.MoveDown()
	}
	m = m.PageUp()
	if m.Cursor != 2 || m.Offset != 2 {
		t.Fatalf("got %d/%d, want 2/2", m.Cursor, m.Offset)
	}
}

// TestPageDownAdvancesOffset verifies PageDown scrolls the window forward.
func TestPageDownAdvancesOffset(t *testing.T) {
	m := New(manyCheats(10)).Resize(120, 1)
	// First PageDown from no selection lands on the first row.
	m = m.PageDown()
	// Second PageDown advances one full page.
	m = m.PageDown()
	if m.Cursor != 3 || m.Offset != 1 {
		t.Fatalf("got %d/%d, want 3/1", m.Cursor, m.Offset)
	}
}

// TestUpdatePageKeys verifies PgUp/PgDn keys route through Update.
func TestUpdatePageKeys(t *testing.T) {
	m := New(testCheats()).Resize(120, 40)
	m.Cursor = 0
	nm, _, oc, _ := m.Update(keyMsg(keys.PgDown), testCheats())
	if oc != Stay || nm.Cursor != 2 {
		t.Fatalf("pgdown: oc=%v cursor=%d", oc, nm.Cursor)
	}
	nm, _, oc, _ = nm.Update(keyMsg(keys.PgUp), testCheats())
	if oc != Stay || nm.Cursor != -1 {
		t.Fatalf("pgup: oc=%v cursor=%d", oc, nm.Cursor)
	}
}
