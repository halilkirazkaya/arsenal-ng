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
// View Rendering
// =============================================================================

// View implements tea.Model and delegates rendering to the active view.
func (m App) View() tea.View {
	v := tea.NewView(m.render())
	v.AltScreen = true
	return v
}

// render builds the active view as a plain string.
func (m App) render() string {
	switch m.view {
	case viewSearch:
		return m.search.View(m.globals.All())
	case viewArgs:
		return m.args.View(m.Width())
	case viewVariables:
		return m.variables.View(m.globals.All())
	case viewTools:
		return m.tools.View(m.cheats)
	case viewHelp:
		return help.View()
	}
	return ""
}
