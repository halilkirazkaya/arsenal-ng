// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

// =============================================================================
// Path Completion
// =============================================================================

// pathSuggestions lists filesystem entries that extend value, keeping the
// directory part exactly as typed (including a leading "~/") so Tab
// completion only appends. Entries that already equal value are skipped so
// a longer sibling ("./report" vs "./report.txt") stays completable.
// Directories end in "/" so Tab can continue into them. Hidden entries
// only appear once the typed name starts with ".". Only values that look
// like paths touch the filesystem, so IPs, URLs and usernames cost nothing.
func pathSuggestions(value string) []string {
	if !hasAnyPrefix(value, "/", "~/", "./", "../") {
		return nil
	}

	dir, base := filepath.Split(value)
	readDir := dir
	if strings.HasPrefix(readDir, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil
		}
		readDir = home + readDir[1:]
	}

	entries, err := os.ReadDir(readDir)
	if err != nil {
		return nil
	}

	var out []string
	for _, e := range entries {
		name := e.Name()
		if !strings.HasPrefix(name, base) {
			continue
		}
		if strings.HasPrefix(name, ".") && !strings.HasPrefix(base, ".") {
			continue
		}
		// Bubbles stores suggestions as runes without sanitizing them, so a
		// name with control characters (e.g. a newline) would land in the
		// field literally; see completableName.
		if !completableName(name) {
			continue
		}
		if isDir(readDir, e) {
			name += "/"
		}
		full := dir + name
		if full == value {
			continue
		}
		out = append(out, full)
	}
	return out
}

// hasAnyPrefix reports whether s starts with any of the prefixes.
func hasAnyPrefix(s string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

// completableName reports whether a directory entry can round-trip through
// Bubbles' textinput suggestions. Control characters (newlines, tabs, ...)
// would land in the single-line field literally, and invalid UTF-8
// collapses to U+FFFD when stored as runes, so "bad\xff" becomes
// indistinguishable from a real "bad�" and Tab could complete the wrong
// file. Valid UTF-8 containing a literal U+FFFD stays completable.
func completableName(name string) bool {
	return utf8.ValidString(name) && strings.IndexFunc(name, unicode.IsControl) < 0
}

// isDir reports whether the entry is a directory, following symlinks.
// dir is concatenated, not filepath.Join'ed: dir may intentionally contain
// ".." (e.g. "./alias/../"), which Join would clean away and resolve the
// symlink target in the wrong directory. dir always ends in a separator
// here because it comes from filepath.Split (or the "~/" expansion above).
func isDir(dir string, e fs.DirEntry) bool {
	if e.Type()&fs.ModeSymlink == 0 {
		return e.IsDir()
	}
	info, err := os.Stat(dir + e.Name())
	return err == nil && info.IsDir()
}

// suggest refreshes path completions for the focused field.
func (m Model) suggest() {
	in := &m.Inputs[m.Cursor]
	in.SetSuggestions(pathSuggestions(in.Value()))
}
