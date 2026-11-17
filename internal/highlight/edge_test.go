// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
)

// =============================================================================
// Edge Case Tests
// =============================================================================

// pad renders with left padding so styled spans stay observable.
func pad(s string) string {
	return lipgloss.NewStyle().PaddingLeft(1).Render(s)
}

// TestSyntaxTruncatedArg verifies truncated placeholders take the var style.
func TestSyntaxTruncatedArg(t *testing.T) {
	defer padStyles()()
	// Truncated "{{arg" and "{{arg|def" take the truncated-arg branch.
	if got, want := Syntax(`echo {{ip`), pad("echo")+" "+pad("{{ip"); got != want {
		t.Fatalf("truncated arg = %q, want %q", got, want)
	}
	if got, want := Syntax(`echo {{w|x`), pad("echo")+" "+pad("{{w|x"); got != want {
		t.Fatalf("truncated default = %q, want %q", got, want)
	}
}

// TestSyntaxNormalToken verifies plain words take the default branch.
func TestSyntaxNormalToken(t *testing.T) {
	defer padStyles()()
	// Plain words after the tool take the default (normal) branch.
	if got, want := Syntax(`nmap target`), pad("nmap")+" "+pad("target"); got != want {
		t.Fatalf("normal token = %q, want %q", got, want)
	}
}

// TestTokenizeUnclosedQuote verifies unclosed quotes flush at end of line.
func TestTokenizeUnclosedQuote(t *testing.T) {
	// Unclosed quotes flush the buffer at end of line.
	got := tokenize(`grep "abc`)
	want := []string{"grep", " ", `"abc`}
	if strings.Join(got, "\x00") != strings.Join(want, "\x00") {
		t.Fatalf("tokenize = %q, want %q", got, want)
	}
}

// TestTokenizeUnclosedBraces verifies unclosed placeholders flush at end.
func TestTokenizeUnclosedBraces(t *testing.T) {
	// Unclosed {{ }} flushes the buffer at end of line.
	got := tokenize(`echo {{ip`)
	want := []string{"echo", " ", "{{ip"}
	if strings.Join(got, "\x00") != strings.Join(want, "\x00") {
		t.Fatalf("tokenize = %q, want %q", got, want)
	}
}

// TestTokenizeEscapedQuote verifies escaped quotes do not close the span.
func TestTokenizeEscapedQuote(t *testing.T) {
	// Backslash-escaped quotes don't close the quoted span.
	got := tokenize(`echo "a\"b" x`)
	want := []string{"echo", " ", `"a\"b"`, " ", "x"}
	if strings.Join(got, "\x00") != strings.Join(want, "\x00") {
		t.Fatalf("tokenize = %q, want %q", got, want)
	}
}

// TestTokenizeFlushBeforeDelimiters verifies text flushes before delimiters.
func TestTokenizeFlushBeforeDelimiters(t *testing.T) {
	// Text accumulated before {{ or a quote is flushed as its own token.
	got := tokenize(`pre{{x}}`)
	want := []string{"pre", "{{x}}"}
	if strings.Join(got, "\x00") != strings.Join(want, "\x00") {
		t.Fatalf("tokenize = %q, want %q", got, want)
	}
	got = tokenize(`a"b c"`)
	want = []string{"a", `"b c"`}
	if strings.Join(got, "\x00") != strings.Join(want, "\x00") {
		t.Fatalf("tokenize = %q, want %q", got, want)
	}
}

// TestTokenizeTrailingOperator verifies a trailing operator has no lookahead.
func TestTokenizeTrailingOperator(t *testing.T) {
	// A line ending in an operator has no lookahead char.
	got := tokenize(`echo hi |`)
	want := []string{"echo", " ", "hi", " ", "|"}
	if strings.Join(got, "\x00") != strings.Join(want, "\x00") {
		t.Fatalf("tokenize = %q, want %q", got, want)
	}
}
