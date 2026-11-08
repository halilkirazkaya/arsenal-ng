// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"strings"
	"testing"
)

// =============================================================================
// Field Navigation
// =============================================================================

// TestNextWrapsForward verifies Next advances and wraps to field zero.
func TestNextWrapsForward(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	m = m.Next()
	if m.Cursor != 1 || !m.Inputs[1].Focused() {
		t.Fatalf("Next: cursor=%d", m.Cursor)
	}
	m = m.Next()
	if m.Cursor != 0 {
		t.Fatalf("Next wraps: cursor=%d, want 0", m.Cursor)
	}
}

// TestPrevWrapsBackward verifies Prev retreats and wraps to the last.
func TestPrevWrapsBackward(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	// Prev from the first field wraps to the last one.
	m = m.Prev()
	if m.Cursor != 1 {
		t.Fatalf("Prev wraps: cursor=%d, want 1", m.Cursor)
	}
	m = m.Prev()
	if m.Cursor != 0 {
		t.Fatalf("Prev: cursor=%d, want 0", m.Cursor)
	}
}

// =============================================================================
// Args Rendering
// =============================================================================

// TestViewShowsPreviewRowsFooter verifies the form renders all sections.
func TestViewShowsPreviewRowsFooter(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	m.Inputs[0].SetValue("10.0.0.1")
	v := m.View(120)
	for _, want := range []string{"Command Preview", "Arguments:", "ip", "port", "tab", "esc: back"} {
		if !strings.Contains(v, want) {
			t.Errorf("view missing %q", want)
		}
	}
	if !strings.Contains(v, "▸") {
		t.Fatal("focused row must show cursor")
	}
}
