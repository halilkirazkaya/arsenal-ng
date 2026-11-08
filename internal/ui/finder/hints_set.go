// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"fmt"
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Set Hints
// =============================================================================

// setHints lists the common variables for the bare "set" command.
func setHints() string {
	var b strings.Builder

	b.WriteString(theme.Title.Render("💡 Common variables you can set:"))
	b.WriteString("\n\n")

	for _, v := range CommonHints {
		varName := theme.ArgName.Render(fmt.Sprintf("  set %-10s", v.Name+"="))
		desc := theme.Help.Render(v.Desc)
		b.WriteString(fmt.Sprintf("%s %s\n", varName, desc))
	}

	b.WriteString("\n")
	b.WriteString(theme.Help.Render("  Type the full command and press Enter"))

	return b.String()
}

// setSuggestions lists variables matching the partial "set" input.
func setSuggestions(partial string) string {
	var b strings.Builder

	b.WriteString(theme.Title.Render("💡 Suggestions:"))
	b.WriteString("\n\n")

	found := false
	for _, v := range CommonHints {
		if strings.HasPrefix(v.Name, partial) {
			varName := theme.ArgName.Render(fmt.Sprintf("  set %s=", v.Name))
			desc := theme.Help.Render(v.Desc)
			b.WriteString(fmt.Sprintf("%s %s\n", varName, desc))
			found = true
		}
	}

	// Unknown names are still settable as custom variables.
	if !found {
		b.WriteString(theme.Help.Render(fmt.Sprintf("  set %s=<value>  (custom variable)\n", partial)))
	}

	b.WriteString("\n")
	b.WriteString(theme.Help.Render("  Add '=' and value, then press Enter"))

	return b.String()
}
