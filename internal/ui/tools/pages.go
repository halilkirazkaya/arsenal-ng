// Package tools provides the tools view: a paginated table of available
// tools with their command counts.
package tools

import (
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Page Navigation
// =============================================================================

// PrevPage moves to the previous tools page.
func (m Model) PrevPage(cheats []*model.Cheat, width, height int) Model {
	if m.Paginator.Page > 0 {
		m.Paginator.PrevPage()
		m = m.Rebuild(cheats, width, height)
	}
	return m
}

// NextPage moves to the next tools page.
func (m Model) NextPage(cheats []*model.Cheat, width, height int) Model {
	if m.Paginator.Page < m.Paginator.TotalPages-1 {
		m.Paginator.NextPage()
		m = m.Rebuild(cheats, width, height)
	}
	return m
}

// =============================================================================
// Current Selection
// =============================================================================

// Selected returns the highlighted tool name, or empty when none.
func (m Model) Selected() string {
	row := m.Table.SelectedRow()
	if len(row) == 0 {
		return ""
	}
	return row[0]
}
