// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Navigation Fixtures
// =============================================================================

// manyCheats returns n identical cheats for scroll tests.
func manyCheats(n int) []*model.Cheat {
	cheats := make([]*model.Cheat, 0, n)
	for range n {
		cheats = append(cheats, &model.Cheat{
			Tool:    "tool",
			Title:   "cmd",
			Command: "run it",
		})
	}
	return cheats
}

// =============================================================================
// Navigation Limits
// =============================================================================

// TestMoveDownFromEmptySelectsFirst verifies advancing selects item zero.
func TestMoveDownFromEmptySelectsFirst(t *testing.T) {
	m := New(testCheats())
	m = m.MoveDown()
	if m.Cursor != 0 || m.Offset != 0 {
		t.Fatalf("got %d/%d, want 0/0", m.Cursor, m.Offset)
	}
}

// TestMoveDownAtEndDoesNotWrap verifies the cursor stops at the last row.
func TestMoveDownAtEndDoesNotWrap(t *testing.T) {
	m := New(testCheats())
	m.Cursor = 2
	m = m.MoveDown()
	if m.Cursor != 2 {
		t.Fatalf("Cursor = %d, want 2 (no wrap)", m.Cursor)
	}
}

// TestMoveUpFromFirstClearsSelection verifies backing off row zero clears.
func TestMoveUpFromFirstClearsSelection(t *testing.T) {
	m := New(testCheats())
	m.Cursor = 0
	m = m.MoveUp()
	if m.Cursor != -1 {
		t.Fatalf("Cursor = %d, want -1", m.Cursor)
	}
}

// TestMoveUpWithNoSelectionStays verifies MoveUp on empty selection stays.
func TestMoveUpWithNoSelectionStays(t *testing.T) {
	m := New(testCheats())
	m = m.MoveUp()
	if m.Cursor != -1 {
		t.Fatalf("Cursor = %d, want -1", m.Cursor)
	}
}

// TestMoveDownOnEmptyListStays verifies MoveDown on an empty list stays.
func TestMoveDownOnEmptyListStays(t *testing.T) {
	m := New(nil)
	m = m.MoveDown()
	if m.Cursor != -1 {
		t.Fatalf("Cursor = %d, want -1", m.Cursor)
	}
}
