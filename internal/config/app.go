// Package config provides application configuration and constants.
//
// This file contains application metadata (name, version, description), ASCII art
// logo, and UI default values including terminal dimensions and input constraints.
package config

// =============================================================================
// Application Metadata
// =============================================================================

var Version = "0.8"

// Application info
const (
	AppName        = "arsenal-ng"
	AppDescription = "Modern pentest command launcher"
)

// GetVersionString returns formatted version for display.
func GetVersionString() string {
	return "v" + Version
}

// =============================================================================
// ASCII Art Logo
// =============================================================================

// Logo is the ASCII art banner displayed at the top of the TUI.
const Logo = `
 █████╗ ██████╗ ███████╗███████╗███╗   ██╗ █████╗ ██╗      ███╗   ██╗ ██████╗ 
██╔══██╗██╔══██╗██╔════╝██╔════╝████╗  ██║██╔══██╗██║      ████╗  ██║██╔════╝ 
███████║██████╔╝███████╗█████╗  ██╔██╗ ██║███████║██║█████╗██╔██╗ ██║██║  ███╗
██╔══██║██╔══██║╚════██║██╔══╝  ██║╚██╗██║██╔══██║██║╚════╝██║╚██╗██║██║   ██║
██║  ██║██║  ██║███████║███████╗██║ ╚████║██║  ██║███████╗ ██║ ╚████║╚██████╔╝
╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝ ╚═╝  ╚═══╝ ╚═════╝`

// =============================================================================
// UI Defaults
// =============================================================================

// Terminal dimensions
const (
	DefaultTerminalWidth  = 80
	DefaultTerminalHeight = 24
)

// Input constraints
const (
	SearchPlaceholder   = "Search commands..."
	SearchMaxLength     = 100
	ArgumentMaxLength   = 200
)

// Backwards compatibility aliases
const (
	DefaultWidth  = DefaultTerminalWidth
	DefaultHeight = DefaultTerminalHeight
	SearchCharLimit = SearchMaxLength
	ArgCharLimit    = ArgumentMaxLength
)

// GetVersionInfo is an alias for GetVersionString (backwards compatibility).
func GetVersionInfo() string {
	return GetVersionString()
}

