// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

import (
	"strings"
)

// =============================================================================
// Command Tokenizer
// =============================================================================

// tokenize splits a command line into tokens for syntax highlighting.
// Quoted spans, {{placeholders}}, and multi-char operators stay whole.
func tokenize(line string) []string {
	var tokens []string
	var current strings.Builder
	inQuote := rune(0)
	inBraces := 0
	skipNext := false

	for i, char := range line {
		if skipNext {
			skipNext = false
			continue
		}
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

		if inQuote != 0 {
			current.WriteRune(char)
			if char == inQuote && (i == 0 || line[i-1] != '\\') {
				tokens = append(tokens, current.String())
				current.Reset()
				inQuote = 0
			}
			continue
		}

		if char == '{' && i+1 < len(line) && line[i+1] == '{' {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			inBraces++
			current.WriteRune(char)
			continue
		}

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
			if op, ok := doubleOp(line, i, char); ok {
				tokens = append(tokens, op)
				skipNext = true
				continue
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
