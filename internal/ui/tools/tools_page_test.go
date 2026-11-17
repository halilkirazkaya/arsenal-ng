// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"testing"
)

// =============================================================================
// Paging Limits
// =============================================================================

// TestNextPageClampsToLast verifies paging stops at the last page.
func TestNextPageClampsToLast(t *testing.T) {
	cheats := manyCheats(60)
	m := New(cheats, 120, 40)
	last := m.Paginator.TotalPages - 1
	for range last + 5 {
		m = m.NextPage(cheats, 120, 40)
	}
	if m.Paginator.Page != last {
		t.Fatalf("Page = %d, want last %d", m.Paginator.Page, last)
	}
	m = m.NextPage(cheats, 120, 40)
	if m.Paginator.Page != last {
		t.Fatal("NextPage past last must not move")
	}
}

// TestPrevPageStepsBack verifies one step back from the last page.
func TestPrevPageStepsBack(t *testing.T) {
	cheats := manyCheats(60)
	m := New(cheats, 120, 40)
	last := m.Paginator.TotalPages - 1
	for range last {
		m = m.NextPage(cheats, 120, 40)
	}
	m = m.PrevPage(cheats, 120, 40)
	if m.Paginator.Page != last-1 {
		t.Fatalf("PrevPage: page=%d, want %d", m.Paginator.Page, last-1)
	}
}
