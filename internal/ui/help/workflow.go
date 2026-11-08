// Package help provides the help view: keyboard shortcuts, usage
// examples, and workflow instructions.
package help

import (
	"fmt"
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Workflow Content
// =============================================================================

// workflow renders the example workflow and before/after comparison.
func workflow() string {
	var b strings.Builder

	b.WriteString(sectionStyle.Render("  📋 EXAMPLE WORKFLOW"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("1."), descStyle.Render("Type: ")+cmdStyle.Render("set key=value")+" "+keyStyle.Render("[Enter]")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("2."), descStyle.Render("Example: ")+cmdStyle.Render("set ip=10.10.10.10")+" "+keyStyle.Render("[Enter]")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("3."), descStyle.Render("Select any command with ")+varStyle.Render("{{ip}}")+" "+descStyle.Render("or ")+varStyle.Render("{{domain}}")))
	b.WriteString(fmt.Sprintf("  %s %s\n", exampleStyle.Render("4."), theme.Success.Render("Arguments will be auto-filled! ✓")))
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("  %s nmap -sV %s %s %s\n",
		theme.Help.Render("Before:"), varStyle.Render("{{ip}}"), arrowStyle.Render("→"), theme.Help.Render("You type IP every time")))
	b.WriteString(fmt.Sprintf("  %s nmap -sV %s  %s %s\n",
		theme.Success.Render("After: "), exampleStyle.Render("10.10.10.10"), arrowStyle.Render("→"), theme.Success.Render("Auto-filled! ✓")))

	return b.String()
}

// helpFooter renders the back hint.
func helpFooter() string {
	return theme.Help.Render("  Press ESC or Enter to go back")
}
