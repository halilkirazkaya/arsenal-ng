// Package help provides the help view: keyboard shortcuts, usage
// examples, and workflow instructions.
package help

import (
	"fmt"
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/frame"
)

// =============================================================================
// Help Rendering
// =============================================================================

// View renders the help screen.
func View() string {
	var b strings.Builder

	b.WriteString(frame.Header())
	b.WriteString("\n\n")

	b.WriteString(sectionStyle.Render("📖 HELP"))
	b.WriteString("\n\n")

	b.WriteString(navigationSection())
	b.WriteString("\n")

	b.WriteString(argumentsSection())
	b.WriteString("\n")

	b.WriteString(variablesSection())
	b.WriteString("\n")

	b.WriteString(workflow())

	b.WriteString("\n\n")
	b.WriteString(helpFooter())

	return b.String()
}

// =============================================================================
// Help Sections
// =============================================================================

// navigationSection renders the navigation key bindings.
func navigationSection() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("  NAVIGATION"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("↑ / ↓        "), descStyle.Render("Move up/down in the list")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("PgUp / PgDown"), descStyle.Render("Jump one page up/down")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Enter        "), descStyle.Render("Select the highlighted command")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("ESC          "), descStyle.Render("Exit arsenal or go back")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("help         "), descStyle.Render("Show this help screen")))
	return b.String()
}

// argumentsSection renders the argument input key bindings.
func argumentsSection() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("  ARGUMENT INPUT"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Tab          "), descStyle.Render("Move to next argument field")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Shift+Tab    "), descStyle.Render("Move to previous argument field")))
	b.WriteString(fmt.Sprintf("  %s  %s\n", keyStyle.Render("Enter        "), descStyle.Render("Execute command with filled arguments")))
	return b.String()
}

// variablesSection renders the global variables command reference.
func variablesSection() string {
	var b strings.Builder
	b.WriteString(sectionStyle.Render("  🚀 GLOBAL VARIABLES"))
	b.WriteString("\n")
	b.WriteString(lineStyle.Render(separator))
	b.WriteString("\n")
	b.WriteString(descStyle.Render("  Set variables once, use them everywhere!\n"))
	b.WriteString("\n")
	// Align arrows by padding commands to fixed width.
	cmdWidth := 20
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("set key=value"), strings.Repeat(" ", cmdWidth-len("set key=value")), arrowStyle.Render("→"), descStyle.Render("Sets a variable")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("unset key"), strings.Repeat(" ", cmdWidth-len("unset key")), arrowStyle.Render("→"), descStyle.Render("Removes a variable")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("variables"), strings.Repeat(" ", cmdWidth-len("variables")), arrowStyle.Render("→"), descStyle.Render("Lists all your variables")))
	b.WriteString(fmt.Sprintf("  %s%s %s %s\n", cmdStyle.Render("tools"), strings.Repeat(" ", cmdWidth-len("tools")), arrowStyle.Render("→"), descStyle.Render("Lists all available tools")))
	return b.String()
}
