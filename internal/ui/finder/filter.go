// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/search"
)

// =============================================================================
// Search Filtering
// =============================================================================

// Filter updates the filtered list based on the search query.
func (m Model) Filter(cheats []*model.Cheat) Model {
	if !IsSpecialCommand(m.Input.Value()) {
		m.Filtered = search.Query(cheats, m.Input.Value())
		// Filtering invalidates the old cursor position.
		m.Cursor = -1
		m.Offset = 0
	}
	return m
}

// IsSpecialCommand checks if the query is a special command.
func IsSpecialCommand(query string) bool {
	q := strings.TrimSpace(strings.ToLower(query))
	return q == keys.Set ||
		strings.HasPrefix(q, keys.Set+" ") ||
		q == keys.Unset ||
		strings.HasPrefix(q, keys.Unset+" ") ||
		q == keys.Variables ||
		q == keys.HelpCmd ||
		q == keys.Tools
}
