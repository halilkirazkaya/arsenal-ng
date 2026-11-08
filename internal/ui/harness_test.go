// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"io"
	"log"
	"reflect"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// newTestApp builds an isolated App with temp HOME and silenced logs.
func newTestApp(t *testing.T, cheats []*model.Cheat) App {
	t.Helper()
	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	t.Setenv("XDG_CONFIG_HOME", tmp)
	log.SetOutput(io.Discard)
	if cheats == nil {
		cheats = fixtureCheats()
	}
	return New(cheats)
}

// fixtureCheats returns one arg-ful and one arg-less cheat for tests.
func fixtureCheats() []*model.Cheat {
	return []*model.Cheat{
		{Tool: "nmap", Title: "Port scan", Desc: "Scan ports", Command: "nmap {{ip}}", Filename: "nmap.yaml"},
		{Tool: "ping", Title: "Ping host", Desc: "Ping it", Command: "ping google.com", Filename: "ping.yaml"},
	}
}

// argCheat returns a cheat requiring one argument value.
func argCheat() *model.Cheat {
	return &model.Cheat{Tool: "nmap", Title: "Port scan", Command: "nmap {{ip}}"}
}

// noArgCheat returns a cheat that runs without arguments.
func noArgCheat() *model.Cheat {
	return &model.Cheat{Tool: "ping", Title: "Ping", Command: "ping google.com"}
}

// =============================================================================
// Key Helpers
// =============================================================================

// keyEnter builds an enter key message.
func keyEnter() tea.KeyPressMsg { return tea.KeyPressMsg{Code: tea.KeyEnter} }

// keyEsc builds an escape key message.
func keyEsc() tea.KeyPressMsg { return tea.KeyPressMsg{Code: tea.KeyEsc} }

// keyCtrlC builds a ctrl+c key message.
func keyCtrlC() tea.KeyPressMsg { return tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl} }

// keyRune builds a printable-rune key message.
func keyRune(s string) tea.KeyPressMsg {
	return tea.KeyPressMsg{Text: string(s)}
}

// isQuit reports whether the command is bubbletea quit.
func isQuit(cmd tea.Cmd) bool {
	if cmd == nil {
		return false
	}
	a, b := reflect.ValueOf(cmd), reflect.ValueOf(tea.Quit)
	return a.Kind() == reflect.Func && b.Kind() == reflect.Func &&
		a.Pointer() == b.Pointer()
}

// init pins the logo color so rendered snapshots stay deterministic.
func init() {
	theme.SetLogoColor(lipgloss.Color("#FFFFFF"))
}
