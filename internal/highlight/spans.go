// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Styled Spans
// =============================================================================

// highlightArgs highlights {{arg}} and {{arg|default}} patterns.
// Placeholders with defaults use the default style, plain ones the var style.
func highlightArgs(token string) string {
	return model.ArgPattern.ReplaceAllStringFunc(token, func(match string) string {
		if strings.Contains(match, "|") {
			return theme.CmdDefault.Render(match)
		}
		return theme.CmdVar.Render(match)
	})
}

// highlightQuotedString highlights a quoted string.
// Tokens containing {{args}} never reach here because highlightToken
// routes those to highlightArgs first.
func highlightQuotedString(token string) string {
	return theme.CmdString.Render(token)
}
