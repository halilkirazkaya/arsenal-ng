// Package model defines the core data types for arsenal-ng.
//
// This file builds final executable commands by substituting argument
// values into {{arg}} placeholders.
package model

// =============================================================================
// Command Building
// =============================================================================

// BuildCommand replaces all {{arg}} placeholders with their values.
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
