// Package tablex builds the styled tables shared by the tools and
// variables views: columns, rows, focus, height, and selection styles.
package tablex

import (
	"testing"
)

// =============================================================================
// Width Fitting
// =============================================================================

// TestFitFillsTarget verifies fitted columns sum to the target width.
func TestFitFillsTarget(t *testing.T) {
	fitted := Fit(cols(), 116)
	if got := Width(fitted); got != 116 {
		t.Fatalf("Width = %d, want 116", got)
	}
}

// TestFitLeavesInputUntouched verifies the input slice keeps base widths.
func TestFitLeavesInputUntouched(t *testing.T) {
	base := cols()
	_ = Fit(base, 116)
	for i, col := range cols() {
		if base[i].Width != col.Width {
			t.Fatalf("col %d = %d, want %d", i, base[i].Width, col.Width)
		}
	}
}

// TestFitKeepsBaseOnNarrow verifies tiny targets keep the base widths.
func TestFitKeepsBaseOnNarrow(t *testing.T) {
	base := cols()
	fitted := Fit(base, 10)
	for i := range base {
		if fitted[i].Width != base[i].Width {
			t.Fatalf("col %d = %d, want %d", i, fitted[i].Width, base[i].Width)
		}
	}
}

// TestNewWideSetsViewport verifies viewport and columns match the target.
func TestNewWideSetsViewport(t *testing.T) {
	tbl := NewWide(cols(), rows(4), 6, 116)
	if tbl.Width() != 116 {
		t.Fatalf("viewport = %d, want 116", tbl.Width())
	}
	if got := Width(tbl.Columns()); got != 116 {
		t.Fatalf("columns = %d, want 116", got)
	}
}

// TestNewWideNarrowRenders verifies tiny targets render without panic.
func TestNewWideNarrowRenders(t *testing.T) {
	tbl := NewWide(cols(), rows(2), 4, 10)
	if len(tbl.Rows()) != 2 {
		t.Fatalf("rows = %d, want 2", len(tbl.Rows()))
	}
	_ = tbl.View()
}

// TestNewWideZeroTargetFallsBack verifies zero keeps the base widths.
func TestNewWideZeroTargetFallsBack(t *testing.T) {
	tbl := NewWide(cols(), rows(1), 3, 0)
	if got := Width(tbl.Columns()); got != Width(cols()) {
		t.Fatalf("columns = %d, want %d", got, Width(cols()))
	}
	_ = tbl.View()
}
