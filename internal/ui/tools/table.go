// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"strconv"

	"charm.land/bubbles/v2/paginator"
	"charm.land/bubbles/v2/table"
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/tablex"
)

// =============================================================================
// Row Construction
// =============================================================================

// pageRows builds the table rows for the given page.
func pageRows(tools []model.ToolInfo, page, perPage int) []table.Row {
	rows := make([]table.Row, 0, perPage)
	startIdx := page * perPage
	endIdx := startIdx + perPage
	if endIdx > len(tools) {
		endIdx = len(tools)
	}

	for i := startIdx; i < endIdx; i++ {
		rows = append(rows, table.Row{
			tools[i].Name,
			strconv.Itoa(tools[i].CommandCount),
		})
	}
	return rows
}

// =============================================================================
// Widget Construction
// =============================================================================

// newTable creates a styled tools table stretched to the terminal width.
func newTable(rows []table.Row, width, height int) table.Model {
	columns := []table.Column{
		{Title: "Tool", Width: 30},
		{Title: "Commands", Width: 15},
	}
	// The frame keeps a margin around the viewport.
	return tablex.NewWide(columns, rows, height, width-4)
}

// newPaginator creates a dots paginator with the given page count.
func newPaginator(totalPages, page int) paginator.Model {
	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 1
	p.SetTotalPages(totalPages)
	p.Page = page
	p.ActiveDot = lipgloss.NewStyle().Foreground(theme.Accent).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(theme.Dim).Render("•")
	return p
}
