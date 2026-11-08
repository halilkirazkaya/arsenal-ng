// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"testing"
)

// =============================================================================
// Pick Tests
// =============================================================================

// TestPickArgLessQuitsWithCommand verifies arg-less picks quit with a command.
func TestPickArgLessQuitsWithCommand(t *testing.T) {
	b := newTestApp(t, nil)
	b.search.Cursor = 1
	b.search.Filtered = fixtureCheats()
	m, cmd := b.Update(keyEnter())
	got := m.(App)
	if got.FinalCommand != "ping google.com" {
		t.Fatalf("FinalCommand = %q", got.FinalCommand)
	}
	if !isQuit(cmd) {
		t.Fatal("arg-less pick must quit")
	}
}

// TestPickArgEntersArgsWithPrefill verifies picks prefill args from globals.
func TestPickArgEntersArgsWithPrefill(t *testing.T) {
	a := newTestApp(t, nil)
	if err := a.globals.Set("ip", "10.0.0.1"); err != nil {
		t.Fatal(err)
	}
	a.search.Cursor = 0
	a.search.Filtered = fixtureCheats()
	m, cmd := a.Update(keyEnter())
	got := m.(App)
	if got.view != viewArgs {
		t.Fatalf("view = %v, want args", got.view)
	}
	if isQuit(cmd) {
		t.Fatal("entering args must not quit")
	}
	if got.args.Inputs[0].Value() != "10.0.0.1" {
		t.Fatalf("prefill = %q, want stored var", got.args.Inputs[0].Value())
	}
}

// =============================================================================
// Command Tests
// =============================================================================

// TestFlowShowVariablesToolsHelp verifies special queries open each view.
func TestFlowShowVariablesToolsHelp(t *testing.T) {
	for query, want := range map[string]view{
		"variables": viewVariables, "tools": viewTools, "help": viewHelp,
	} {
		a := newTestApp(t, nil)
		a.search = a.search.SetQuery(query)
		m, _ := a.Update(keyEnter())
		if m.(App).view != want {
			t.Fatalf("%q: view = %v, want %v", query, m.(App).view, want)
		}
	}
}

// =============================================================================
// Cancel Tests
// =============================================================================

// TestFlowCancel verifies ctrl+c cancels the app and quits.
func TestFlowCancel(t *testing.T) {
	a := newTestApp(t, nil)
	m, cmd := a.Update(keyCtrlC())
	got := m.(App)
	if !got.Cancelled {
		t.Fatal("ctrl+c must cancel")
	}
	if !isQuit(cmd) {
		t.Fatal("cancel must quit")
	}
}
