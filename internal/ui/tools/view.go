// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"fmt"
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Tools Rendering
// =============================================================================

// View renders the paginated tools table. Pure: no model mutation.
func (m Model) View(cheats []*model.Cheat) string {
	var b strings.Builder

	b.WriteString(frame.Header())
	b.WriteString("\n\n")

	b.WriteString(theme.Title.Render(title))
	b.WriteString("\n\n")

	tools := model.ToolSummary(cheats)
	if len(tools) == 0 {
		b.WriteString(theme.Help.Render("No tools found."))
	} else {
		b.WriteString(m.Table.View())
		b.WriteString("\n\n")

		b.WriteString("  ")
		b.WriteString(m.Paginator.View())
		b.WriteString("\n\n")

		b.WriteString(theme.Help.Render(fmt.Sprintf("Total: %d tools", len(tools))))
	}

	b.WriteString("\n\n")
	b.WriteString(theme.Help.Render(hint))

	return b.String()
}
