// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Search Rendering
// =============================================================================

// View renders the search view. Variables is needed for unset hints.
func (m Model) View(variables map[string]string) string {
	var b strings.Builder
	width := m.Width()

	// Header renders the logo banner.
	b.WriteString(frame.Header())
	b.WriteString("\n\n")

	// Status lines surface feedback above the list.
	if m.StatusMsg != "" {
		style := theme.Success
		if m.StatusIsError {
			style = theme.Error
		}
		b.WriteString(style.Render(m.StatusMsg))
		b.WriteString("\n\n")
	}

	// The info box previews the highlighted cheat.
	if len(m.Filtered) > 0 && m.Cursor >= 0 && m.Cursor < len(m.Filtered) {
		b.WriteString(InfoBox(m.Filtered[m.Cursor], width))
		b.WriteString("\n\n")
	}

	// Search prompt and input come before the results.
	b.WriteString(theme.Prompt.Render("❯ "))
	b.WriteString(m.Input.View())
	b.WriteString("\n\n")

	// Special commands show hints instead of the cheat list.
	query := m.Input.Value()
	if !IsSpecialCommand(query) {
		b.WriteString(m.cheatList(query))
	} else {
		b.WriteString(Hints(query, variables))
	}

	// Footer reports position and help.
	b.WriteString("\n")
	b.WriteString(m.Footer(len(variables)))

	return b.String()
}

// Footer renders the search view footer with status and help text.
func (m Model) Footer(varCount int) string {
	cursorPos := m.Cursor + 1
	if m.Cursor < 0 {
		cursorPos = 0
	}
	return frame.Footer(cursorPos, len(m.Filtered), varCount)
}
