// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Syntax Highlighting
// =============================================================================

// Syntax applies syntax highlighting to a command string.
// Multi-line commands are highlighted line by line.
func Syntax(cmd string) string {
	lines := strings.Split(cmd, "\n")
	result := make([]string, len(lines))

	for i, line := range lines {
		result[i] = highlightLine(line)
	}

	return strings.Join(result, "\n")
}

// highlightLine applies syntax highlighting to a single line.
// Blank lines pass through unstyled.
func highlightLine(line string) string {
	if strings.TrimSpace(line) == "" {
		return line
	}

	var result strings.Builder
	tokens := tokenize(line)
	isFirst := true

	for _, token := range tokens {
		result.WriteString(highlightToken(token, &isFirst))
	}

	return result.String()
}

// highlightToken applies the appropriate style to a single token.
// The isFirst flag tracks whether the next word opens a new command.
func highlightToken(token string, isFirst *bool) string {
	switch {
	case model.ArgPattern.MatchString(token):
		return highlightArgs(token)

	case strings.HasPrefix(token, "{{") && !strings.HasSuffix(token, "}}"):
		// Truncated arg pattern
		if strings.Contains(token, "|") {
			return theme.CmdDefault.Render(token)
		}
		return theme.CmdVar.Render(token)

	case *isFirst && !strings.HasPrefix(token, " "):
		*isFirst = false
		return theme.CmdTool.Render(token)

	case strings.HasPrefix(token, "--") || strings.HasPrefix(token, "-"):
		return theme.CmdFlag.Render(token)

	case strings.HasPrefix(token, "'") || strings.HasPrefix(token, "\""):
		return highlightQuotedString(token)

	case strings.HasPrefix(token, "$"):
		return theme.CmdVar.Render(token)

	case isPipeOrOperator(token):
		*isFirst = true
		return theme.CmdPipe.Render(token)

	case strings.TrimSpace(token) == "":
		return token

	default:
		if !strings.HasPrefix(token, " ") && strings.TrimSpace(token) != "" {
			*isFirst = false
		}
		return theme.CmdNormal.Render(token)
	}
}
