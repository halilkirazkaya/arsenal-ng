// Package ui provides the terminal user interface for arsenal-ng.
//
// App implements tea.Model and composes one sub-model per view. Views
// report outcomes; the parent owns cross-view transitions and shared data.
package ui

import (
	tea "charm.land/bubbletea/v2"

	uivariables "github.com/halilkirazkaya/arsenal-ng/internal/ui/variables"
)

// =============================================================================
// Variables Flow
// =============================================================================

// updateVariables forwards messages to the variables view and applies its outcome.
func (m App) updateVariables(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Lazily rebuild so direct entry always has a table to render.
	if m.variables.Empty() {
		m.variables = m.variables.Rebuild(m.globals.All(), m.height)
	}

	var cmd tea.Cmd
	var outcome uivariables.Outcome

	m.variables, cmd, outcome = m.variables.Update(msg)
	if outcome == uivariables.Back {
		return m.backToSearch(), nil
	}
	return m, cmd
}
