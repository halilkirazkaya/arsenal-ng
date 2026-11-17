// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"strings"
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/ui/finder"
)

// =============================================================================
// View Rendering
// =============================================================================

// TestViewsContainTitles verifies each view renders its heading or prompt.
func TestViewsContainTitles(t *testing.T) {
	a := newTestApp(t, nil)
	if v := a.render(); !strings.Contains(v, "❯") {
		t.Fatalf("search view missing prompt: %q", trunc(v))
	}
	m, _ := a.pickCheat(finder.Pick{Cheat: argCheat()})
	if v := m.(App).render(); !strings.Contains(v, "Command Preview") {
		t.Fatalf("args view missing preview: %q", trunc(v))
	}
	if v := a.showVariables().render(); !strings.Contains(v, "Global Variables") {
		t.Fatalf("variables view missing title: %q", trunc(v))
	}
	if v := a.showTools().render(); !strings.Contains(v, "Available Tools") {
		t.Fatalf("tools view missing title: %q", trunc(v))
	}
	if v := a.showHelp().render(); !strings.Contains(v, "HELP") {
		t.Fatalf("help view missing heading: %q", trunc(v))
	}
}

// TestUnknownViewEmpty verifies an invalid view renders nothing.
func TestUnknownViewEmpty(t *testing.T) {
	a := newTestApp(t, nil)
	a.view = view(42)
	if v := a.render(); v != "" {
		t.Fatalf("unknown view = %q, want empty", trunc(v))
	}
}

// TestViewSetsAltScreen verifies the root view enables the alt screen.
func TestViewSetsAltScreen(t *testing.T) {
	a := newTestApp(t, nil)
	v := a.View()
	if !v.AltScreen {
		t.Fatal("root view must enable AltScreen")
	}
	if v.Content == "" {
		t.Fatal("root view must carry rendered content")
	}
}

// trunc shortens long rendered views for failure messages.
func trunc(s string) string {
	if len(s) > 200 {
		return s[:200]
	}
	return s
}
