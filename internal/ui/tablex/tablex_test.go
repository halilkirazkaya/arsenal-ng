// Package tablex builds the styled tables shared by the tools and
// variables views: columns, rows, focus, height, and selection styles.
package tablex

import (
	"testing"

	"charm.land/bubbles/v2/table"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// cols returns a two-column fixture for table tests.
func cols() []table.Column {
	return []table.Column{{Title: "A", Width: 10}, {Title: "B", Width: 10}}
}

// rows returns n identical two-cell rows for table tests.
func rows(n int) []table.Row {
	r := make([]table.Row, 0, n)
	for range n {
		r = append(r, table.Row{"a", "b"})
	}
	return r
}

// =============================================================================
// Height Clamping
// =============================================================================

// TestHeightFitsRowsPlusHeader verifies short tables size to rows plus header.
func TestHeightFitsRowsPlusHeader(t *testing.T) {
	if got := Height(5, 40, 10); got != 7 {
		t.Fatalf("Height = %d, want 7", got)
	}
}

// TestHeightClampsToAvailableSpace verifies tall tables clamp to the terminal.
func TestHeightClampsToAvailableSpace(t *testing.T) {
	// Twenty rows plus header exceed the space left after the reserve.
	if got := Height(20, 40, 30); got != 10 {
		t.Fatalf("Height = %d, want 10 (max available)", got)
	}
}

// TestHeightEnforcesMinimumThree verifies degenerate sizes clamp to three.
func TestHeightEnforcesMinimumThree(t *testing.T) {
	cases := [][3]int{
		{0, 40, 10},
		{5, 10, 20},
		{1, 5, 100},
		{0, 3, 3},
	}
	for _, c := range cases {
		if got := Height(c[0], c[1], c[2]); got != 3 {
			t.Errorf("Height(%d,%d,%d) = %d, want 3", c[0], c[1], c[2], got)
		}
	}
}

// =============================================================================
// Table Construction
// =============================================================================

// TestNewIsFocusedWithRowsAndHeight verifies New focuses and sizes the table.
func TestNewIsFocusedWithRowsAndHeight(t *testing.T) {
	tbl := New(cols(), rows(4), 6)
	if !tbl.Focused() {
		t.Fatal("table must be focused")
	}
	if len(tbl.Rows()) != 4 {
		t.Fatalf("rows = %d, want 4", len(tbl.Rows()))
	}
	// Height reports the viewport: requested height minus the header row.
	if tbl.Height() != 5 {
		t.Fatalf("height = %d, want 5", tbl.Height())
	}
}
