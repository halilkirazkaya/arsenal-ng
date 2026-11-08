// Package help provides the help view: keyboard shortcuts, usage
// examples, and workflow instructions.
package help

import (
	"strings"
	"testing"
)

// =============================================================================
// View Sections
// =============================================================================

// TestViewContainsAllSections verifies every help section renders.
func TestViewContainsAllSections(t *testing.T) {
	v := View()
	for _, want := range []string{
		"HELP", "NAVIGATION", "ARGUMENT INPUT", "GLOBAL VARIABLES",
		"EXAMPLE WORKFLOW", "set key=value", "unset key", "variables", "tools",
		"Press ESC or Enter to go back",
	} {
		if !strings.Contains(v, want) {
			t.Errorf("help view missing %q", want)
		}
	}
}
