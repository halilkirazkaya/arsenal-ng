// Package variables provides the variables view: a table of the global
// variables set via the "set" command.
package variables

import (
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// View Constants
// =============================================================================

// title is the view heading.
const title = "Global Variables"

// hint is the footer help text.
const hint = "Press ESC or Enter to go back │ ↑/↓: navigate"

// =============================================================================
// Layout Calculation
// =============================================================================

// reserved returns the vertical space taken by header, title, and footer.
func reserved(height int) int {
	header := frame.Header()
	titleView := theme.Title.Render(title)
	footer := theme.Help.Render(hint)
	return lipgloss.Height(header) + 2 +
		lipgloss.Height(titleView) + 2 +
		lipgloss.Height(footer) + 2
}
