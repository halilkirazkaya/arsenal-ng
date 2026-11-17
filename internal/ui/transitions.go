// Package ui provides the terminal user interface for arsenal-ng.
//
// App implements tea.Model and composes one sub-model per view. Views
// report outcomes; the parent owns cross-view transitions and shared data.
package ui

import (
	"log"

	"github.com/halilkirazkaya/arsenal-ng/internal/ui/tools"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/variables"
)

// =============================================================================
// View Transitions
// =============================================================================

// showVariables switches to the variables view with a cleared query.
func (m App) showVariables() App {
	log.Printf("User navigated to variables view")
	m.view = viewVariables
	m.search = m.search.SetQuery("")
	m.variables = variables.New(m.globals.All(), m.height, m.Width())
	return m
}

// showTools switches to the tools view with a cleared query.
func (m App) showTools() App {
	log.Printf("User navigated to tools view")
	m.view = viewTools
	m.search = m.search.SetQuery("")
	m.tools = tools.New(m.cheats, m.Width(), m.height)
	return m
}

// showHelp switches to the help view with a cleared query.
func (m App) showHelp() App {
	m.view = viewHelp
	m.search = m.search.SetQuery("")
	return m
}

// backToSearch returns to the search view and refocuses the input.
func (m App) backToSearch() App {
	m.view = viewSearch
	m.search = m.search.Focus()
	return m
}
