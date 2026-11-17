// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Tools Flow Tests
// =============================================================================

// TestToolsPickFiltersSearch verifies picking a tool filters the search.
func TestToolsPickFiltersSearch(t *testing.T) {
	a := newTestApp(t, nil).showTools()
	m, _ := a.Update(keyEnter())
	got := m.(App)
	if got.view != viewSearch {
		t.Fatal("pick must return to search")
	}
	if got.search.Input.Value() != "nmap" {
		t.Fatalf("query = %q, want selected tool", got.search.Input.Value())
	}
	if len(got.search.Filtered) != 1 {
		t.Fatalf("filtered = %d, want 1", len(got.search.Filtered))
	}
}

// TestToolsEmptyPickGoesBack verifies an empty tools pick returns home.
func TestToolsEmptyPickGoesBack(t *testing.T) {
	a := newTestApp(t, []*model.Cheat{}).showTools()
	m, _ := a.Update(keyEnter())
	if m.(App).view != viewSearch {
		t.Fatal("empty pick must go back to search")
	}
}

// TestToolsBack verifies escape leaves the tools view.
func TestToolsBack(t *testing.T) {
	a := newTestApp(t, nil).showTools()
	m, _ := a.Update(keyEsc())
	if m.(App).view != viewSearch {
		t.Fatal("esc must leave tools view")
	}
}

// =============================================================================
// Variables Flow Tests
// =============================================================================

// TestVarsBack verifies enter leaves the variables view.
func TestVarsBack(t *testing.T) {
	a := newTestApp(t, nil).showVariables()
	m, _ := a.Update(keyEnter())
	if m.(App).view != viewSearch {
		t.Fatal("enter must leave variables view")
	}
}

// =============================================================================
// Prefill Tests
// =============================================================================

// TestCountApplied verifies stored variables count toward prefill coverage.
func TestCountApplied(t *testing.T) {
	cheat := &model.Cheat{Command: "nmap {{ip}} -p {{port}}"}
	if n := countApplied(cheat, newTestApp(t, nil).globals); n != 0 {
		t.Fatalf("count = %d, want 0", n)
	}
	b := newTestApp(t, nil)
	if err := b.globals.Set("ip", "1.1.1.1"); err != nil {
		t.Fatal(err)
	}
	if n := countApplied(cheat, b.globals); n != 1 {
		t.Fatalf("count = %d, want 1", n)
	}
	if n := countApplied(&model.Cheat{Command: "ping google.com"}, b.globals); n != 0 {
		t.Fatalf("arg-less count = %d, want 0", n)
	}
}
