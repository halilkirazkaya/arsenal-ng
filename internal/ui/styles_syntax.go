// Package ui provides the terminal user interface for arsenal-ng.
//
// This file defines syntax highlighting styles for command rendering. It
// includes styles for tools, flags, arguments, strings, variables, and
// operators used in command syntax highlighting.
package ui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Command Syntax Highlighting Styles
// =============================================================================

var (
	cmdToolStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorCmdTool)).Bold(true)
	cmdFlagStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorCmdFlag))
	cmdArgStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorCmdArg))
	cmdDefaultStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorCmdDefault)).Bold(true)
	cmdStringStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorCmdString))
	cmdVarStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorCmdVar)).Bold(true)
	cmdNormalStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorCmdNormal))
	cmdPipeStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorCmdPipe))
)

