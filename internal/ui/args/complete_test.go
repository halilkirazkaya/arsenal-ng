// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
)

// TestTabCompletesPathThenAdvances verifies Tab completes a unique path
// prefix (dirs get "/", "~/" kept, dotfiles hidden) and moves to the next
// field once nothing is left to complete.
func TestTabCompletesPathThenAdvances(t *testing.T) {
	g := testStore(t, nil) // sets HOME to a temp dir
	home, _ := os.UserHomeDir()
	for _, d := range []string{"wordlists", "web"} {
		if err := os.MkdirAll(filepath.Join(home, d), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	for _, f := range []string{"wordlists/rockyou.txt", ".hidden"} {
		if err := os.WriteFile(filepath.Join(home, f), nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}

	if got := pathSuggestions("~/w"); !slices.Equal(got, []string{"~/web/", "~/wordlists/"}) {
		t.Fatalf("pathSuggestions(~/w) = %v", got)
	}
	if slices.Contains(pathSuggestions("~/"), "~/.hidden") || !slices.Contains(pathSuggestions("~/."), "~/.hidden") {
		t.Fatal("dotfiles must only appear after a typed dot")
	}
	if got := pathSuggestions("10.0.0.1"); got != nil {
		t.Fatalf("non-path value must not complete: %v", got)
	}

	tab := tea.KeyPressMsg{Code: tea.KeyTab}
	m := New(plainCheat(), g)
	m.Inputs[0].SetValue("~/wo")
	m.suggest()

	m, _, _, _ = m.Update(tab)
	if got := m.Inputs[0].Value(); got != "~/wordlists/" {
		t.Fatalf("after 1st tab: %q", got)
	}
	m, _, _, _ = m.Update(tab)
	if got := m.Inputs[0].Value(); got != "~/wordlists/rockyou.txt" {
		t.Fatalf("after 2nd tab: %q", got)
	}
	// Several matches are listed under the field.
	m.Inputs[0].SetValue("~/w")
	m.suggest()
	if v := m.View(120); !strings.Contains(v, "▸ web/") || !strings.Contains(v, "wordlists/") {
		t.Fatalf("view missing completion list:\n%s", v)
	}
	m.Inputs[0].SetValue("~/wordlists/rockyou.txt")
	m.suggest()

	m, _, _, _ = m.Update(tab)
	if m.Cursor != 1 || m.Inputs[0].Value() != "~/wordlists/rockyou.txt" {
		t.Fatalf("3rd tab: cursor=%d value=%q", m.Cursor, m.Inputs[0].Value())
	}
}

// chdirTemp creates a temp dir, moves the test into it, and restores the
// previous working directory on cleanup.
func chdirTemp(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(old) })
	return dir
}

// TestPathSuggestionsCoversRelativeAbsoluteAndSymlinks verifies the ./ and
// ../ prefixes, absolute paths, symlink-to-directory trailing slashes, the
// exact-match exclusion, and failed directory reads.
func TestPathSuggestionsCoversRelativeAbsoluteAndSymlinks(t *testing.T) {
	g := testStore(t, nil) // sets HOME to a temp dir
	_ = g
	dir := chdirTemp(t)
	work := filepath.Join(dir, "work")
	docs := filepath.Join(dir, "docs")
	for _, d := range []string{work, docs} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	for _, f := range []string{"work/report", "work/report.txt"} {
		if err := os.WriteFile(filepath.Join(dir, f), nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	// "work/link" is a symlink to ../docs, a directory.
	if err := os.Symlink(docs, filepath.Join(work, "link")); err != nil {
		t.Fatal(err)
	}

	// Path reads below are relative to "work".
	if err := os.Chdir(work); err != nil {
		t.Fatal(err)
	}

	// Exact matches are excluded so the longer sibling stays completable.
	if got := pathSuggestions("./report"); !slices.Equal(got, []string{"./report.txt"}) {
		t.Fatalf("pathSuggestions(./report) = %v", got)
	}

	// ./ completes entries in the working dir; symlinked dirs end in "/".
	if got, want := pathSuggestions("./"), []string{"./link/", "./report", "./report.txt"}; !slices.Equal(got, want) {
		t.Fatalf("pathSuggestions(./) = %v, want %v", got, want)
	}
	if got, want := pathSuggestions("../"), []string{"../docs/", "../work/"}; !slices.Equal(got, want) {
		t.Fatalf("pathSuggestions(../) = %v, want %v", got, want)
	}

	// Absolute paths and nonexistent directories.
	if got := pathSuggestions(filepath.Join(dir, "work", "re")); !slices.Equal(got, []string{filepath.Join(dir, "work", "report"), filepath.Join(dir, "work", "report.txt")}) {
		t.Fatalf("pathSuggestions(abs) = %v", got)
	}
	if got := pathSuggestions("./nope/x"); got != nil {
		t.Fatalf("pathSuggestions under missing dir = %v, want nil", got)
	}
}

// TestTabAfterCtrlPWithoutSuggestionsAdvances is the regression test for
// the panic CodeRabbit found: with no matches, Ctrl+P leaves the bubbles
// suggestion index at -1 and a plain CurrentSuggestion() call panics. Tab
// must instead move to the next field.
func TestTabAfterCtrlPWithoutSuggestionsAdvances(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	m.Inputs[0].SetValue("10.0.0.1") // not a path: no suggestions
	m.suggest()

	ctrlP := tea.KeyPressMsg{Code: 'p', Mod: tea.ModCtrl}
	tab := tea.KeyPressMsg{Code: tea.KeyTab}

	m, _, _, _ = m.Update(ctrlP)
	if idx := m.Inputs[0].CurrentSuggestionIndex(); idx != -1 {
		t.Fatalf("suggestion index after ctrl+p = %d, want -1", idx)
	}

	m, _, _, _ = m.Update(tab) // must not panic
	if m.Cursor != 1 {
		t.Fatalf("cursor after tab = %d, want 1", m.Cursor)
	}

	// Shift+Tab back must refresh suggestions for the refocused field.
	home, _ := os.UserHomeDir()
	if err := os.MkdirAll(filepath.Join(home, "web"), 0o755); err != nil {
		t.Fatal(err)
	}
	m.Inputs[0].SetValue("~/w")
	m = m.Prev() // refocuses field 0 and re-runs suggest()
	if got := m.Inputs[0].MatchedSuggestions(); len(got) == 0 || got[0] != "~/web/" {
		t.Fatalf("prev did not refresh suggestions: %v", got)
	}
}

// TestTabCompletesCtrlNSelection verifies Tab honors the match chosen
// with Ctrl+N instead of always taking the first suggestion.
func TestTabCompletesCtrlNSelection(t *testing.T) {
	g := testStore(t, nil)
	home, _ := os.UserHomeDir()
	for _, d := range []string{"web", "wordlists"} {
		if err := os.MkdirAll(filepath.Join(home, d), 0o755); err != nil {
			t.Fatal(err)
		}
	}

	ctrlN := tea.KeyPressMsg{Code: 'n', Mod: tea.ModCtrl}
	tab := tea.KeyPressMsg{Code: tea.KeyTab}
	m := New(plainCheat(), g)
	m.Inputs[0].SetValue("~/w")
	m.suggest()

	m, _, _, _ = m.Update(ctrlN)
	if idx := m.Inputs[0].CurrentSuggestionIndex(); idx != 1 {
		t.Fatalf("suggestion index after ctrl+n = %d, want 1", idx)
	}
	m, _, _, _ = m.Update(tab)
	if got := m.Inputs[0].Value(); got != "~/wordlists/" {
		t.Fatalf("after ctrl+n + tab: %q, want ~/wordlists/", got)
	}
}

// TestMatchesPagination verifies the completion list pages around the
// Ctrl+N selection and reports the total count.
func TestMatchesPagination(t *testing.T) {
	g := testStore(t, nil)
	home, _ := os.UserHomeDir()
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("f%02d.txt", i)
		if err := os.WriteFile(filepath.Join(home, name), nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}

	ctrlN := tea.KeyPressMsg{Code: 'n', Mod: tea.ModCtrl}
	m := New(plainCheat(), g)
	m.Inputs[0].SetValue("~/f")
	m.suggest()

	v := m.View(120)
	for _, want := range []string{"▸ f01.txt", "f08.txt", "1-8 of 10"} {
		if !strings.Contains(v, want) {
			t.Fatalf("first page missing %q:\n%s", want, v)
		}
	}

	// Eight Ctrl+N presses move the selection onto the second page.
	for range 8 {
		m, _, _, _ = m.Update(ctrlN)
	}
	v = m.View(120)
	for _, want := range []string{"▸ f09.txt", "f10.txt", "9-10 of 10"} {
		if !strings.Contains(v, want) {
			t.Fatalf("second page missing %q:\n%s", want, v)
		}
	}
}
