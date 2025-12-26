// Package ui provides the terminal user interface for arsenal-ng.
//
// This file defines the main App struct which implements the tea.Model interface.
// It contains all the state needed for the TUI application including cheat data,
// UI components, navigation state, and view-specific state.
package ui

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/state"
)

// =============================================================================
// Application Model
// =============================================================================

// App is the main TUI application model.
// It implements tea.Model interface.
type App struct {
	// Current application state
	state appState

	// Cheat data
	cheats   []*model.Cheat
	filtered []*model.Cheat

	// Search input component
	searchInput textinput.Model

	// List navigation
	cursor int
	offset int

	// Terminal dimensions
	width  int
	height int

	// Global variables (session-scoped)
	globals *state.Global

	// Status message for user feedback
	statusMsg     string
	statusIsError bool

	// Argument input state
	selectedCheat *model.Cheat
	args          []model.Argument
	argInputs     []textinput.Model
	argCursor     int

	// Result (exported for main to access)
	FinalCommand string
	Cancelled    bool

	// Tools view state
	toolsTable    table.Model
	toolsPaginator paginator.Model
	toolsPerPage  int

	// Variables view state
	varsTable table.Model
}

