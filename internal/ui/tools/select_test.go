// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"testing"
)

// =============================================================================
// Current Selection
// =============================================================================

// TestSelectedEmptyVsFilled verifies Selected on empty and filled models.
func TestSelectedEmptyVsFilled(t *testing.T) {
	if got := New(nil, 120, 40).Selected(); got != "" {
		t.Fatalf("empty Selected = %q", got)
	}
	if got := New(manyCheats(2), 120, 40).Selected(); got == "" {
		t.Fatal("filled Selected must not be empty")
	}
}
