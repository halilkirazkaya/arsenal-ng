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
	// Tool is the tool name (e.g., "nmap", "ffuf").
	Tool string `yaml:"tool"`
	// Tags categorizes the tool (e.g., ["scan", "recon"]).
	Tags []string `yaml:"tags"`
	// Actions lists the commands in this file.
	Actions []Action `yaml:"actions"`
}

// Action represents a single command entry in a cheat file.
type Action struct {
	// Title is the display title.
	Title string `yaml:"title"`
	// Desc is the description.
	Desc string `yaml:"desc,omitempty"`
	// Command is the actual command template.
	Command string `yaml:"command"`
}

// =============================================================================
// Runtime Types
// =============================================================================

// Cheat is the runtime representation of a command.
// It is an enriched, flattened version of CheatFile + Action.
type Cheat struct {
	// Tool is the parent tool name.
	Tool string
	// Tags inherits from the parent CheatFile.
	Tags []string
	// Title is the command title.
	Title string
	// Desc is the command description.
	Desc string
	// Command is the template with {{placeholders}}.
	Command string
	// Filename is the source file path, used for debugging.
	Filename string
}

// Argument represents a placeholder in a command template.
// Placeholders use the format {{name}} or {{name|default}}.
type Argument struct {
	// Name is the argument name (e.g., "ip", "port").
	Name string
	// DefaultValue is the default when specified with "|".
	DefaultValue string
	// Value is the current value, either user input or default.
	Value string
	// Optional marks {{name|...}} args that may be left empty.
	Optional bool
	// Position is the byte offset in the command string.
	Position int
}
