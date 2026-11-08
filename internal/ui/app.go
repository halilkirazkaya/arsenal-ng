// Package ui provides the terminal user interface for arsenal-ng.
//
// App implements tea.Model and composes one sub-model per view. Views
// report outcomes; the parent owns cross-view transitions and shared data.
package ui

import (
	"log"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/finder"
	"github.com/halilkirazkaya/arsenal-ng/internal/variables"
)

// =============================================================================
// Application Construction
// =============================================================================

// New creates a TUI application starting on the search view.
func New(cheats []*model.Cheat) App {
	// Fall back to an empty store so a broken config dir never blocks startup.
	globals, err := variables.New()
	if err != nil {
		log.Printf("WARNING: Failed to load global variables, continuing with empty store: %v", err)
		globals = &variables.Store{}
	}

	return App{
		view:    viewSearch,
		cheats:  cheats,
		globals: globals,
		width:   config.DefaultWidth,
		height:  config.DefaultHeight,
		search:  finder.New(cheats),
	}
}

// Init implements tea.Model and returns the initial cursor blink command.
func (m App) Init() tea.Cmd {
	return textinput.Blink
}
