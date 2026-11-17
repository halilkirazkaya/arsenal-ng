// Package tablex builds the styled tables shared by the tools and
// variables views: columns, rows, focus, height, and selection styles.
package tablex

import (
	"testing"

	"charm.land/bubbles/v2/table"
)

// =============================================================================
// Table Styles
// =============================================================================

// TestStylesSelectedDiffersFromBase verifies shared styles override defaults.
func TestStylesSelectedDiffersFromBase(t *testing.T) {
	base := table.DefaultStyles()
	got := Styles()
	// Compare style attributes since plain Render carries no ANSI to observe.
	if got.Selected.GetForeground() == base.Selected.GetForeground() &&
		got.Selected.GetBackground() == base.Selected.GetBackground() {
		t.Fatal("Selected style must differ from bubbles default")
	}
	if got.Header.GetForeground() == base.Header.GetForeground() {
		t.Fatal("Header style must differ from bubbles default")
	}
	if got.Selected.GetBackground() == got.Header.GetBackground() &&
		got.Selected.GetForeground() == got.Header.GetForeground() {
		t.Fatal("Selected and Header styles must differ from each other")
	}
}
