// Package ui provides the terminal user interface for arsenal-ng.
//
// App implements tea.Model and composes one sub-model per view. Views
// report outcomes; the parent owns cross-view transitions and shared data.
package ui

import (
	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/ui/help"
)

// =============================================================================
// Message Routing
// =============================================================================

// Update implements tea.Model and routes messages to the active view.
func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Ignore zero-size events so headless resizes never shrink the input.
	if msg, ok := msg.(tea.WindowSizeMsg); ok && msg.Width > 0 && msg.Height > 0 {
		m.width = msg.Width
		m.height = msg.Height
		m.search = m.search.Resize(msg.Width, msg.Height)
		// Tables cache their rows, so rebuild them after a resize.
		if m.view == viewTools {
			m.tools = m.tools.Rebuild(m.cheats, m.Width(), m.height)
		}
		if m.view == viewVariables {
			m.variables = m.variables.Rebuild(m.globals.All(), m.height, m.Width())
		}
	}

	switch m.view {
	case viewSearch:
		return m.updateSearch(msg)
	case viewArgs:
		return m.updateArgs(msg)
	case viewVariables:
		return m.updateVariables(msg)
	case viewTools:
		return m.updateTools(msg)
	case viewHelp:
		if help.Update(msg) == help.Back {
			m = m.backToSearch()
		}
		return m, nil
	}
	return m, nil
}
