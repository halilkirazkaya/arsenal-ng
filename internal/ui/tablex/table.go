// Package tablex builds the styled tables shared by the tools and
// variables views: columns, rows, focus, height, and selection styles.
package tablex

import (
	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Table Construction
// =============================================================================

// New creates a focused table with the shared selection styles.
func New(columns []table.Column, rows []table.Row, height int) table.Model {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(height),
	)
	t.SetStyles(Styles())
	// v2 viewports render no rows at zero width: size from columns.
	t.SetWidth(Width(columns))
	return t
}

// Styles returns the shared table styles for tools and variables.
func Styles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(theme.Secondary).
		Foreground(theme.Bright).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#4A4A4A")).
		Bold(true)
	return s
}

// =============================================================================
// Size Calculation
// =============================================================================

// Width sums column widths plus inter-column spacing.
func Width(columns []table.Column) int {
	width := 0
	for _, col := range columns {
		width += col.Width + 2
	}
	return width
}

// Height fits rows plus header into the terminal clamped to a minimum.
func Height(rowCount, termHeight, reserved int) int {
	// Add space for the header rows above the data.
	height := rowCount + 2

	maxHeight := termHeight - reserved
	if height > maxHeight {
		height = maxHeight
	}
	if height < 3 {
		height = 3
	}
	return height
}
