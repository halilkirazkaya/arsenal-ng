// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the application initialization functions: New() creates
// a new App instance with default settings, and Init() implements the tea.Model
// interface to return the initial command for the TUI.
package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/state"
)

// =============================================================================
// Application Initialization
// =============================================================================

// New creates a new TUI application with the given cheats.
func New(cheats []*model.Cheat) App {
	ti := textinput.New()
	ti.Placeholder = config.SearchPlaceholder
	ti.Focus()
	ti.CharLimit = config.SearchCharLimit
	ti.Width = 50
	ti.PromptStyle = promptStyle
	ti.TextStyle = lipgloss.NewStyle().Foreground(brightColor)

	// Load persistent variables
	globals, err := state.NewGlobal()
	if err != nil {
		// If can't load variables, continue with empty store!
		// This allows the app to work even if config dir can't be created
		log.Printf("WARNING: Failed to load global variables, continuing with empty store: %v", err)
		globals = &state.Global{}
	}

	return App{
		state:        stateSearch,
		cheats:       cheats,
		filtered:     cheats,
		searchInput:  ti,
		cursor:       -1, // No item selected initially
		width:        config.DefaultWidth,
		height:       config.DefaultHeight,
		globals:      globals,
		toolsPerPage: 25,
	}
}

// Init implements tea.Model. Returns initial command (cursor blink).
func (m App) Init() tea.Cmd {
	return textinput.Blink
}

