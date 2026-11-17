// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"strings"
	"testing"
)

// =============================================================================
// View Branches
// =============================================================================

// TestViewNormalListShowsRowsAndFooter verifies the list renders rows.
func TestViewNormalListShowsRowsAndFooter(t *testing.T) {
	m := New(testCheats()).Resize(120, 40)
	m.Cursor = 0
	m.Filtered = testCheats()
	v := m.View(nil)
	for _, want := range []string{"Basic scan", "ffuf", "1/3", "❯"} {
		if !strings.Contains(v, want) {
			t.Errorf("view missing %q", want)
		}
	}
}

// TestViewStatusErrorVsSuccess verifies both status styles render text.
func TestViewStatusErrorVsSuccess(t *testing.T) {
	m := New(testCheats()).Resize(120, 40)
	m.StatusMsg = "boom"
	m.StatusIsError = true
	if !strings.Contains(m.View(nil), "boom") {
		t.Fatal("error status must render")
	}
	m.StatusIsError = false
	m.StatusMsg = "ok"
	if !strings.Contains(m.View(nil), "ok") {
		t.Fatal("success status must render")
	}
}

// TestViewSelectedShowsInfoBox verifies the info box follows selection.
func TestViewSelectedShowsInfoBox(t *testing.T) {
	m := New(testCheats()).Resize(120, 40)
	m.Cursor = 0
	if !strings.Contains(m.View(nil), "Scan a host") {
		t.Fatal("selected cheat desc should appear in InfoBox")
	}
	m.Cursor = -1
	if strings.Contains(m.View(nil), "Scan a host") {
		t.Fatal("unselected view must not show InfoBox")
	}
}

// TestViewFooterCursorMath verifies cursor math in the view footer.
func TestViewFooterCursorMath(t *testing.T) {
	m := New(testCheats())
	if f := m.Footer(2); !strings.Contains(f, "0/3") {
		t.Fatalf("unselected footer = %q, want 0/3", f)
	}
	m.Cursor = 1
	if f := m.Footer(0); !strings.Contains(f, "2/3") {
		t.Fatalf("selected footer = %q, want 2/3", f)
	}
}

// TestViewSpecialCommandShowsHints verifies hints replace the cheat list.
func TestViewSpecialCommandShowsHints(t *testing.T) {
	m := New(testCheats()).Resize(120, 40).SetQuery("set ")
	v := m.View(nil)
	if !strings.Contains(v, "Common variables") || strings.Contains(v, "Basic scan") {
		t.Fatal("hints should replace list for special commands")
	}
}
