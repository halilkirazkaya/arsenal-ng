// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Cheat Row Selection
// =============================================================================

// selectedRowWidth returns the visible width of a row without ANSI codes.
func selectedRowWidth(t *testing.T, row string) int {
	t.Helper()
	// Trailing newline is a row separator, not part of the band.
	return lipgloss.Width(strings.TrimSuffix(row, "\n"))
}

// TestCheatRowSelectedSpansFullWidth verifies the band fills the view.
func TestCheatRowSelectedSpansFullWidth(t *testing.T) {
	for _, width := range []int{80, 120} {
		m := New(testCheats()).Resize(width, 40)
		m.Cursor = 0
		m.Filtered = testCheats()
		row := m.cheatRow(0, m.Filtered[0], "", lipgloss.NewStyle())
		if got := selectedRowWidth(t, row); got != m.Width() {
			t.Fatalf("width %d: selected band = %d, want %d", width, got, m.Width())
		}
	}
}

// TestCheatRowSelectedCoversWholeLine verifies tags sit inside the band.
func TestCheatRowSelectedCoversWholeLine(t *testing.T) {
	m := New(testCheats()).Resize(100, 40)
	m.Cursor = 0
	m.Filtered = testCheats()
	row := m.cheatRow(0, m.Filtered[0], "", lipgloss.NewStyle())
	stripped := stripANSI(row)
	for _, want := range []string{"nmap", "Basic scan", "scan"} {
		if !strings.Contains(stripped, want) {
			t.Fatalf("selected band missing %q: %q", want, stripped)
		}
	}
	if got := selectedRowWidth(t, row); got != m.Width() {
		t.Fatalf("selected band = %d, want %d", got, m.Width())
	}
}

// TestCheatRowSelectedDiffersFromUnselected verifies only one band renders.
func TestCheatRowSelectedDiffersFromUnselected(t *testing.T) {
	m := New(testCheats()).Resize(100, 40)
	m.Cursor = 0
	m.Filtered = testCheats()
	match := lipgloss.NewStyle()
	sel := m.cheatRow(0, m.Filtered[0], "", match)
	other := m.cheatRow(1, m.Filtered[1], "", match)
	if sel == other {
		t.Fatal("selected and unselected rows must differ")
	}
	// Only the selected row fills the full view width.
	if got := selectedRowWidth(t, other); got == m.Width() {
		t.Fatalf("unselected row must not span full width: %d", got)
	}
}

// TestCheatRowUnselectedKeepsColors verifies plain rows keep tag colors.
func TestCheatRowUnselectedKeepsColors(t *testing.T) {
	m := New(testCheats()).Resize(100, 40)
	m.Cursor = 0
	m.Filtered = testCheats()
	match := lipgloss.NewStyle()
	row := m.cheatRow(1, m.Filtered[1], "", match)
	if want := theme.ColoredTags(m.Filtered[1].Tags); !strings.Contains(row, want) {
		t.Fatalf("unselected row lost colored tags:\n%q\nwant %q", row, want)
	}
}

// stripANSI removes escape sequences for plain-text assertions.
func stripANSI(s string) string {
	var b strings.Builder
	inSeq := false
	for _, r := range s {
		switch {
		case r == 0x1b:
			inSeq = true
		case inSeq && (r == 'm' || r == 'K'):
			inSeq = false
		case !inSeq:
			b.WriteRune(r)
		}
	}
	return b.String()
}
