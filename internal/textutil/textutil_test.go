// Package textutil provides plain string utilities for terminal output:
// truncation with ellipsis and word-boundary wrapping.
package textutil

import (
	"strings"
	"testing"
)

// =============================================================================
// Truncation Cases
// =============================================================================

// TestTruncate verifies shortening, ellipsis, and boundary widths.
func TestTruncate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		max  int
		want string
	}{
		{"short untouched", "nmap", 10, "nmap"},
		{"exact fit", "nmap", 4, "nmap"},
		{"ellipsis", "nmap --top-ports 1000", 10, "nmap --..."},
		{"newline flattened", "a\nb", 10, "a b"},
		{"newline flattened then cut", "aa\nbbbbbbbb", 5, "aa..."},
		{"empty input", "", 5, ""},
		{"zero max", "nmap", 0, ""},
		{"negative max", "nmap", -1, ""},
		{"tiny max raw slice", "nmap", 2, "nm"},
		{"tiny max shorter input", "n", 3, "n"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Truncate(tc.s, tc.max); got != tc.want {
				t.Fatalf("Truncate(%q, %d) = %q, want %q", tc.s, tc.max, got, tc.want)
			}
		})
	}
}

// =============================================================================
// Wrapping Cases
// =============================================================================

// TestWrap verifies word-boundary wrapping and passthrough widths.
func TestWrap(t *testing.T) {
	tests := []struct {
		name  string
		text  string
		width int
		want  string
	}{
		{"fits one line", "nmap scan", 20, "nmap scan"},
		{"word boundary", "aa bb cc", 4, "aa\nbb\ncc"},
		{"newline to space", "aa\nbb", 20, "aa bb"},
		{"empty input", "", 10, ""},
		{"nonpositive width passthrough", "a b", 0, "a b"},
		{"negative width passthrough", "a b", -3, "a b"},
		{"long word kept whole", "superlongword x", 5, "superlongword\nx"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Wrap(tc.text, tc.width); got != tc.want {
				t.Fatalf("Wrap(%q, %d) = %q, want %q", tc.text, tc.width, got, tc.want)
			}
		})
	}
}

// TestWrapPreservesAllWords verifies wrapping never drops input words.
func TestWrapPreservesAllWords(t *testing.T) {
	got := Wrap("the quick brown fox jumps", 10)
	for _, w := range []string{"the", "quick", "brown", "fox", "jumps"} {
		found := false
		for _, line := range strings.Split(got, "\n") {
			for _, f := range strings.Fields(line) {
				if f == w {
					found = true
					break
				}
			}
		}
		if !found {
			t.Fatalf("Wrap dropped word %q: %q", w, got)
		}
	}
}
