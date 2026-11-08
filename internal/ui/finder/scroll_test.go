// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"
)

// =============================================================================
// Scroll Offset
// =============================================================================

// TestMoveDownScrollsOffset verifies scrolling follows the cursor down.
func TestMoveDownScrollsOffset(t *testing.T) {
	// Short terminal clamps visible rows to minVisible.
	m := New(manyCheats(10)).Resize(120, 1)
	for range 5 {
		m = m.MoveDown()
	}
	if m.Cursor != 4 || m.Offset != 2 {
		t.Fatalf("got %d/%d, want 4/2", m.Cursor, m.Offset)
	}
}

// TestMoveUpPullsOffsetBack verifies scrolling follows the cursor up.
func TestMoveUpPullsOffsetBack(t *testing.T) {
	m := New(manyCheats(10)).Resize(120, 1)
	for range 5 {
		m = m.MoveDown()
	}
	for range 3 {
		m = m.MoveUp()
	}
	if m.Cursor != 1 || m.Offset != 1 {
		t.Fatalf("got %d/%d, want 1/1", m.Cursor, m.Offset)
	}
}
