// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	"testing"
)

// =============================================================================
// Update Mapping
// =============================================================================

// TestUpdateBackKeys verifies escape, quit, and enter leave the view.
func TestUpdateBackKeys(t *testing.T) {
	for _, k := range []string{"esc", "q", "enter"} {
		m := New(map[string]string{"a": "1"}, 40)
		if _, _, oc := m.Update(keyMsg(k)); oc != Back {
			t.Errorf("%s = %v, want Back", k, oc)
		}
	}
}

// TestUpdateFallsThroughToTable verifies other keys move the table cursor.
func TestUpdateFallsThroughToTable(t *testing.T) {
	m := New(map[string]string{"a": "1", "b": "2"}, 40)
	nm, _, oc := m.Update(keyMsg("down"))
	if oc != Stay {
		t.Fatalf("down oc = %v, want Stay", oc)
	}
	if nm.table.Cursor() == 0 {
		t.Fatal("down should move the table cursor from 0")
	}
}
