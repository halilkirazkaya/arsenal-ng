// Package search ranks cheats against a query.
//
// This file contains search functionality that filters cheats based on query
// strings. It supports multi-word queries where all terms must match, and
// searches across tool names, tags, titles, commands, and descriptions.
package search

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Search
// =============================================================================

// Query filters cheats based on a query string.
// Multi-word queries require all terms to match somewhere, and results are
// ranked by match location (see score.go). Search covers tool name, tags,
// title, command, and description.
func Query(cheats []*model.Cheat, query string) []*model.Cheat {
	if query == "" {
		return cheats
	}

	query = strings.ToLower(query)
	terms := strings.Fields(query)

	return rank(cheats, terms)
}
