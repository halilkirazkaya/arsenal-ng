// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Construction Tests
// =============================================================================

// TestNewStartsInSearch verifies New opens on search with defaults.
func TestNewStartsInSearch(t *testing.T) {
	a := newTestApp(t, nil)
	if a.view != viewSearch {
		t.Fatalf("view = %v, want viewSearch", a.view)
	}
	if len(a.search.Filtered) != 2 {
		t.Fatalf("filtered = %d, want 2", len(a.search.Filtered))
	}
	if a.width != config.DefaultWidth || a.height != config.DefaultHeight {
		t.Fatalf("dims = %dx%d", a.width, a.height)
	}
}

// TestNewEmptyCheats verifies New handles an empty cheat list.
func TestNewEmptyCheats(t *testing.T) {
	a := newTestApp(t, []*model.Cheat{})
	if len(a.search.Filtered) != 0 {
		t.Fatalf("filtered = %d, want 0", len(a.search.Filtered))
	}
}

// TestInitReturnsBlink verifies Init returns the cursor blink command.
func TestInitReturnsBlink(t *testing.T) {
	a := newTestApp(t, nil)
	if a.Init() == nil {
		t.Fatal("Init must return cursor blink cmd")
	}
}
