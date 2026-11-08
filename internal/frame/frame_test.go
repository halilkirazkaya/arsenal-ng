// Package frame renders the shared application frame: header with logo
// and version, plus the search view footer with status and help text.
package frame

import (
	"strings"
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Width Tests
// =============================================================================

// TestWidth verifies usable widths pass through and tiny ones clamp.
func TestWidth(t *testing.T) {
	tests := []struct {
		name  string
		width int
		want  int
	}{
		{"normal passes through", 120, 120},
		{"boundary minimum", 40, 40},
		{"below minimum clamps", 39, config.DefaultWidth},
		{"zero clamps", 0, config.DefaultWidth},
		{"negative clamps", -10, config.DefaultWidth},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Width(tc.width); got != tc.want {
				t.Fatalf("Width(%d) = %d, want %d", tc.width, got, tc.want)
			}
		})
	}
}

// =============================================================================
// Section Tests
// =============================================================================

// TestHeaderContainsLogoAndVersion verifies the header shows branding.
func TestHeaderContainsLogoAndVersion(t *testing.T) {
	got := Header()
	// First logo row proves the banner rendered, not just the version.
	if !strings.Contains(got, "█████╗") {
		t.Fatalf("Header missing logo art: %q", got)
	}
	if !strings.Contains(got, config.GetVersionInfo()) {
		t.Fatalf("Header missing version %q: %q", config.GetVersionInfo(), got)
	}
}

// TestFooter verifies status counts and conditional variable badges.
func TestFooter(t *testing.T) {
	tests := []struct {
		name     string
		cursor   int
		total    int
		varCount int
		contains []string
		absent   []string
	}{
		{"no variables", 1, 10, 0,
			[]string{" 1/10 ", "↑/↓: nav", "esc: quit"},
			[]string{"active variable count"}},
		{"with variables", 3, 7, 2,
			[]string{" 3/7 ", "active variable count: 2"},
			nil},
		{"zero total", 0, 0, 0,
			[]string{" 0/0 "},
			[]string{"active variable count"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Footer(tc.cursor, tc.total, tc.varCount)
			for _, s := range tc.contains {
				if !strings.Contains(got, s) {
					t.Fatalf("Footer(%d,%d,%d) missing %q: %q", tc.cursor, tc.total, tc.varCount, s, got)
				}
			}
			for _, s := range tc.absent {
				if strings.Contains(got, s) {
					t.Fatalf("Footer(%d,%d,%d) must not contain %q: %q", tc.cursor, tc.total, tc.varCount, s, got)
				}
			}
		})
	}
}
