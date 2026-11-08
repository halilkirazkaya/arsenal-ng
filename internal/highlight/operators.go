// Package highlight provides terminal highlighting helpers: fuzzy match
// highlighting for search results and syntax highlighting for commands.
package highlight

// =============================================================================
// Operator Detection
// =============================================================================

// doubleOp matches a two-char operator (>>, &&, ||) at position i.
// It reports the operator and true when line[i] starts one.
func doubleOp(line string, i int, char rune) (string, bool) {
	if i+1 >= len(line) {
		return "", false
	}
	next := line[i+1]
	if (char == '>' && next == '>') ||
		(char == '&' && next == '&') ||
		(char == '|' && next == '|') {
		return string(char) + string(next), true
	}
	return "", false
}

// isPipeOrOperator checks if a token is a pipe or operator.
// It keeps highlighting state resets consistent with tokenize.
func isPipeOrOperator(token string) bool {
	switch token {
	case "|", ">", "<", ">>", "&&", "||", ";":
		return true
	}
	return false
}
