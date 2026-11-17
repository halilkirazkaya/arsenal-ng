// Package search ranks cheats against a query.
//
// This file scores individual cheats against query terms. Scores combine
// field weights (tool name beats description) with match-quality multipliers
// (exact beats prefix, prefix beats substring).
package search

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Score Weights
// =============================================================================

// Field weights: higher means a more significant match location.
const (
	weightTool    = 100
	weightTag     = 60
	weightTitle   = 40
	weightCommand = 20
	weightDesc    = 10
)

// Match quality multipliers within a field.
const (
	matchExact  = 4
	matchPrefix = 2
	matchSub    = 1
)

// =============================================================================
// Cheat Scoring
// =============================================================================

// scoreCheat sums per-term scores, reporting false when any term matches
// nowhere. The false case preserves the "all terms must match" result set.
func scoreCheat(cheat *model.Cheat, terms []string) (score int, ok bool) {
	tool := strings.ToLower(cheat.Tool)
	title := strings.ToLower(cheat.Title)
	command := strings.ToLower(cheat.Command)
	desc := strings.ToLower(cheat.Desc)
	tags := make([]string, len(cheat.Tags))
	for i, tag := range cheat.Tags {
		tags[i] = strings.ToLower(tag)
	}

	for _, term := range terms {
		// A term matching tool AND title outranks one matching only one field.
		sum := matchScore(tool, term, weightTool)
		for _, tag := range tags {
			sum += matchScore(tag, term, weightTag)
		}
		sum += matchScore(title, term, weightTitle)
		sum += matchScore(command, term, weightCommand)
		sum += matchScore(desc, term, weightDesc)
		if sum == 0 {
			return 0, false
		}
		score += sum
	}
	return score, true
}

// matchScore weights a single term against one field value.
// Exact matches outrank prefix matches, which outrank substring matches.
func matchScore(field, term string, weight int) int {
	switch {
	case field == "":
		return 0
	case field == term:
		return weight * matchExact
	case strings.HasPrefix(field, term):
		return weight * matchPrefix
	case strings.Contains(field, term):
		return weight * matchSub
	}
	return 0
}
