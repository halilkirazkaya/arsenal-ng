// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/tablex"
)

// =============================================================================
// View Constants
// =============================================================================

// title is the view heading.
const title = "Available Tools"

// hint is the footer help text.
const hint = "Press ESC to go back │ Enter: search selected tool │ ↑/↓: navigate │ ←/→: change page"

// =============================================================================
// Layout Calculation
// =============================================================================

// tableHeight calculates the table height from row count and terminal size.
func (m Model) tableHeight(rowCount, termHeight int) int {
	header := frame.Header()
	titleView := theme.Title.Render(title)
	paginatorView := m.Paginator.View()
	footer := theme.Help.Render(hint)
	// Sum each section height plus its blank-line spacing.
	reserved := lipgloss.Height(header) + 2 +
		lipgloss.Height(titleView) + 2 +
		lipgloss.Height(paginatorView) + 2 +
		lipgloss.Height(footer) + 2

	return tablex.Height(rowCount, termHeight, reserved)
}
