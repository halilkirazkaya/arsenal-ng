// Package theme provides the shared lipgloss styles for the TUI: base
// palette, component styles, command syntax styles, and tag rendering.
package theme

// =============================================================================
// Primary Colors
// =============================================================================

// ColorPrimary is red for titles, cursors, and prompts.
const ColorPrimary = "#FF6B6B"

// ColorSecondary is cyan for tools and arguments.
const ColorSecondary = "#4ECDC4"

// ColorAccent is yellow for tags, highlights, and matches.
const ColorAccent = "#FFE66D"

// ColorDim is gray for help text and separators.
const ColorDim = "#666666"

// ColorBright is white for selected and important text.
const ColorBright = "#FFFFFF"

// ColorDesc is light gray for descriptions.
const ColorDesc = "#888888"

// =============================================================================
// Syntax Colors
// =============================================================================

// ColorCmdTool is red for command and tool names.
const ColorCmdTool = "#FF6B6B"

// ColorCmdFlag is cyan for flags like --help and -v.
const ColorCmdFlag = "#4ECDC4"

// ColorCmdDefault is orange for defaults in {{arg|default}}.
const ColorCmdDefault = "#FF8C00"

// ColorCmdString is light green for quoted strings.
const ColorCmdString = "#98D8C8"

// ColorCmdVar is gold for $variables.
const ColorCmdVar = "#F7DC6F"

// ColorCmdNormal is light gray for normal text.
const ColorCmdNormal = "#CCCCCC"

// ColorCmdPipe is purple for pipes and operators.
const ColorCmdPipe = "#BB8FCE"

// =============================================================================
// Tag Palette
// =============================================================================

// TagColors holds vibrant tag colors indexed by tag hash.
var TagColors = []string{
	"#FF6B6B",
	"#4ECDC4",
	"#FFE66D",
	"#95E1D3",
	"#F38181",
	"#AA96DA",
	"#78C4D4",
	"#F9ED69",
	"#F08A5D",
	"#B83B5E",
	"#6A0572",
	"#00B8A9",
	"#F6416C",
	"#FCBAD3",
	"#A8D8EA",
}
