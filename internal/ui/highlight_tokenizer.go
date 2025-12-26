// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains the tokenizer that splits command lines into tokens for
// syntax highlighting. It handles quotes, braces, operators, and whitespace
// to properly parse command syntax.
package ui

import (
	"strings"
)

// =============================================================================
// Tokenizer
// =============================================================================

// tokenize splits a command line into tokens for syntax highlighting.
func tokenize(line string) []string {
	var tokens []string
	var current strings.Builder
	inQuote := rune(0)
	inBraces := 0

	for i, char := range line {
		// Inside {{ }}, collect everything
		if inBraces > 0 {
			current.WriteRune(char)
			if char == '}' && i > 0 && line[i-1] == '}' {
				inBraces--
				if inBraces == 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
			}
			continue
		}

		// Inside quotes, collect everything
		if inQuote != 0 {
			current.WriteRune(char)
			if char == inQuote && (i == 0 || line[i-1] != '\\') {
				tokens = append(tokens, current.String())
				current.Reset()
				inQuote = 0
			}
			continue
		}

		// Check for {{ start
		if char == '{' && i+1 < len(line) && line[i+1] == '{' {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			inBraces++
			current.WriteRune(char)
			continue
		}

		// Handle different character types
		switch {
		case char == '"' || char == '\'':
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			current.WriteRune(char)
			inQuote = char

		case char == ' ' || char == '\t':
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(char))

		case char == '|' || char == '>' || char == '<' || char == ';' || char == '&':
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			// Check for double operators
			if i+1 < len(line) {
				next := line[i+1]
				if (char == '>' && next == '>') ||
					(char == '&' && next == '&') ||
					(char == '|' && next == '|') {
					tokens = append(tokens, string(char)+string(next))
					continue
				}
			}
			tokens = append(tokens, string(char))

		default:
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

