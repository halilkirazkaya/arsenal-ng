// Package help provides the help view: keyboard shortcuts, usage
// examples, and workflow instructions.
package help

import (
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Help Styles
// =============================================================================

// Local styles for help screen.
var (
	// sectionStyle renders section headings.
	sectionStyle = lipgloss.NewStyle().Foreground(theme.Primary).Bold(true)
	// lineStyle renders separator lines.
	lineStyle = lipgloss.NewStyle().Foreground(theme.Secondary)
	// keyStyle renders key names.
	keyStyle = lipgloss.NewStyle().Foreground(theme.Accent).Bold(true)
	// descStyle renders description text.
	descStyle = lipgloss.NewStyle().Foreground(theme.Bright)
	// cmdStyle renders example commands.
	cmdStyle = lipgloss.NewStyle().Foreground(theme.Secondary)
	// arrowStyle renders mapping arrows.
	arrowStyle = lipgloss.NewStyle().Foreground(theme.Dim)
	// exampleStyle renders workflow step numbers.
	exampleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#98D8C8"))
	// varStyle renders variable placeholders.
	varStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F7DC6F")).Bold(true)
)

// separator is the horizontal rule under each section heading.
const separator = "  ════════════════════════════════════════════════"
