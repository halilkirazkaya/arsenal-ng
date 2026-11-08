// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Layout Metrics
// =============================================================================

// minVisible is the minimum number of visible list items.
const minVisible = 3

// MaxVisible calculates available space for list items dynamically.
func (m Model) MaxVisible() int {
	fixed := m.headerHeight() + m.infoBoxHeight() + m.footerHeight()
	available := m.height - fixed

	if available < minVisible {
		return minVisible
	}
	return available
}

// headerHeight returns the height of the header (logo plus spacing).
func (m Model) headerHeight() int {
	header := frame.Header()
	return lipgloss.Height(header) + 2
}

// infoBoxHeight returns the height of the info box.
func (m Model) infoBoxHeight() int {
	if len(m.Filtered) == 0 || m.Cursor < 0 || m.Cursor >= len(m.Filtered) {
		return 0
	}
	infoBox := InfoBox(m.Filtered[m.Cursor], m.Width())
	return lipgloss.Height(infoBox)
}

// footerHeight returns the height of footer elements.
func (m Model) footerHeight() int {
	footer := m.Footer(0)
	searchInput := theme.Prompt.Render("❯ ") + m.Input.View()
	return lipgloss.Height(footer) + lipgloss.Height(searchInput) + 2
}

// Width returns the usable width for the search view.
func (m Model) Width() int {
	return frame.Width(m.width)
}
