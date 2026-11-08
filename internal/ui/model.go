// Package ui provides the terminal user interface for arsenal-ng.
//
// App implements tea.Model and composes one sub-model per view. Views
// report outcomes; the parent owns cross-view transitions and shared data.
package ui

import (
	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/args"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/finder"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/tools"
	uivariables "github.com/halilkirazkaya/arsenal-ng/internal/ui/variables"
	"github.com/halilkirazkaya/arsenal-ng/internal/variables"
)

// =============================================================================
// View State
// =============================================================================

// view identifies the currently active TUI screen.
type view int

const (
	// viewSearch is the main cheat search screen.
	viewSearch view = iota
	// viewArgs collects argument values for the picked cheat.
	viewArgs
	// viewVariables lists the stored global variables.
	viewVariables
	// viewTools lists the available tools for filtering.
	viewTools
	// viewHelp shows the key binding help screen.
	viewHelp
)

// =============================================================================
// Application Model
// =============================================================================

// App is the main TUI application model implementing tea.Model.
type App struct {
	// view is the currently active screen.
	view view
	// cheats holds the loaded cheats shared read-only after load.
	cheats []*model.Cheat
	// globals stores the session-scoped global variables.
	globals *variables.Store
	// width tracks the terminal width in cells.
	width int
	// height tracks the terminal height in cells.
	height int
	// search is the sub-model for the search screen.
	search finder.Model
	// args is the sub-model for the argument input screen.
	args args.Model
	// tools is the sub-model for the tools screen.
	tools tools.Model
	// variables is the sub-model for the variables screen.
	variables uivariables.Model
	// FinalCommand holds the built command after a successful pick.
	FinalCommand string
	// Cancelled reports whether the user aborted the session.
	Cancelled bool
}

// Width returns the usable terminal width clamped to a sane minimum.
func (m App) Width() int {
	return frame.Width(m.width)
}
