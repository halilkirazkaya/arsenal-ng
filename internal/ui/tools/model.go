// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"charm.land/bubbles/v2/paginator"
	"charm.land/bubbles/v2/table"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Model Definition
// =============================================================================

// defaultPerPage is the number of tools shown per page.
const defaultPerPage = 25

// Model holds the state of the show tools view.
type Model struct {
	Table     table.Model
	Paginator paginator.Model
	PerPage   int
}

// =============================================================================
// Model Construction
// =============================================================================

// New creates a tools model showing the first page.
func New(cheats []*model.Cheat, width, height int) Model {
	m := Model{PerPage: defaultPerPage}
	return m.Rebuild(cheats, width, height)
}

// Rebuild recreates the table and paginator for the current page.
func (m Model) Rebuild(cheats []*model.Cheat, width, height int) Model {
	tools := model.ToolSummary(cheats)

	if len(tools) == 0 {
		m.Table = newTable(nil, width, 10)
		m.Paginator = newPaginator(1, 0)
		return m
	}

	perPage := m.PerPage
	if perPage <= 0 {
		perPage = defaultPerPage
	}
	totalPages := (len(tools) + perPage - 1) / perPage

	// Preserve current page if paginator already exists.
	currentPage := 0
	if m.Paginator.TotalPages > 0 {
		currentPage = m.Paginator.Page
	}

	// Ensure page is within valid range.
	if currentPage >= totalPages {
		currentPage = totalPages - 1
	}
	if currentPage < 0 {
		currentPage = 0
	}

	m.Paginator = newPaginator(totalPages, currentPage)
	rows := pageRows(tools, currentPage, perPage)
	m.Table = newTable(rows, width, m.tableHeight(len(rows), height))
	return m
}
