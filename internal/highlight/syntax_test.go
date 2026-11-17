// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Tokenizer Tests
// =============================================================================

// TestTokenize verifies token boundaries for flags, pipes, and quotes.
func TestTokenize(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []string
	}{
		{"flag split", "nmap -sV", []string{"nmap", " ", "-sV"}},
		{"pipe split", "a | b", []string{"a", " ", "|", " ", "b"}},
		{"quoted kept whole", `grep "a b"`, []string{"grep", " ", `"a b"`}},
		{"single quoted", `grep 'a b'`, []string{"grep", " ", `'a b'`}},
		{"arg placeholder whole", "nmap {{ip}}", []string{"nmap", " ", "{{ip}}"}},
		{"arg default whole", "ffuf {{word|admin}}", []string{"ffuf", " ", "{{word|admin}}"}},
		{"dollar var kept", "echo $HOME", []string{"echo", " ", "$HOME"}},
		{"semicolon operator", "a; b", []string{"a", ";", " ", "b"}},
		{"redirect", "a > b", []string{"a", " ", ">", " ", "b"}},
		{"double gt single token", "a >> b", []string{"a", " ", ">>", " ", "b"}},
		{"and single token", "a && b", []string{"a", " ", "&&", " ", "b"}},
		{"or single token", "a || b", []string{"a", " ", "||", " ", "b"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tokenize(tc.line)
			if strings.Join(got, "\x00") != strings.Join(tc.want, "\x00") {
				t.Fatalf("tokenize(%q) = %q, want %q", tc.line, got, tc.want)
			}
		})
	}
}

// TestTokenizeJoinIsIdentity verifies tokenizing then joining restores input.
func TestTokenizeJoinIsIdentity(t *testing.T) {
	lines := []string{
		"nmap -sV --top-ports 1000 {{ip}}",
		`grep "a b" x | sort > out; echo 'done'`,
		"curl {{url|http://x}} | jq .",
		"echo $HOME; ls",
		"a >> b && c || d",
	}
	for _, line := range lines {
		if got := strings.Join(tokenize(line), ""); got != line {
			t.Fatalf("join(tokenize(%q)) = %q", line, got)
		}
	}
}

// TestIsPipeOrOperator verifies operator classification of tokens.
func TestIsPipeOrOperator(t *testing.T) {
	for _, op := range []string{"|", ">", "<", ">>", "&&", "||", ";"} {
		if !isPipeOrOperator(op) {
			t.Fatalf("isPipeOrOperator(%q) = false", op)
		}
	}
	for _, not := range []string{"", " ", "nmap", "&", "-"} {
		if isPipeOrOperator(not) {
			t.Fatalf("isPipeOrOperator(%q) = true", not)
		}
	}
}

// =============================================================================
// Style Fixtures
// =============================================================================

// padStyles swaps theme syntax styles for space-padding ones so assertions
// stay observable without ANSI colors, and returns a restore func.
func padStyles() func() {
	old := map[string]lipgloss.Style{
		"tool": theme.CmdTool, "flag": theme.CmdFlag, "def": theme.CmdDefault,
		"str": theme.CmdString, "var": theme.CmdVar,
		"normal": theme.CmdNormal, "pipe": theme.CmdPipe,
	}
	pad := lipgloss.NewStyle().PaddingLeft(1)
	theme.CmdTool, theme.CmdFlag, theme.CmdDefault = pad, pad, pad
	theme.CmdString, theme.CmdVar, theme.CmdNormal, theme.CmdPipe = pad, pad, pad, pad
	return func() {
		theme.CmdTool, theme.CmdFlag, theme.CmdDefault = old["tool"], old["flag"], old["def"]
		theme.CmdString, theme.CmdVar, theme.CmdNormal, theme.CmdPipe = old["str"], old["var"], old["normal"], old["pipe"]
	}
}
