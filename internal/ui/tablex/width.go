// Package tablex builds the styled tables shared by the tools and
// variables views: columns, rows, focus, height, and selection styles.
package tablex

import (
	"charm.land/bubbles/v2/table"
)

// =============================================================================
// Width Fitting
// =============================================================================

// NewWide creates a focused table stretched to the target viewport width.
func NewWide(columns []table.Column, rows []table.Row, height, target int) table.Model {
	t := New(Fit(columns, target), rows, height)
	if target > 0 {
		t.SetWidth(target)
	}
	return t
}

// Fit stretches the last column so the columns fill the target width.
func Fit(columns []table.Column, target int) []table.Column {
	fitted := make([]table.Column, len(columns))
	copy(fitted, columns)
	if target <= 0 || len(fitted) == 0 {
		return fitted
	}
	fixed := 0
	for _, col := range fitted[:len(fitted)-1] {
		fixed += col.Width + 2
	}
	last := target - fixed - 2
	// Narrow targets keep the base widths: zero widths would drop columns.
	if last < 1 {
		return fitted
	}
	fitted[len(fitted)-1].Width = last
	return fitted
}
