// Package theme provides the shared lipgloss styles for the TUI: base
// palette, component styles, command syntax styles, and tag rendering.
package theme

import (
	"charm.land/lipgloss/v2"
)

// =============================================================================
// Syntax Styles
// =============================================================================

// CmdTool renders command and tool names in bold red.
var CmdTool = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorCmdTool)).Bold(true)

// CmdFlag renders flags like --help in cyan.
var CmdFlag = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorCmdFlag))

// CmdDefault renders {{arg|default}} defaults in bold orange.
var CmdDefault = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorCmdDefault)).Bold(true)

// CmdString renders quoted strings in light green.
var CmdString = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorCmdString))

// CmdVar renders $variables in bold gold.
var CmdVar = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorCmdVar)).Bold(true)

// CmdNormal renders ordinary command text in light gray.
var CmdNormal = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorCmdNormal))

// CmdPipe renders pipes and operators in purple.
var CmdPipe = lipgloss.NewStyle().Foreground(lipgloss.Color(ColorCmdPipe))
