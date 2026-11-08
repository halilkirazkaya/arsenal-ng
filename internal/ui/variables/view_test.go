// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	"strings"
	"testing"
)

// =============================================================================
// View Rendering
// =============================================================================

// TestViewEmptyShowsUsage verifies the empty view shows usage instructions.
func TestViewEmptyShowsUsage(t *testing.T) {
	v := New(nil, 40).View(nil)
	for _, want := range []string{"No variables set.", "Usage:", "set key=value", "unset key", "variables"} {
		if !strings.Contains(v, want) {
			t.Errorf("empty view missing %q", want)
		}
	}
}

// TestViewFilledShowsTable verifies the filled view lists names and values.
func TestViewFilledShowsTable(t *testing.T) {
	variables := map[string]string{"ip": "10.0.0.1"}
	v := New(variables, 40).View(variables)
	for _, want := range []string{"Global Variables", "ip", "10.0.0.1"} {
		if !strings.Contains(v, want) {
			t.Errorf("filled view missing %q", want)
		}
	}
	if strings.Contains(v, "No variables set.") {
		t.Fatal("filled view must not show usage")
	}
}
