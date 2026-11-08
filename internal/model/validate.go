// Package model defines the core data types for arsenal-ng.
//
// This file validates whether required arguments still lack values.
package model

// =============================================================================
// Argument Validation
// =============================================================================

// HasEmptyArgs checks if any required argument is still empty.
// It returns true when at least one required argument has no value.
func HasEmptyArgs(args []Argument) bool {
	for _, arg := range args {
		if arg.Value == "" && !arg.Optional {
			return true
		}
	}
	return false
}

// GetEmptyArgs returns required argument names that have no value.
// Optional arguments are never included, even when empty.
func GetEmptyArgs(args []Argument) []string {
	var empty []string
	for _, arg := range args {
		if arg.Value == "" && !arg.Optional {
			empty = append(empty, arg.Name)
		}
	}
	return empty
}
