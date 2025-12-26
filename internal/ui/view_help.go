// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the help view rendering function. It displays a
// comprehensive help screen with keyboard shortcuts, usage examples, and
// workflow instructions for the application.
package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Help View
// =============================================================================

func (m App) viewHelp() string {
	var b strings.Builder

	// Local styles for help screen
	sectionStyle := lipgloss.NewStyle().Foreground(primaryColor).Bold(true)
	lineStyle := lipgloss.NewStyle().Foreground(secondaryColor)
	keyStyle := lipgloss.NewStyle().Foreground(accentColor).Bold(true)
	helpDescStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(config.ColorBright))
	cmdStyle := lipgloss.NewStyle().Foreground(secondaryColor)
	arrowStyle := lipgloss.NewStyle().Foreground(dimColor)
	exampleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#98D8C8"))
	varStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F7DC6F")).Bold(true)

	separator := "  ════════════════════════════════════════════════"

	b.WriteString(m.renderHeader())
	b.WriteString("\n\n")

	// Title
	b.WriteString(sectionStyle.Render("📖 HELP"))
	b.WriteString("\n\n")

	// Navigation section
	b.WriteString(sectionStyle.Render("  NAVIGATION"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("↑ / ↓        "), helpDescStyle.Render("Move up/down in the list")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("PgUp / PgDown"), helpDescStyle.Render("Jump one page up/down")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Enter        "), helpDescStyle.Render("Select the highlighted command")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("ESC          "), helpDescStyle.Render("Exit arsenal or go back")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("help         "), helpDescStyle.Render("Show this help screen")))
	b.WriteString("\n")

	// Arguments section
	b.WriteString(sectionStyle.Render("  ARGUMENT INPUT"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Tab          "), helpDescStyle.Render("Move to next argument field")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Shift+Tab    "), helpDescStyle.Render("Move to previous argument field")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Enter        "), helpDescStyle.Render("Execute command with filled arguments")))
	b.WriteString("\n")

	// Global Variables section
	b.WriteString(sectionStyle.Render("  🚀 GLOBAL VARIABLES"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(helpDescStyle.Render("  Set variables once, use them everywhere!\n"))
	b.WriteString("\n")
	// Align arrows by padding commands to fixed width
	cmdWidth := 20
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("set key=value"), strings.Repeat(" ", cmdWidth-len("set key=value")), arrowStyle.Render("→"), helpDescStyle.Render("Sets a variable")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("unset key"), strings.Repeat(" ", cmdWidth-len("unset key")), arrowStyle.Render("→"), helpDescStyle.Render("Removes a variable")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("variables"), strings.Repeat(" ", cmdWidth-len("variables")), arrowStyle.Render("→"), helpDescStyle.Render("Lists all your variables")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("tools"), strings.Repeat(" ", cmdWidth-len("tools")), arrowStyle.Render("→"), helpDescStyle.Render("Lists all available tools")))
	b.WriteString("\n")

	// Example workflow
	b.WriteString(sectionStyle.Render("  📋 EXAMPLE WORKFLOW"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("1."), helpDescStyle.Render("Type: ")+cmdStyle.Render("set key=value")+" "+keyStyle.Render("[Enter]")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("2."), helpDescStyle.Render("Example: ")+cmdStyle.Render("set ip=10.10.10.10")+" "+keyStyle.Render("[Enter]")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("3."), helpDescStyle.Render("Select any command with ")+varStyle.Render("{{ip}}")+" "+helpDescStyle.Render("or ")+varStyle.Render("{{domain}}")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("4."), successStyle.Render("Arguments will be auto-filled! ✓")))
	b.WriteString("\n")

	// Before/After example
	b.WriteString(fmt.Sprintf("  %s nmap -sV %s %s %s\n",
		helpStyle.Render("Before:"), varStyle.Render("{{ip}}"), arrowStyle.Render("→"), helpStyle.Render("You type IP every time")))
	b.WriteString(fmt.Sprintf("  %s nmap -sV %s  %s %s\n",
		successStyle.Render("After: "), exampleStyle.Render("10.10.10.10"), arrowStyle.Render("→"), successStyle.Render("Auto-filled! ✓")))

	b.WriteString("\n\n")
	b.WriteString(helpStyle.Render("  Press ESC or Enter to go back"))

	return b.String()
}

