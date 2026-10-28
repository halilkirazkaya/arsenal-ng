// Package model defines the core data types for arsenal-ng.
//
// This file contains argument parsing and command building functionality.
// It handles {{arg}} and {{arg|default}} placeholders in command templates,
// extracts arguments, validates them, and builds final executable commands.
package model

import (
	"regexp"
)

// =============================================================================
// Argument Pattern
// =============================================================================

// ArgPattern matches {{arg}} and {{arg|default}} placeholders in commands.
// Examples:
//   - {{ip}}           -> name="ip", default=""
//   - {{port|8080}}    -> name="port", default="8080"
//   - {{output|scan}}  -> name="output", default="scan"
//   - {{domain|}}      -> name="domain", default="" (optional, may be left empty)

var ArgPattern = regexp.MustCompile(`\{\{([^{}|]+)(?:\|([^{}]*))?\}\}`)

// =============================================================================
// Argument Parsing
// =============================================================================

// ParseArguments extracts all unique arguments from a command template.
// Arguments are returned in order of first appearance.
func ParseArguments(command string) []Argument {
	matches := ArgPattern.FindAllStringSubmatchIndex(command, -1)
	seen := make(map[string]bool)
	var args []Argument

	for _, match := range matches {
		name := command[match[2]:match[3]]

		// Skip duplicates
		if seen[name] {
			continue
		}
		seen[name] = true

		// Extract default value if present. Any argument written with "|"
		// is optional, so {{arg|}} may be submitted empty.
		var defaultVal string
		optional := match[4] != -1
		if optional {
			defaultVal = command[match[4]:match[5]]
		}

		args = append(args, Argument{
			Name:         name,
			DefaultValue: defaultVal,
			Value:        defaultVal, // Pre-fill with default
			Optional:     optional,
			Position:     match[0],
		})
	}

	return args
}

// =============================================================================
// Command Building
// =============================================================================

// BuildCommand replaces all {{arg}} placeholders with their values.
// Returns the final executable command string.
// Placeholders are replaced in a single pass, so a value that itself contains
// "{{...}}" is inserted literally instead of being substituted again.
func BuildCommand(command string, args []Argument) string {
	values := make(map[string]string, len(args))
	for _, arg := range args {
		values[arg.Name] = arg.Value
	}

	return ArgPattern.ReplaceAllStringFunc(command, func(match string) string {
		name := ArgPattern.FindStringSubmatch(match)[1]
		if value, ok := values[name]; ok {
			return value
		}
		return match
	})
}

// =============================================================================
// Validation
// =============================================================================

// HasEmptyArgs checks if any required argument is still empty.
// Returns true if at least one required argument has no value.
func HasEmptyArgs(args []Argument) bool {
	for _, arg := range args {
		if arg.Value == "" && !arg.Optional {
			return true
		}
	}
	return false
}

// GetEmptyArgs returns a list of required argument names that have no value.
func GetEmptyArgs(args []Argument) []string {
	var empty []string
	for _, arg := range args {
		if arg.Value == "" && !arg.Optional {
			empty = append(empty, arg.Name)
		}
	}
	return empty
}

