// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
	"github.com/halilkirazkaya/arsenal-ng/internal/highlight"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Args Rendering
// =============================================================================

// View renders the argument input form with a command preview.
func (m Model) View(width int) string {
	var b strings.Builder
	width = frame.Width(width)

	// Preview shows the command with current values filled in.
	preview := model.BuildCommand(m.Cheat.Command, m.Args)
	previewBox := theme.InfoBox.Width(width - 4).Render(
		theme.Title.Render("Command Preview") + "\n" + highlight.Syntax(preview),
	)
	b.WriteString(previewBox)
	b.WriteString("\n\n")

	// Form lists one row per argument.
	b.WriteString(theme.Title.Render("Arguments:"))
	b.WriteString("\n\n")

	for i, arg := range m.Args {
		b.WriteString(m.row(i, arg))
	}

	// Footer documents the form keys.
	b.WriteString("\n")
	b.WriteString(theme.Help.Render("tab/↓: next │ shift+tab/↑: prev │ ctrl+n/p: cycle │ enter: confirm │ esc: back"))

	return b.String()
}

// row renders a single argument row with cursor, name, and input.
func (m Model) row(index int, arg model.Argument) string {
	// The cursor marker distinguishes the focused row.
	cursor := "  "
	if index == m.Cursor {
		cursor = theme.Cursor.Render("▸ ")
	}

	name := theme.ArgName.Render(fmt.Sprintf("%-15s", arg.Name))
	input := m.Inputs[index].View()

	return fmt.Sprintf("%s%s = %s\n", cursor, name, input) + m.matches(index)
}

// matchPage is how many path completions show under the focused field.
const matchPage = 8

// matches lists the focused field's path completions when more than one
// fits, paged around the one Ctrl+N/Ctrl+P has selected.
func (m Model) matches(index int) string {
	if index != m.Cursor {
		return ""
	}
	in := m.Inputs[index]
	all := in.MatchedSuggestions()
	if len(all) < 2 {
		return ""
	}

	dir, _ := filepath.Split(in.Value())
	cur := in.CurrentSuggestionIndex()
	start := cur - cur%matchPage
	end := min(start+matchPage, len(all))

	var b strings.Builder
	pad := strings.Repeat(" ", 22)
	for i := start; i < end; i++ {
		name := strings.TrimPrefix(all[i], dir)
		if i == cur {
			b.WriteString(pad + theme.Cursor.Render("▸ "+name) + "\n")
		} else {
			b.WriteString(pad + theme.Help.Render("  "+name) + "\n")
		}
	}
	if len(all) > matchPage {
		b.WriteString(pad + theme.Help.Render(fmt.Sprintf("  %d-%d of %d", start+1, end, len(all))) + "\n")
	}
	return b.String()
}
