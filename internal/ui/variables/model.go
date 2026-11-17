// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	"sort"

	"charm.land/bubbles/v2/table"

	"github.com/halilkirazkaya/arsenal-ng/internal/ui/tablex"
)

// =============================================================================
// Model Definition
// =============================================================================

// Model holds the state of the show variables view.
type Model struct {
	table table.Model
}

// =============================================================================
// Model Construction
// =============================================================================

// New creates a variables model with a table built from the variables.
func New(variables map[string]string, height int, width ...int) Model {
	return Model{}.Rebuild(variables, height, width...)
}

// Rebuild recreates the table from the current variables.
func (m Model) Rebuild(variables map[string]string, height int, width ...int) Model {
	// Sort keys alphabetically.
	names := make([]string, 0, len(variables))
	for name := range variables {
		names = append(names, name)
	}
	sort.Strings(names)

	// Create rows.
	rows := make([]table.Row, 0, len(names))
	for _, name := range names {
		rows = append(rows, table.Row{name, variables[name]})
	}

	columns := []table.Column{
		{Title: "KEY", Width: 20},
		{Title: "VALUE", Width: 50},
	}
	// Without a width the base columns stand: older callers pass none.
	target := 0
	if len(width) > 0 {
		target = width[0] - 4
	}
	m.table = tablex.NewWide(columns, rows, tablex.Height(len(rows), height, reserved(height)), target)
	return m
}

// =============================================================================
// Model State
// =============================================================================

// Empty reports whether the table has no rows.
func (m Model) Empty() bool {
	return len(m.table.Rows()) == 0
}
