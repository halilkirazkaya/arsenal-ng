// Package frame renders the shared application frame: header with logo
// and version, plus the search view footer with status and help text.
package frame

import (
	"fmt"

	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Width Policy
// =============================================================================

// minWidth is the smallest terminal width still treated as usable.
const minWidth = 40

// Width clamps tiny terminals to the configured default width.
func Width(width int) int {
	if width < minWidth {
		return config.DefaultWidth
	}
	return width
}

// =============================================================================
// Frame Sections
// =============================================================================

// Header renders the logo beside the version at the top.
func Header() string {
	logo := theme.Logo.Render(config.Logo)
	version := theme.Version.Render(" " + config.GetVersionInfo())
	return lipgloss.JoinHorizontal(lipgloss.Bottom, logo, version)
}

// Footer renders status counts plus key hints at the bottom.
func Footer(cursorPos, total, varCount int) string {
	// Position readout anchors the left edge of the footer.
	status := theme.StatusBar.Render(fmt.Sprintf(" %d/%d ", cursorPos, total))

	// Variable badge appears only when variables are active.
	varInfo := ""
	if varCount > 0 {
		varInfo = theme.VarCount.Render(fmt.Sprintf("│ active variable count: %d ", varCount))
	}

	// Key hints close the footer with the available actions.
	help := theme.StatusBar.Render("│ ↑/↓: nav │ set key=value │ unset key │ variables │ tools │ help │ esc: quit")

	return status + varInfo + help
}
