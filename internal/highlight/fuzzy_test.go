// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

import (
	"testing"

	"charm.land/lipgloss/v2"
)

// =============================================================================
// Fuzzy Tests
// =============================================================================

// Padding-based styles stay observable even when the test env strips ANSI
// colors: every match-styled char gains a leading space.
var (
	plain = lipgloss.NewStyle()
	match = lipgloss.NewStyle().PaddingLeft(1)
)

// TestFuzzy verifies subsequence matching and case-insensitive rendering.
func TestFuzzy(t *testing.T) {
	P, M := plain.Render, match.Render
	tests := []struct {
		name  string
		text  string
		query string
		want  string
	}{
		{"empty query renders plain", "nmap", "", P("nmap")},
		{"full subsequence", "nmap", "nmap", M("n") + M("m") + M("a") + M("p")},
		{"partial subsequence", "nmap", "nm", M("n") + M("m") + P("a") + P("p")},
		{"case insensitive", "Nmap", "nm", M("N") + M("m") + P("a") + P("p")},
		{"no match all plain", "nmap", "xyz", P("n") + P("m") + P("a") + P("p")},
		{"query longer than text", "nm", "nmapx", M("n") + M("m")},
		{"empty text", "", "nm", ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Fuzzy(tc.text, tc.query, plain, match); got != tc.want {
				t.Fatalf("Fuzzy(%q, %q) = %q, want %q", tc.text, tc.query, got, tc.want)
			}
		})
	}
}

// TestFuzzyMultibyte verifies rune-based matching handles multi-byte chars.
func TestFuzzyMultibyte(t *testing.T) {
	// Rune-based matching handles multi-byte chars in the fuzzy loop.
	P, M := plain.Render, match.Render
	got := Fuzzy("café", "é", plain, match)
	want := P("c") + P("a") + P("f") + M("é")
	if got != want {
		t.Fatalf("Fuzzy multibyte = %q, want %q", got, want)
	}
}

// TestFuzzyOutOfOrderNoMatch verifies greedy matching only moves forward.
func TestFuzzyOutOfOrderNoMatch(t *testing.T) {
	// "ab" in "ba": 'b' is plain, 'a' matches the first query rune.
	P, M := plain.Render, match.Render
	got := Fuzzy("ba", "ab", plain, match)
	want := P("b") + M("a")
	if got != want {
		t.Fatalf("Fuzzy out-of-order = %q, want %q", got, want)
	}
}
