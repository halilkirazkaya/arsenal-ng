// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the variables view rendering functions. It displays all
// global variables in a table format and provides initialization for the
// variables table component.
package ui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Show Variables View
// =============================================================================

func (m App) viewShowVars() string {
	var b strings.Builder

	b.WriteString(m.renderHeader())
	b.WriteString("\n\n")

	b.WriteString(titleStyle.Render("Global Variables"))
	b.WriteString("\n\n")

	vars := m.globals.All()
	if len(vars) == 0 {
		b.WriteString(helpStyle.Render("No variables set."))
		b.WriteString("\n\n")
		b.WriteString(titleStyle.Render("Usage:"))
		b.WriteString("\n\n")
		b.WriteString(fmt.Sprintf("  %s  %s\n", argNameStyle.Render("set key=value     "), helpStyle.Render("Set a variable")))
		b.WriteString(fmt.Sprintf("  %s  %s\n", argNameStyle.Render("unset key         "), helpStyle.Render("Remove a variable")))
		b.WriteString(fmt.Sprintf("  %s  %s\n", argNameStyle.Render("variables         "), helpStyle.Render("List all variables")))
	} else {
		b.WriteString(m.varsTable.View())
	}

	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("Press ESC or Enter to go back │ ↑/↓: navigate"))

	return b.String()
}

// initVarsTable initializes the variables table with variable data.
func (m App) initVarsTable() App {
	vars := m.globals.All()

	// Sort keys alphabetically
	keys := make([]string, 0, len(vars))
	for name := range vars {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	// Create columns
	columns := []table.Column{
		{Title: "KEY", Width: 20},
		{Title: "VALUE", Width: 50},
	}

	// Create rows
	rows := make([]table.Row, 0, len(keys))
	for _, key := range keys {
		rows = append(rows, table.Row{
			key,
			vars[key],
		})
	}

	// Calculate table height based on actual rows and terminal size
	tableHeight := len(rows) + 2 // +2 for header

	// Calculate reserved space dynamically
	header := m.renderHeader()
	title := titleStyle.Render("Global Variables")
	footer := helpStyle.Render("Press ESC or Enter to go back │ ↑/↓: navigate")
	reservedHeight := lipgloss.Height(header) + 2 + // header + spacing
		lipgloss.Height(title) + 2 + // title + spacing
		lipgloss.Height(footer) + 2 // footer + spacing

	maxHeight := m.height - reservedHeight
	if tableHeight > maxHeight {
		tableHeight = maxHeight
	}
	if tableHeight < 3 {
		tableHeight = 3 // Minimum height
	}
	if len(rows) == 0 {
		tableHeight = 3 // Empty table
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(tableHeight),
	)

	// Style the table
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(config.ColorSecondary)).
		Foreground(lipgloss.Color(config.ColorBright)).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#4A4A4A")).
		Bold(true)
	t.SetStyles(s)

	m.varsTable = t
	return m
}

