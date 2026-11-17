// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
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
