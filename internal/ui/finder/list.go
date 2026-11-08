// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/highlight"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/textutil"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Cheat List
// =============================================================================

// InfoBox renders the detail box for the selected cheat.
func InfoBox(cheat *model.Cheat, width int) string {
	var content strings.Builder

	content.WriteString(theme.Title.Render(cheat.Title))
	if cheat.Desc != "" {
		content.WriteString("\n")
		// Descriptions wrap so long text stays inside the box.
		wrappedDesc := textutil.Wrap(cheat.Desc, width-6)
		content.WriteString(theme.Desc.Render(wrappedDesc))
	}
	content.WriteString("\n")
	content.WriteString(highlight.Syntax(textutil.Truncate(cheat.Command, width-6)))

	return theme.InfoBox.Width(width - 4).Render(content.String())
}

// cheatList renders the visible slice of filtered cheats.
func (m Model) cheatList(query string) string {
	var b strings.Builder

	maxVisible := m.MaxVisible()
	end := m.Offset + maxVisible
	if end > len(m.Filtered) {
		end = len(m.Filtered)
	}

	matchStyle := lipgloss.NewStyle().
		Foreground(theme.Accent).
		Bold(true)

	for i := m.Offset; i < end; i++ {
		b.WriteString(m.cheatRow(i, m.Filtered[i], query, matchStyle))
	}

	return b.String()
}

// cheatRow renders a single cheat row with cursor, tool, title, and tags.
func (m Model) cheatRow(index int, cheat *model.Cheat, query string, matchStyle lipgloss.Style) string {
	// The highlighted row renders as one band so the whole line is covered.
	if index == m.Cursor {
		return m.selectedRow(cheat)
	}
	// Blank space keeps unselected rows aligned with the cursor marker.
	cursor := "  "
	// Tool names align in a fixed-width column.
	toolText := textutil.Truncate(cheat.Tool, 12)
	tool := highlight.Fuzzy(fmt.Sprintf("%-12s", toolText), query, theme.Tool, matchStyle)
	// Titles align in a fixed-width column with fuzzy match highlighting.
	titleText := textutil.Truncate(cheat.Title, 40)
	title := highlight.Fuzzy(fmt.Sprintf("%-40s", titleText), query, lipgloss.NewStyle(), matchStyle)
	// Tags render after the title.
	tags := theme.ColoredTags(cheat.Tags)
	return fmt.Sprintf("%s%s %s %s\n", cursor, tool, title, tags)
}

// selectedRow renders the highlighted row as a single full-width band.
func (m Model) selectedRow(cheat *model.Cheat) string {
	// Plain fields stay readable on the selection background.
	tool := fmt.Sprintf("%-12s", textutil.Truncate(cheat.Tool, 12))
	title := fmt.Sprintf("%-40s", textutil.Truncate(cheat.Title, 40))
	line := fmt.Sprintf("▸ %s %s %s", tool, title, strings.Join(cheat.Tags, ", "))
	// Width pads the band to the view so resizing keeps it full-bleed.
	return theme.Selected.Width(m.Width()).Render(line) + "\n"
}
