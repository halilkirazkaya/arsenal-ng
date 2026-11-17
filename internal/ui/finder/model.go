// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Model State
// =============================================================================

// Model holds the state of the main search view.
type Model struct {
	// Search input component.
	Input textinput.Model

	// Filtered results.
	Filtered []*model.Cheat
	// Cursor position in the filtered list.
	Cursor int
	// First visible row for scrolling.
	Offset int

	// Terminal dimensions for layout calculation.
	width int
	// Terminal height for layout calculation.
	height int

	// Status message for user feedback.
	StatusMsg string
	// Whether the status message is an error.
	StatusIsError bool
}

// =============================================================================
// Model Operations
// =============================================================================

// New creates a search model showing all cheats with no selection.
func New(cheats []*model.Cheat) Model {
	ti := textinput.New()
	ti.Placeholder = config.SearchPlaceholder
	ti.Focus()
	ti.CharLimit = config.SearchCharLimit
	ti.SetWidth(50)
	styles := ti.Styles()
	styles.Focused.Prompt = theme.Prompt
	styles.Focused.Text = lipgloss.NewStyle().Foreground(theme.Bright)
	ti.SetStyles(styles)
	// No item is selected until the user moves the cursor.
	return Model{
		Input:    ti,
		Filtered: cheats,
		Cursor:   -1,
		width:    config.DefaultWidth,
		height:   config.DefaultHeight,
	}
}

// Resize adapts the model to new terminal dimensions.
func (m Model) Resize(width, height int) Model {
	m.width = width
	m.height = height
	m.Input.SetWidth(width - 10)
	return m
}

// Focus focuses the search input.
func (m Model) Focus() Model {
	m.Input.Focus()
	return m
}

// SetQuery replaces the search query.
func (m Model) SetQuery(query string) Model {
	m.Input.SetValue(query)
	return m
}

// ResetList restores the full cheat list with no selection.
func (m Model) ResetList(cheats []*model.Cheat) Model {
	m.Filtered = cheats
	m.Cursor = -1
	m.Offset = 0
	return m
}
