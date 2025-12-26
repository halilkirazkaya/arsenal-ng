// Package model defines the core data types for arsenal-ng.
//
// This file contains argument parsing and command building functionality.
// It handles {{arg}} and {{arg|default}} placeholders in command templates,
// extracts arguments, validates them, and builds final executable commands.
package model

import (
	"regexp"
	"strings"
)

// =============================================================================
// Argument Pattern
// =============================================================================

// ArgPattern matches {{arg}} and {{arg|default}} placeholders in commands.
// Examples:
//   - {{ip}}           -> name="ip", default=""
//   - {{port|8080}}    -> name="port", default="8080"
//   - {{output|scan}}  -> name="output", default="scan"

var ArgPattern = regexp.MustCompile(`\{\{([^{}|]+)(?:\|([^{}]+))?\}\}`)

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

		// Extract default value if present
		var defaultVal string
		if match[4] != -1 {
			defaultVal = command[match[4]:match[5]]
		}

		args = append(args, Argument{
			Name:         name,
			DefaultValue: defaultVal,
			Value:        defaultVal, // Pre-fill with default
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
func BuildCommand(command string, args []Argument) string {
	// First normalize all {{arg|default}} to {{arg}}
	result := ArgPattern.ReplaceAllString(command, "{{$1}}")

	// Then replace each placeholder with its value
	for _, arg := range args {
		placeholder := "{{" + arg.Name + "}}"
		result = strings.ReplaceAll(result, placeholder, arg.Value)
	}

	return result
}

// =============================================================================
// Validation
// =============================================================================

// HasEmptyArgs checks if any required argument is still empty.
// Returns true if at least one argument has no value.
func HasEmptyArgs(args []Argument) bool {
	for _, arg := range args {
		if arg.Value == "" {
			return true
		}
	}
	return false
}

// GetEmptyArgs returns a list of argument names that have no value.
func GetEmptyArgs(args []Argument) []string {
	var empty []string
	for _, arg := range args {
		if arg.Value == "" {
			empty = append(empty, arg.Name)
		}
	}
	return empty
}

