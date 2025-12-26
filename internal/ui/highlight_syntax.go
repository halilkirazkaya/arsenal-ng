// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains syntax highlighting functions for command strings. It
// applies color coding to different parts of commands (tools, flags, arguments,
// strings, variables) to improve readability.
package ui

import (
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Command Syntax Highlighting
// =============================================================================

// syntaxHighlight applies syntax highlighting to a command string.
func syntaxHighlight(cmd string) string {
	lines := strings.Split(cmd, "\n")
	result := make([]string, len(lines))

	for i, line := range lines {
		result[i] = highlightLine(line)
	}

	return strings.Join(result, "\n")
}

// highlightLine applies syntax highlighting to a single line.
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
func highlightToken(token string, isFirst *bool) string {
	switch {
	case model.ArgPattern.MatchString(token):
		return highlightArgs(token)

	case strings.HasPrefix(token, "{{") && !strings.HasSuffix(token, "}}"):
		// Truncated arg pattern
		if strings.Contains(token, "|") {
			return cmdDefaultStyle.Render(token)
		}
		return cmdVarStyle.Render(token)

	case *isFirst && !strings.HasPrefix(token, " "):
		*isFirst = false
		return cmdToolStyle.Render(token)

	case strings.HasPrefix(token, "--") || strings.HasPrefix(token, "-"):
		return cmdFlagStyle.Render(token)

	case strings.HasPrefix(token, "'") || strings.HasPrefix(token, "\""):
		return highlightQuotedString(token)

	case strings.HasPrefix(token, "$"):
		return cmdVarStyle.Render(token)

	case isPipeOrOperator(token):
		*isFirst = true
		return cmdPipeStyle.Render(token)

	case strings.TrimSpace(token) == "":
		return token

	default:
		if !strings.HasPrefix(token, " ") && strings.TrimSpace(token) != "" {
			*isFirst = false
		}
		return cmdNormalStyle.Render(token)
	}
}

// isPipeOrOperator checks if a token is a pipe or operator.
func isPipeOrOperator(token string) bool {
	switch token {
	case "|", ">", "<", ">>", "&&", "||", ";":
		return true
	}
	return false
}

// highlightArgs highlights {{arg}} and {{arg|default}} patterns.
func highlightArgs(token string) string {
	return model.ArgPattern.ReplaceAllStringFunc(token, func(match string) string {
		if strings.Contains(match, "|") {
			return cmdDefaultStyle.Render(match)
		}
		return cmdVarStyle.Render(match)
	})
}

// highlightQuotedString highlights a quoted string, including {{args}} inside.
func highlightQuotedString(token string) string {
	if !model.ArgPattern.MatchString(token) {
		return cmdStringStyle.Render(token)
	}

	var result strings.Builder
	lastEnd := 0
	matches := model.ArgPattern.FindAllStringSubmatchIndex(token, -1)

	for _, match := range matches {
		if match[0] > lastEnd {
			result.WriteString(cmdStringStyle.Render(token[lastEnd:match[0]]))
		}

		argMatch := token[match[0]:match[1]]
		if strings.Contains(argMatch, "|") {
			result.WriteString(cmdDefaultStyle.Render(argMatch))
		} else {
			result.WriteString(cmdVarStyle.Render(argMatch))
		}
		lastEnd = match[1]
	}

	if lastEnd < len(token) {
		result.WriteString(cmdStringStyle.Render(token[lastEnd:]))
	}

	return result.String()
}

