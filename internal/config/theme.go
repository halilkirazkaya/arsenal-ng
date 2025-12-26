// Package config provides application configuration and constants.
//
// This file defines the color palette used throughout the application including
// primary UI colors, command syntax highlighting colors, and tag colors for
// visual distinction in the TUI.
package config

// =============================================================================
// UI Color Palette
// =============================================================================

// Primary UI colors

const (
	ColorPrimary   = "#FF6B6B" // Red - titles, cursor, prompts
	ColorSecondary = "#4ECDC4" // Cyan - tools, arguments
	ColorAccent    = "#FFE66D" // Yellow - tags, highlights, matches
	ColorDim       = "#666666" // Gray - help text, separators
	ColorBright    = "#FFFFFF" // White - selected text, important
	ColorSelected  = "#3D3D3D" // Dark gray - selection background
	ColorDesc      = "#888888" // Light gray - descriptions
)

// =============================================================================
// Command Syntax Highlighting Colors
// =============================================================================

const (
	ColorCmdTool    = "#FF6B6B" // Red - command/tool name
	ColorCmdFlag    = "#4ECDC4" // Cyan - flags like --help, -v
	ColorCmdArg     = "#FFE66D" // Yellow - {{arg}} placeholders
	ColorCmdDefault = "#FF8C00" // Orange - {{arg|default}} with defaults
	ColorCmdString  = "#98D8C8" // Light green - quoted strings
	ColorCmdVar     = "#F7DC6F" // Gold - $variables
	ColorCmdNormal  = "#CCCCCC" // Light gray - normal text
	ColorCmdPipe    = "#BB8FCE" // Purple - pipes and operators
)

// =============================================================================
// Tag Color Palette
// =============================================================================

// TagColors provides vibrant, distinct colors for tag rendering.
// Each tag gets a consistent color based on its hash.

var TagColors = []string{
	"#FF6B6B", // Red
	"#4ECDC4", // Cyan
	"#FFE66D", // Yellow
	"#95E1D3", // Mint
	"#F38181", // Coral
	"#AA96DA", // Lavender
	"#78C4D4", // Sky Blue
	"#F9ED69", // Lemon
	"#F08A5D", // Orange
	"#B83B5E", // Magenta
	"#6A0572", // Purple
	"#00B8A9", // Teal
	"#F6416C", // Pink
	"#FCBAD3", // Light Pink
	"#A8D8EA", // Light Blue
}

