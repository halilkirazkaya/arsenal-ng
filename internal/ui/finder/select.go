// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"log"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Cheat Selection
// =============================================================================

// enter processes the enter key: special commands or cheat selection.
func (m Model) enter(cheats []*model.Cheat) (Model, tea.Cmd, Outcome, Pick) {
	query := strings.TrimSpace(m.Input.Value())
	qLower := strings.ToLower(query)

	switch {
	case qLower == keys.Variables:
		return m, nil, ShowVariables, Pick{}

	case qLower == keys.HelpCmd:
		return m, nil, ShowHelp, Pick{}

	case qLower == keys.Tools:
		return m, nil, ShowTools, Pick{}

	case strings.HasPrefix(qLower, keys.Set+" "):
		return m, nil, SetVar, Pick{}

	case strings.HasPrefix(qLower, keys.Unset+" "):
		return m, nil, UnsetVar, Pick{}
	}

	// Guard against picking with an invalid cursor.
	if len(m.Filtered) == 0 || m.Cursor < 0 || m.Cursor >= len(m.Filtered) {
		log.Printf("Invalid cheat selection: cursor=%d, filtered=%d", m.Cursor, len(m.Filtered))
		return m, nil, Stay, Pick{}
	}

	cheat := m.Filtered[m.Cursor]
	log.Printf("Selected cheat: %s (tool: %s, file: %s)", cheat.Title, cheat.Tool, cheat.Filename)
	args := model.ParseArguments(cheat.Command)

	// Cheats without arguments execute immediately without the args view.
	if len(args) == 0 {
		return m, nil, Picked, Pick{Cheat: cheat, Quit: true}
	}

	return m, nil, Picked, Pick{Cheat: cheat}
}
