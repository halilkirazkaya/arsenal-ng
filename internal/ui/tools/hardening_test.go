// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

// =============================================================================
// Hardening Cases
// =============================================================================

// TestRebuildZeroPerPage verifies zero PerPage falls back without panicking.
func TestRebuildZeroPerPage(t *testing.T) {
	// Zero PerPage must not panic (division by zero); default applies.
	m := Model{}
	m = m.Rebuild(manyCheats(3), 100, 30)
	if m.Paginator.TotalPages != 1 {
		t.Fatalf("pages = %d, want 1", m.Paginator.TotalPages)
	}
	if got := m.Selected(); got == "" {
		t.Fatal("first tool must be selected")
	}
}

// TestUpdateNonKeyPassesToTable verifies non-key messages stay in the view.
func TestUpdateNonKeyPassesToTable(t *testing.T) {
	m := New(manyCheats(3), 100, 30)
	_, cmd, oc, tool := m.Update(tea.WindowSizeMsg{Width: 80, Height: 24}, manyCheats(3), 100, 30)
	if oc != Stay || tool != "" {
		t.Fatalf("outcome = %v/%q, want Stay/empty", oc, tool)
	}
	_ = cmd
}

// TestRebuildClampsNegativePage verifies a negative page resets to zero.
func TestRebuildClampsNegativePage(t *testing.T) {
	m := New(manyCheats(3), 100, 30)
	m.Paginator.Page = -5
	m = m.Rebuild(manyCheats(3), 100, 30)
	if m.Paginator.Page != 0 {
		t.Fatalf("page = %d, want 0", m.Paginator.Page)
	}
}
