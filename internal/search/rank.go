// Package search ranks cheats against a query.
//
// This file ranks cheats against a query. Every term must match somewhere
// (same result set as before), but results are ordered by where the terms
// match: tool name beats title, title beats command, and so on. Exact and
// prefix matches outrank substring matches within each field.
package search

import (
	"sort"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Result Ranking
// =============================================================================

// rank orders cheats by relevance to the terms. Ties keep input order.
func rank(cheats []*model.Cheat, terms []string) []*model.Cheat {
	type scored struct {
		cheat *model.Cheat
		score int
	}
	items := make([]scored, 0, len(cheats))
	for _, cheat := range cheats {
		if score, ok := scoreCheat(cheat, terms); ok {
			items = append(items, scored{cheat, score})
		}
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].score > items[j].score
	})

	results := make([]*model.Cheat, len(items))
	for i, item := range items {
		results[i] = item.cheat
	}
	return results
}
