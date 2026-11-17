// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"testing"
)

// =============================================================================
// Page Navigation
// =============================================================================

// TestPageDownFromEmptySelectsFirst verifies PageDown selects item zero.
func TestPageDownFromEmptySelectsFirst(t *testing.T) {
	m := New(testCheats()).Resize(120, 40)
	m = m.PageDown()
	if m.Cursor != 0 {
		t.Fatalf("Cursor = %d, want 0", m.Cursor)
	}
}

// TestPageUpWithNoSelectionStays verifies PageUp on empty selection stays.
func TestPageUpWithNoSelectionStays(t *testing.T) {
	m := New(testCheats()).Resize(120, 40)
	m = m.PageUp()
	if m.Cursor != -1 || m.Offset != 0 {
		t.Fatalf("got %d/%d, want -1/0", m.Cursor, m.Offset)
	}
}

// TestPageDownClampsToLast verifies PageDown overshoot clamps to last.
func TestPageDownClampsToLast(t *testing.T) {
	// A tall terminal makes one page larger than the list.
	m := New(testCheats()).Resize(120, 40)
	m.Cursor = 0
	m = m.PageDown()
	if m.Cursor != 2 {
		t.Fatalf("Cursor = %d, want 2 (clamped)", m.Cursor)
	}
}

// TestPageUpOvershootGoesToTop verifies PageUp overshoot clears selection.
func TestPageUpOvershootGoesToTop(t *testing.T) {
	m := New(testCheats()).Resize(120, 40)
	m.Cursor = 1
	m = m.PageUp()
	if m.Cursor != -1 || m.Offset != 0 {
		t.Fatalf("got %d/%d, want -1/0", m.Cursor, m.Offset)
	}
}
