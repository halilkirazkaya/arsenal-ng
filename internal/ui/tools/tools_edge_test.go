// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"strings"
	"testing"
)

// =============================================================================
// Update Mapping
// =============================================================================

// TestUpdateBackAndPick verifies esc, quit, and enter outcomes with tool names.
func TestUpdateBackAndPick(t *testing.T) {
	cheats := manyCheats(2)
	m := New(cheats, 120, 40)
	if _, _, oc, _ := m.Update(keyMsg("esc"), cheats, 120, 40); oc != Back {
		t.Errorf("esc = %v, want Back", oc)
	}
	if _, _, oc, _ := m.Update(keyMsg("q"), cheats, 120, 40); oc != Back {
		t.Errorf("q = %v, want Back", oc)
	}
	if _, _, oc, tool := m.Update(keyMsg("enter"), cheats, 120, 40); oc != Pick || tool != "tool-000" {
		t.Errorf("enter = %v/%q, want Pick/tool-000", oc, tool)
	}
	me := New(nil, 120, 40)
	if _, _, oc, tool := me.Update(keyMsg("enter"), nil, 120, 40); oc != Pick || tool != "" {
		t.Errorf("enter empty = %v/%q, want Pick/empty", oc, tool)
	}
}

// TestUpdatePageKeys verifies vim and arrow keys flip pages both ways.
func TestUpdatePageKeys(t *testing.T) {
	cheats := manyCheats(60)
	m := New(cheats, 120, 40)
	nm, _, oc, _ := m.Update(keyMsg("l"), cheats, 120, 40)
	if oc != Stay || nm.Paginator.Page != 1 {
		t.Fatalf("l: oc=%v page=%d", oc, nm.Paginator.Page)
	}
	nm, _, oc, _ = nm.Update(keyMsg("h"), cheats, 120, 40)
	if oc != Stay || nm.Paginator.Page != 0 {
		t.Fatalf("h: oc=%v page=%d", oc, nm.Paginator.Page)
	}
	nm, _, oc, _ = nm.Update(keyMsg("right"), cheats, 120, 40)
	if oc != Stay || nm.Paginator.Page != 1 {
		t.Fatalf("right: oc=%v page=%d", oc, nm.Paginator.Page)
	}
	nm, _, oc, _ = nm.Update(keyMsg("left"), cheats, 120, 40)
	if oc != Stay || nm.Paginator.Page != 0 {
		t.Fatalf("left: oc=%v page=%d", oc, nm.Paginator.Page)
	}
}

// TestUpdateFallsThroughToTable verifies other keys move the table cursor.
func TestUpdateFallsThroughToTable(t *testing.T) {
	cheats := manyCheats(3)
	m := New(cheats, 120, 40)
	before := m.Selected()
	nm, _, oc, _ := m.Update(keyMsg("down"), cheats, 120, 40)
	if oc != Stay {
		t.Fatalf("down oc = %v, want Stay", oc)
	}
	if before != "tool-000" {
		t.Fatalf("precondition: selected=%q, want tool-000", before)
	}
	if nm.Selected() == before {
		t.Fatal("down should move table cursor")
	}
}

// =============================================================================
// View Rendering
// =============================================================================

// TestViewEmptyVsFilled verifies the empty notice and filled view contents.
func TestViewEmptyVsFilled(t *testing.T) {
	empty := New(nil, 120, 40).View(nil)
	if !strings.Contains(empty, "No tools found.") {
		t.Fatal("empty view must say no tools")
	}
	full := New(manyCheats(2), 120, 40).View(manyCheats(2))
	for _, want := range []string{"Available Tools", "tool-000", "Total: 2 tools", "ESC"} {
		if !strings.Contains(full, want) {
			t.Errorf("view missing %q", want)
		}
	}
}

// TestViewIsPure verifies rendering does not mutate the model.
func TestViewIsPure(t *testing.T) {
	m := New(manyCheats(2), 120, 40)
	before := m.Table.View()
	_ = m.View(manyCheats(2))
	if after := m.Table.View(); before != after {
		t.Fatal("View must not mutate the model")
	}
}
