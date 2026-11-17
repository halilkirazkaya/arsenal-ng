// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"fmt"
	"slices"
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Unset Hints
// =============================================================================

// unsetHints lists the currently set variables for the "unset" command.
func unsetHints(variables map[string]string) string {
	var b strings.Builder

	if len(variables) == 0 {
		b.WriteString(theme.Title.Render("  ⚠ No variables to unset"))
		b.WriteString("\n\n")
		b.WriteString(theme.Help.Render("  Set variables first:"))
		b.WriteString("\n\n")
		b.WriteString(fmt.Sprintf("    %s\n", theme.ArgName.Render("set key=value")))
		b.WriteString(fmt.Sprintf("    %s\n", theme.ArgName.Render("set ip=10.10.10.10")))
		b.WriteString(fmt.Sprintf("    %s\n", theme.ArgName.Render("set domain=target.com")))
	} else {
		b.WriteString(theme.Title.Render("  💡 Variables you can unset:"))
		b.WriteString("\n\n")
		names := make([]string, 0, len(variables))
		for name := range variables {
			names = append(names, name)
		}
		// Map iteration order is random, so sort for stable output.
		slices.Sort(names)
		for _, name := range names {
			value := variables[name]
			varName := theme.ArgName.Render(fmt.Sprintf("unset %-12s", name))
			val := theme.Help.Render(fmt.Sprintf("(current: %s)", value))
			b.WriteString(fmt.Sprintf("    %s  %s\n", varName, val))
		}
		b.WriteString("\n")
		b.WriteString(theme.Help.Render("  Type variable name and press Enter"))
	}

	return b.String()
}
