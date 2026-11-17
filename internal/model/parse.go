// Package model defines the core data types for arsenal-ng.
//
// This file parses {{arg}} and {{arg|default}} placeholders from command
// templates into ordered, deduplicated arguments.
package model

import (
	"regexp"
)

// =============================================================================
// Argument Pattern
// =============================================================================

// ArgPattern matches {{arg}} and {{arg|default}} placeholders in commands.
// Examples:
//
//	{{ip}}        -> name="ip", default=""
//	{{port|8080}} -> name="port", default="8080"
//	{{domain|}}   -> name="domain", default="" (optional, may be left empty)
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

		if seen[name] {
			continue
		}
		seen[name] = true

		// Any argument written with "|" is optional, so {{arg|}} may be left empty.
		var defaultVal string
		optional := match[4] != -1
		if optional {
			defaultVal = command[match[4]:match[5]]
		}

		args = append(args, Argument{
			Name:     name,
			Optional: optional,
			Position: match[0],
			// Pre-fill with the default value.
			DefaultValue: defaultVal,
			Value:        defaultVal,
		})
	}

	return args
}
