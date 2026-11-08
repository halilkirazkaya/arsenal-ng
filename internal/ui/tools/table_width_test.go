// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"testing"
)

// =============================================================================
// Table Width
// =============================================================================

// TestNewTableStretchesToTerminal verifies the table fills width minus margin.
func TestNewTableStretchesToTerminal(t *testing.T) {
	m := New(manyCheats(2), 120, 40)
	if got := m.Table.Width(); got != 116 {
		t.Fatalf("viewport = %d, want 116", got)
	}
	sum := 0
	for _, col := range m.Table.Columns() {
		sum += col.Width + 2
	}
	if sum != 116 {
		t.Fatalf("columns = %d, want 116", sum)
	}
}

// TestNewTableNarrowRenders verifies a narrow terminal renders without panic.
func TestNewTableNarrowRenders(t *testing.T) {
	cheats := manyCheats(2)
	m := New(cheats, 10, 40)
	if len(m.Table.Rows()) != 2 {
		t.Fatalf("rows = %d, want 2", len(m.Table.Rows()))
	}
	_ = m.Table.View()
	_ = m.View(cheats)
}

// TestNewTableKeepsRowHeight verifies stretching preserves the row count.
func TestNewTableKeepsRowHeight(t *testing.T) {
	m := New(manyCheats(3), 120, 40)
	if len(m.Table.Rows()) != 3 {
		t.Fatalf("rows = %d, want 3", len(m.Table.Rows()))
	}
}
