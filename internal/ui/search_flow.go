// Package ui provides the terminal user interface for arsenal-ng.
//
// App implements tea.Model and composes one sub-model per view. Views
// report outcomes; the parent owns cross-view transitions and shared data.
package ui

import (
	"log"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/args"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/finder"
)

// =============================================================================
// Search Flow
// =============================================================================

// updateSearch forwards messages to the finder and applies its outcome.
func (m App) updateSearch(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var outcome finder.Outcome
	var pick finder.Pick

	m.search, cmd, outcome, pick = m.search.Update(msg, m.cheats)

	switch outcome {
	case finder.Picked:
		return m.pickCheat(pick)

	case finder.ShowVariables:
		return m.showVariables(), nil

	case finder.ShowTools:
		return m.showTools(), nil

	case finder.ShowHelp:
		log.Printf("User navigated to help view")
		return m.showHelp(), nil

	case finder.SetVar:
		m = m.applySet(strings.TrimSpace(m.search.Input.Value()))
		return m, nil

	case finder.UnsetVar:
		m = m.applyUnset(strings.TrimSpace(m.search.Input.Value()))
		return m, nil

	case finder.Cancel:
		if key, ok := msg.(tea.KeyPressMsg); ok {
			log.Printf("User cancelled application (key: %s)", key.String())
		}
		m.Cancelled = true
		return m, tea.Quit
	}

	return m, cmd
}

// pickCheat enters args mode for the pick or quits when it needs no args.
func (m App) pickCheat(pick finder.Pick) (tea.Model, tea.Cmd) {
	if pick.Quit {
		m.FinalCommand = model.BuildCommand(pick.Cheat.Command, nil)
		log.Printf("Command ready (no arguments): %s", m.FinalCommand)
		return m, tea.Quit
	}

	applied := countApplied(pick.Cheat, m.globals)
	log.Printf("Entering argument input mode (%d argument(s), %d pre-filled with variables)",
		len(model.ParseArguments(pick.Cheat.Command)), applied)
	m.view = viewArgs
	m.args = args.New(pick.Cheat, m.globals)
	return m, nil
}
