// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	"fmt"
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Variables Rendering
// =============================================================================

// View renders the variables table, or usage help when empty.
func (m Model) View(variables map[string]string) string {
	var b strings.Builder

	b.WriteString(frame.Header())
	b.WriteString("\n\n")

	b.WriteString(theme.Title.Render(title))
	b.WriteString("\n\n")

	if len(variables) == 0 {
		b.WriteString(theme.Help.Render("No variables set."))
		b.WriteString("\n\n")
		b.WriteString(theme.Title.Render("Usage:"))
		b.WriteString("\n\n")
		b.WriteString(fmt.Sprintf("  %s  %s\n", theme.ArgName.Render("set key=value     "), theme.Help.Render("Set a variable")))
		b.WriteString(fmt.Sprintf("  %s  %s\n", theme.ArgName.Render("unset key         "), theme.Help.Render("Remove a variable")))
		b.WriteString(fmt.Sprintf("  %s  %s\n", theme.ArgName.Render("variables         "), theme.Help.Render("List all variables")))
	} else {
		b.WriteString(m.table.View())
	}

	b.WriteString("\n\n")
	b.WriteString(theme.Help.Render(hint))

	return b.String()
}
