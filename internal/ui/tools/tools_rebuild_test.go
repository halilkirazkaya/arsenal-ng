// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"testing"
)

// =============================================================================
// Rebuild Page Handling
// =============================================================================

// TestRebuildPreservesPage verifies rebuild keeps the current page.
func TestRebuildPreservesPage(t *testing.T) {
	m := New(manyCheats(60), 120, 40)
	if m.Paginator.TotalPages < 2 {
		t.Fatalf("need 2+ pages, got %d", m.Paginator.TotalPages)
	}
	m = m.NextPage(manyCheats(60), 120, 40)
	m = m.Rebuild(manyCheats(60), 120, 40)
	if m.Paginator.Page != 1 {
		t.Fatalf("Page = %d, want 1 preserved", m.Paginator.Page)
	}
}

// TestRebuildClampsPageWhenShrinking verifies rebuild clamps a stale page.
func TestRebuildClampsPageWhenShrinking(t *testing.T) {
	m := New(manyCheats(60), 120, 40)
	m = m.NextPage(manyCheats(60), 120, 40)
	m = m.NextPage(manyCheats(60), 120, 40)
	// Now single page.
	m = m.Rebuild(manyCheats(2), 120, 40)
	if m.Paginator.Page != 0 {
		t.Fatalf("Page = %d, want 0 clamped", m.Paginator.Page)
	}
}

// TestPrevPageAtFirstStays verifies PrevPage is a no-op on page zero.
func TestPrevPageAtFirstStays(t *testing.T) {
	cheats := manyCheats(60)
	m := New(cheats, 120, 40)
	first := m.Selected()
	// Already first.
	m = m.PrevPage(cheats, 120, 40)
	if m.Paginator.Page != 0 || m.Selected() != first {
		t.Fatalf("PrevPage at 0 moved: page=%d", m.Paginator.Page)
	}
}
