// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

import (
	"testing"

	"charm.land/lipgloss/v2"
)

// =============================================================================
// Syntax Render Tests
// =============================================================================

// TestSyntax verifies end-to-end rendering of tools, flags, and args.
func TestSyntax(t *testing.T) {
	defer padStyles()()
	p := lipgloss.NewStyle().PaddingLeft(1).Render
	tests := []struct {
		name string
		cmd  string
		want string
	}{
		{"tool then flag", "nmap -sV", p("nmap") + " " + p("-sV")},
		{"pipe resets tool", "a | grep", p("a") + " " + p("|") + " " + p("grep")},
		{"arg no default", "nmap {{ip}}", p("nmap") + " " + p("{{ip}}")},
		{"arg with default", "ffuf {{w|admin}}", p("ffuf") + " " + p("{{w|admin}}")},
		{"dollar var", "echo $HOME", p("echo") + " " + p("$HOME")},
		{"quoted string", `grep "a b"`, p("grep") + " " + p(`"a b"`)},
		{"empty line passthrough", "", ""},
		{"blank line passthrough", "   ", "   "},
		{"multiline joined", "a\nb", p("a") + "\n" + p("b")},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Syntax(tc.cmd); got != tc.want {
				t.Fatalf("Syntax(%q) = %q, want %q", tc.cmd, got, tc.want)
			}
		})
	}
}

// TestSyntaxQuotedWithArg verifies args inside quotes route to arg styling.
func TestSyntaxQuotedWithArg(t *testing.T) {
	defer padStyles()()
	got := Syntax(`echo "{{h}}"`)
	// ArgPattern branch wins over the quoted branch: quotes pass through
	// unstyled, only the {{arg}} span is rendered.
	want := " echo" + " " + `"` + " {{h}}" + `"`
	if got != want {
		t.Fatalf("quoted arg Syntax = %q, want %q", got, want)
	}
}
