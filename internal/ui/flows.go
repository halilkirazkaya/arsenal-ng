// Package ui provides the terminal user interface for arsenal-ng.
//
// App implements tea.Model and composes one sub-model per view. Views
// report outcomes; the parent owns cross-view transitions and shared data.
package ui

import (
	"log"

	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/args"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/tools"
	"github.com/halilkirazkaya/arsenal-ng/internal/variables"
)

// =============================================================================
// Args And Tools Flows
// =============================================================================

// updateArgs forwards messages to the args view and applies its outcome.
func (m App) updateArgs(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var outcome args.Outcome
	var command string

	m.args, cmd, outcome, command = m.args.Update(msg)

	switch outcome {
	case args.Submit:
		if command != "" {
			m.FinalCommand = command
			return m, tea.Quit
		}
		return m, nil

	case args.Back:
		log.Printf("User exited args view, returning to search")
		return m.backToSearch(), nil

	case args.Cancel:
		log.Printf("User cancelled application from args view")
		m.Cancelled = true
		return m, tea.Quit
	}

	return m, cmd
}

// updateTools forwards messages to the tools view and applies its outcome.
func (m App) updateTools(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var outcome tools.Outcome
	var tool string

	m.tools, cmd, outcome, tool = m.tools.Update(msg, m.cheats, m.Width(), m.height)

	switch outcome {
	case tools.Pick:
		if tool == "" {
			return m.backToSearch(), nil
		}
		m.view = viewSearch
		m.search = m.search.SetQuery(tool).Focus().Filter(m.cheats)
		return m, nil

	case tools.Back:
		return m.backToSearch(), nil
	}

	return m, cmd
}

// =============================================================================
// Argument Prefill
// =============================================================================

// countApplied counts cheat args already covered by stored global variables.
func countApplied(cheat *model.Cheat, globals *variables.Store) int {
	applied := 0
	for _, arg := range model.ParseArguments(cheat.Command) {
		if _, ok := globals.Get(arg.Name); ok {
			applied++
		}
	}
	return applied
}
