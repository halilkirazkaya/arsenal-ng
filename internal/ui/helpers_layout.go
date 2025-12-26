// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains layout calculation functions that compute dimensions
// for UI components based on terminal size. It includes functions for
// calculating header height, info box height, footer height, and available
// space for list items.
package ui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Layout Constants
// =============================================================================

const (
	minVisibleItems = 3
	minWidth        = 40
)

// =============================================================================
// Layout Calculations
// =============================================================================

// headerHeight returns the height of the header (logo + spacing).
func (m App) headerHeight() int {
	header := m.renderHeader()
	return lipgloss.Height(header) + 2
}

// infoBoxHeight returns the height of the info box.
func (m App) infoBoxHeight() int {
	if len(m.filtered) == 0 || m.cursor < 0 || m.cursor >= len(m.filtered) {
		return 0
	}
	width := m.effectiveWidth()
	infoBox := m.renderInfoBox(m.filtered[m.cursor], width)
	return lipgloss.Height(infoBox)
}

// footerHeight returns the height of footer elements.
func (m App) footerHeight() int {
	footer := m.renderFooter()
	searchInput := promptStyle.Render("❯ ") + m.searchInput.View()
	return lipgloss.Height(footer) + lipgloss.Height(searchInput) + 2
}

// maxVisibleItems calculates available space for list items dynamically.
func (m App) maxVisibleItems() int {
	fixedHeight := m.headerHeight() + m.infoBoxHeight() + m.footerHeight()
	available := m.height - fixedHeight

	if available < minVisibleItems {
		return minVisibleItems
	}
	return available
}

// effectiveWidth returns the usable terminal width.
func (m App) effectiveWidth() int {
	if m.width < minWidth {
		return config.DefaultWidth
	}
	return m.width
}

