// Package model defines the core data types for arsenal-ng.
//
// This file builds final executable commands by substituting argument
// values into {{arg}} placeholders.
package model

import "strings"

// Command Building
// =============================================================================

// BuildCommand replaces all {{arg}} placeholders with their values.
// Placeholders are replaced in a single pass, so a value that itself contains
// "{{...}}" is inserted literally instead of being substituted again.
// Values that look like filesystem paths are shell-quoted when the
// placeholder stands alone as a word (e.g. "cat {{file}}"), so a completed
// path with spaces stays one shell word. Placeholders already inside quotes
// or embedded in a larger word are left untouched, and values without shell
// metacharacters are never modified.
func BuildCommand(command string, args []Argument) string {
	values := make(map[string]string, len(args))
	for _, arg := range args {
		values[arg.Name] = arg.Value
	}

	var b strings.Builder
	locs := ArgPattern.FindAllStringSubmatchIndex(command, -1)
	pos := 0
	for _, loc := range locs {
		start, end := loc[0], loc[1]
		name := command[loc[2]:loc[3]]
		value, ok := values[name]
		if !ok {
			continue
		}
		b.WriteString(command[pos:start])
		if needsShellQuoteAt(command, start, end, value) {
			b.WriteString(shellQuote(value))
		} else {
			b.WriteString(value)
		}
		pos = end
	}
	b.WriteString(command[pos:])
	return b.String()
}

// needsShellQuoteAt reports whether value must be quoted at the placeholder
// spanning command[start:end]. Only standalone words (bounded by whitespace
// or string edges) are quoted, and only when the same span is not already
// inside single or double quotes in the template. A leading ~/ is excluded
// from the metacharacter check so plain home-relative paths stay unquoted
// and expandable; shellQuote then keeps ~/ outside the quotes.
func needsShellQuoteAt(command string, start, end int, value string) bool {
	if !looksLikePath(value) || !isStandaloneWordAt(command, start, end) || isQuotedAt(command, start) {
		return false
	}
	rest := strings.TrimPrefix(value, "~/")
	return strings.ContainsAny(rest, " \t\n\"'\\$`!*?[]{}()<>|&;~#")
}

// isStandaloneWordAt reports whether command[start:end] spans a whole shell
// word: the characters around it are whitespace or the ends of the string.
func isStandaloneWordAt(command string, start, end int) bool {
	if start > 0 {
		if c := command[start-1]; c != ' ' && c != '\t' && c != '\n' {
			return false
		}
	}
	if end < len(command) {
		if c := command[end]; c != ' ' && c != '\t' && c != '\n' {
			return false
		}
	}
	return true
}

// isQuotedAt reports whether offset sits inside single or double quotes. It
// scans from the start of the command tracking quote state; an odd number
// of unescaped quotes before the match means quoted context.
func isQuotedAt(command string, offset int) bool {
	var single, double bool
	escaped := false
	for j := range offset {
		c := command[j]
		if escaped {
			escaped = false
			continue
		}
		switch c {
		case '\\':
			escaped = true
		case '\'':
			if !double {
				single = !single
			}
		case '"':
			if !single {
				double = !double
			}
		}
	}
	return single || double
}

// looksLikePath reports whether value resembles a filesystem path: the
// same prefixes pathSuggestions completes on.
func looksLikePath(value string) bool {
	return strings.HasPrefix(value, "/") ||
		strings.HasPrefix(value, "~/") ||
		strings.HasPrefix(value, "./") ||
		strings.HasPrefix(value, "../")
}

// shellQuote quotes s as one shell word, keeping a leading ~/ outside the
// quotes so the shell still expands it to the home directory: ~/report.txt
// stays unquoted, ~/annual report.txt becomes ~/'annual report.txt'.
// Embedded single quotes use the standard POSIX '"'"' idiom.
func shellQuote(s string) string {
	if after, ok := strings.CutPrefix(s, "~/"); ok {
		return "~/" + quoteWord(after)
	}
	return quoteWord(s)
}

// quoteWord wraps s in single quotes, escaping embedded quotes the
// standard POSIX way ('foo'"'"'bar').
func quoteWord(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'"
}
