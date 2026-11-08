// Package config provides application metadata, version info, and UI
// defaults such as terminal dimensions and input constraints.
package config

// =============================================================================
// Terminal Dimensions
// =============================================================================

// DefaultTerminalWidth is the fallback width used for narrow terminals.
const DefaultTerminalWidth = 80

// DefaultTerminalHeight is the fallback height used for short terminals.
const DefaultTerminalHeight = 24

// =============================================================================
// Input Constraints
// =============================================================================

// SearchPlaceholder is the default hint shown in the empty search box.
const SearchPlaceholder = "Search commands..."

// SearchMaxLength caps the length of search input.
const SearchMaxLength = 100

// ArgumentMaxLength caps the length of argument input.
const ArgumentMaxLength = 200

// =============================================================================
// Compatibility Aliases
// =============================================================================

// DefaultWidth aliases DefaultTerminalWidth for older callers.
const DefaultWidth = DefaultTerminalWidth

// DefaultHeight aliases DefaultTerminalHeight for older callers.
const DefaultHeight = DefaultTerminalHeight

// SearchCharLimit aliases SearchMaxLength for older callers.
const SearchCharLimit = SearchMaxLength

// ArgCharLimit aliases ArgumentMaxLength for older callers.
const ArgCharLimit = ArgumentMaxLength
