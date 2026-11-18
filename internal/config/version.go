// Package config provides application metadata, version info, and UI
// defaults such as terminal dimensions and input constraints.
package config

// =============================================================================
// Version Metadata
// =============================================================================

// Version holds the release version injected at build time.
var Version = "2.1"

// =============================================================================
// Version Formatting
// =============================================================================

// GetVersionInfo returns the display-ready version with a leading v.
func GetVersionInfo() string {
	return "v" + Version
}
