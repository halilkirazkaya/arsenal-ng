// Package ui provides the terminal user interface for arsenal-ng.
//
// This file defines the application state types and constants used to manage
// different views/modes of the TUI application.
package ui

// =============================================================================
// Application State
// =============================================================================

// appState represents the current view/mode of the application.
type appState int

const (
	stateSearch   appState = iota // Main search view
	stateArgs                     // Argument input view
	stateShowVars                 // Show variables view
	stateShowTools                // Show tools view
	stateHelp                     // Help screen
)

