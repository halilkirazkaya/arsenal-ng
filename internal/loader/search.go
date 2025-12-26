// Package loader handles loading and searching cheat files.
//
// This file contains search functionality that filters cheats based on query
// strings. It supports multi-word queries where all terms must match, and
// searches across tool names, tags, titles, commands, and descriptions.
package loader

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Search
// =============================================================================

// Search filters cheats based on a query string.
// Supports multi-word queries where all terms must match.
// Searches across: tool name, tags, title, command, and description.
func Search(cheats []*model.Cheat, query string) []*model.Cheat {
	if query == "" {
		return cheats
	}

	query = strings.ToLower(query)
	terms := strings.Fields(query)

	var results []*model.Cheat
	for _, cheat := range cheats {
		if matchesAllTerms(cheat, terms) {
			results = append(results, cheat)
		}
	}

	return results
}

// matchesAllTerms checks if a cheat matches all search terms.
func matchesAllTerms(cheat *model.Cheat, terms []string) bool {
	searchText := buildSearchText(cheat)

	for _, term := range terms {
		if !strings.Contains(searchText, term) {
			return false
		}
	}

	return true
}

// buildSearchText creates a searchable string from all cheat fields.
func buildSearchText(cheat *model.Cheat) string {
	return strings.ToLower(
		cheat.Tool + " " +
			strings.Join(cheat.Tags, " ") + " " +
			cheat.Title + " " +
			cheat.Command + " " +
			cheat.Desc,
	)
}

