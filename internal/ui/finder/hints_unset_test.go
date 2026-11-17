// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"strings"
	"testing"
)

// =============================================================================
// Unset Hint Order
// =============================================================================

// TestUnsetHintsListsNamesAlphabetically verifies multi-var output order.
func TestUnsetHintsListsNamesAlphabetically(t *testing.T) {
	variables := map[string]string{"zebra": "1", "apple": "2", "mango": "3"}
	got := unsetHints(variables)
	apple := strings.Index(got, "apple")
	mango := strings.Index(got, "mango")
	zebra := strings.Index(got, "zebra")
	if apple < 0 || mango < 0 || zebra < 0 {
		t.Fatalf("unset hints must list all names, got %q", got)
	}
	if !(apple < mango && mango < zebra) {
		t.Fatalf("unset hints not alphabetical, got %q", got)
	}
}

// TestUnsetHintsRenderIsDeterministic verifies repeated renders match.
func TestUnsetHintsRenderIsDeterministic(t *testing.T) {
	variables := map[string]string{"zebra": "1", "apple": "2", "mango": "3"}
	first := unsetHints(variables)
	for range 20 {
		if got := unsetHints(variables); got != first {
			t.Fatalf("repeated render differs:\nfirst %q\ngot   %q", first, got)
		}
	}
}
