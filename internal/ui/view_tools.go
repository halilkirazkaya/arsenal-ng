// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the tools view rendering functions. It displays all
// available tools with their command counts in a paginated table format
// and provides initialization for the tools table component.
package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Show Tools View
// =============================================================================

func (m App) viewShowTools() string {
	var b strings.Builder
	width := m.effectiveWidth()

	b.WriteString(m.renderHeader())
	b.WriteString("\n\n")

	b.WriteString(titleStyle.Render("Available Tools"))
	b.WriteString("\n\n")

	toolInfoList := m.getToolInfoList()
	if len(toolInfoList) == 0 {
		b.WriteString(helpStyle.Render("No tools found."))
	} else {
		// Render table
		tableWidth := width - 4
		if tableWidth > 0 {
			m.toolsTable.SetWidth(tableWidth)
		}
		b.WriteString(m.toolsTable.View())
		b.WriteString("\n\n")

		// Render paginator
		b.WriteString("  ")
		b.WriteString(m.toolsPaginator.View())
		b.WriteString("\n\n")

		// Total count
		b.WriteString(helpStyle.Render(fmt.Sprintf("Total: %d tools", len(toolInfoList))))
	}

	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("Press ESC to go back │ Enter: search selected tool │ ↑/↓: navigate │ ←/→: change page"))

	return b.String()
}

// initToolsTable initializes the tools table with tool data.
func (m App) initToolsTable() App {
	toolInfoList := m.getToolInfoList()

	if len(toolInfoList) == 0 {
		// Empty table
		columns := []table.Column{
			{Title: "Tool", Width: 30},
			{Title: "Commands", Width: 15},
		}
		t := table.New(
			table.WithColumns(columns),
			table.WithRows([]table.Row{}),
			table.WithFocused(true),
			table.WithHeight(10),
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

		m.toolsTable = t
		m.toolsPaginator = paginator.New()
		m.toolsPaginator.Type = paginator.Dots
		m.toolsPaginator.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorAccent)).Render("•")
		m.toolsPaginator.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorDim)).Render("•")
		return m
	}

	// Initialize or update paginator first
	totalPages := (len(toolInfoList) + m.toolsPerPage - 1) / m.toolsPerPage
	if totalPages == 0 {
		totalPages = 1
	}

	// Preserve current page if paginator already exists
	currentPage := 0
	if m.toolsPaginator.TotalPages > 0 {
		currentPage = m.toolsPaginator.Page
	}

	// Ensure page is within valid range
	if currentPage >= totalPages {
		currentPage = totalPages - 1
	}
	if currentPage < 0 {
		currentPage = 0
	}

	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 1
	p.SetTotalPages(totalPages)
	p.Page = currentPage
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorAccent)).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorDim)).Render("•")
	m.toolsPaginator = p

	// Create columns
	columns := []table.Column{
		{Title: "Tool", Width: 30},
		{Title: "Commands", Width: 15},
	}

	// Create rows for current page
	rows := make([]table.Row, 0, m.toolsPerPage)
	startIdx := currentPage * m.toolsPerPage
	endIdx := startIdx + m.toolsPerPage
	if endIdx > len(toolInfoList) {
		endIdx = len(toolInfoList)
	}

	for i := startIdx; i < endIdx; i++ {
		rows = append(rows, table.Row{
			toolInfoList[i].Name,
			strconv.Itoa(toolInfoList[i].CommandCount),
		})
	}

	// Calculate table height based on actual rows and terminal size
	tableHeight := len(rows) + 2 // +2 for header

	// Calculate reserved space dynamically
	header := m.renderHeader()
	title := titleStyle.Render("Available Tools")
	paginatorView := m.toolsPaginator.View()
	footer := helpStyle.Render("Press ESC or Enter to go back │ ↑/↓: navigate │ ←/→: change page")
	reservedHeight := lipgloss.Height(header) + 2 + // header + spacing
		lipgloss.Height(title) + 2 + // title + spacing
		lipgloss.Height(paginatorView) + 2 + // paginator + spacing
		lipgloss.Height(footer) + 2 // footer + spacing

	maxHeight := m.height - reservedHeight
	if tableHeight > maxHeight {
		tableHeight = maxHeight
	}
	if tableHeight < 3 {
		tableHeight = 3 // Minimum height
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

	m.toolsTable = t
	return m
}

