// Package theme provides the shared lipgloss styles for the TUI: base
// palette, component styles, command syntax styles, and tag rendering.
package theme

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

// =============================================================================
// Tag Color Tests
// =============================================================================

// TestColoredTags verifies empty, single, and multi-tag rendering.
func TestColoredTags(t *testing.T) {
	if got := ColoredTags(nil); got != "" {
		t.Fatalf("ColoredTags(nil) = %q, want empty", got)
	}
	if got := ColoredTags([]string{}); got != "" {
		t.Fatalf("ColoredTags(empty) = %q, want empty", got)
	}

	single := ColoredTags([]string{"scan"})
	if !strings.Contains(single, "scan") {
		t.Fatalf("ColoredTags single missing tag: %q", single)
	}
	if strings.Contains(single, ",") {
		t.Fatalf("single tag must have no separator: %q", single)
	}

	multi := ColoredTags([]string{"scan", "recon", "web"})
	for _, tag := range []string{"scan", "recon", "web"} {
		if !strings.Contains(multi, tag) {
			t.Fatalf("ColoredTags missing %q: %q", tag, multi)
		}
	}
	// Three tags need exactly two separators between them.
	if got := strings.Count(multi, ","); got != 2 {
		t.Fatalf("ColoredTags separator count = %d, want 2: %q", got, multi)
	}
}

// TestTagColorDeterministic verifies stable colors and overflow handling.
func TestTagColorDeterministic(t *testing.T) {
	if tagColor("scan") != tagColor("scan") {
		t.Fatal("tagColor not deterministic")
	}
	seen := map[string]bool{}
	for _, tag := range []string{"scan", "recon", "web", "brute", "dns"} {
		seen[tagColor(tag)] = true
	}
	if len(seen) < 2 {
		t.Fatalf("expected distinct colors across tags, got %v", seen)
	}
	// Long input overflows the hash negative yet must still map.
	if got := tagColor("zzzzzzzzzzzzzz"); got == "" {
		t.Fatal("negative hash produced no color")
	}
}

// TestSetLogoColorDeterministic verifies logo overrides and pins a color.
func TestSetLogoColorDeterministic(t *testing.T) {
	SetLogoColor(lipgloss.Color("#FF0000"))
	if got := Logo.GetForeground(); got != lipgloss.Color("#FF0000") {
		t.Fatalf("Logo foreground = %v, want #FF0000", got)
	}
	SetLogoColor(lipgloss.Color("#00FF00"))
	if got := Logo.GetForeground(); got != lipgloss.Color("#00FF00") {
		t.Fatalf("Logo foreground = %v, want #00FF00", got)
	}
	// Fixed color keeps other tests and snapshots deterministic.
	SetLogoColor(lipgloss.Color("#FF6B6B"))
}
