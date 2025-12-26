// Package model defines the core data types for arsenal-ng.
//
// This file contains the data structures for cheat files (CheatFile, Action)
// and runtime types (Cheat) used throughout the application. It defines the
// structure of YAML cheat files and their runtime representations.
package model

// =============================================================================
// YAML File Structure
// =============================================================================

// CheatFile represents the structure of a YAML cheat file.
// Each file contains one tool with multiple actions.
type CheatFile struct {
	Tool    string   `yaml:"tool"`    // Tool name (e.g., "nmap", "ffuf")
	Tags    []string `yaml:"tags"`    // Tags (e.g., ["scan", "recon"])
	Actions []Action `yaml:"actions"` // List of commands
}

// Action represents a single command entry in a cheat file.
type Action struct {
	Title   string `yaml:"title"`         // Display title
	Desc    string `yaml:"desc,omitempty"` // Description
	Command string `yaml:"command"`       // The actual command template
}

// =============================================================================
// Runtime Types
// =============================================================================

// Cheat is the runtime representation of a command.
// It's an enriched, flattened version of CheatFile + Action.
type Cheat struct {
	Tool     string   // Parent tool name
	Tags     []string // Inherited from parent CheatFile
	Title    string   // Command title
	Desc     string   // Command description
	Command  string   // Command template with {{placeholders}}
	Filename string   // Source file path (for debugging)
}

// Argument represents a placeholder in a command template.
// Placeholders use the format {{name}} or {{name|default}}.
type Argument struct {
	Name         string // Argument name (e.g., "ip", "port")
	DefaultValue string // Default value if specified with |
	Value        string // Current value (user input or default)
	Position     int    // Position in command string
}
