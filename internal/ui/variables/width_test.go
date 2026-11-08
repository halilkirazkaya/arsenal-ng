// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	"testing"
)

// =============================================================================
// Table Width
// =============================================================================

// TestNewWithWidthStretches verifies the table fills width minus margin.
func TestNewWithWidthStretches(t *testing.T) {
	vars := map[string]string{"ip": "10.0.0.1"}
	m := New(vars, 40, 120)
	sum := 0
	for _, col := range m.table.Columns() {
		sum += col.Width + 2
	}
	if sum != 116 {
		t.Fatalf("columns = %d, want 116", sum)
	}
	if m.table.Width() != 116 {
		t.Fatalf("viewport = %d, want 116", m.table.Width())
	}
}

// TestNewWithoutWidthKeepsBase verifies omitted width keeps base columns.
func TestNewWithoutWidthKeepsBase(t *testing.T) {
	m := New(map[string]string{"a": "1"}, 40)
	sum := 0
	for _, col := range m.table.Columns() {
		sum += col.Width + 2
	}
	if sum != 74 {
		t.Fatalf("columns = %d, want 74", sum)
	}
}

// TestNewNarrowRenders verifies a narrow terminal renders without panic.
func TestNewNarrowRenders(t *testing.T) {
	vars := map[string]string{"ip": "10.0.0.1"}
	m := New(vars, 40, 10)
	if len(m.table.Rows()) != 1 {
		t.Fatalf("rows = %d, want 1", len(m.table.Rows()))
	}
	_ = m.table.View()
	_ = m.View(vars)
}
