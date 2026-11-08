// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Command Hints
// =============================================================================

// Hints renders guidance for the special command being typed.
func Hints(query string, variables map[string]string) string {
	var b strings.Builder
	q := strings.ToLower(strings.TrimSpace(query))

	switch {
	case q == keys.Set || q == keys.Set+" ":
		b.WriteString(setHints())

	case strings.HasPrefix(q, keys.Set+" ") && !strings.Contains(q, "="):
		partial := strings.TrimPrefix(q, keys.Set+" ")
		b.WriteString(setSuggestions(partial))

	case strings.HasPrefix(q, keys.Set+" ") && strings.Contains(q, "="):
		b.WriteString(theme.Success.Render("✓ Press Enter to set variable"))

	case q == keys.Unset || q == keys.Unset+" ":
		b.WriteString(unsetHints(variables))

	case q == keys.Variables || q == keys.Variables+" ":
		b.WriteString(theme.Success.Render("✓ Press Enter to show all variables"))

	case q == keys.Tools:
		b.WriteString(theme.Success.Render("✓ Press Enter to show all available tools"))

	case q == keys.HelpCmd || q == keys.HelpCmd+" ":
		b.WriteString(theme.Success.Render("✓ Press Enter to show help menu"))

	default:
		b.WriteString(theme.Help.Render("Press Enter to execute command..."))
	}

	return b.String()
}
