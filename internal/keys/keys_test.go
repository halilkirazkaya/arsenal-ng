// Package keys centralizes the keyboard bindings and special command
// literals used across the TUI views.
package keys

import "testing"

// =============================================================================
// Binding Values
// =============================================================================

// TestKeyBindings verifies every key literal matches bubbletea naming.
func TestKeyBindings(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"Up", Up, "up"},
		{"Down", Down, "down"},
		{"Left", Left, "left"},
		{"Right", Right, "right"},
		{"VimLeft", VimLeft, "h"},
		{"VimRight", VimRight, "l"},
		{"CtrlP", CtrlP, "ctrl+p"},
		{"CtrlN", CtrlN, "ctrl+n"},
		{"PgUp", PgUp, "pgup"},
		{"PgDown", PgDown, "pgdown"},
		{"Enter", Enter, "enter"},
		{"Esc", Esc, "esc"},
		{"CtrlC", CtrlC, "ctrl+c"},
		{"Help", Help, "?"},
		{"Quit", Quit, "q"},
		{"Tab", Tab, "tab"},
		{"ShiftTab", ShiftTab, "shift+tab"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("%s = %q, want %q", tc.name, tc.got, tc.want)
			}
		})
	}
}

// TestSpecialCommands verifies search-box command literals stay stable.
func TestSpecialCommands(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"Set", Set, "set"},
		{"Unset", Unset, "unset"},
		{"Variables", Variables, "variables"},
		{"Tools", Tools, "tools"},
		{"HelpCmd", HelpCmd, "help"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("%s = %q, want %q", tc.name, tc.got, tc.want)
			}
		})
	}
}

// TestNoEmptyBindings verifies no binding or command literal is blank.
func TestNoEmptyBindings(t *testing.T) {
	all := map[string]string{
		"Up": Up, "Down": Down, "Left": Left, "Right": Right,
		"VimLeft": VimLeft, "VimRight": VimRight, "CtrlP": CtrlP, "CtrlN": CtrlN,
		"PgUp": PgUp, "PgDown": PgDown, "Enter": Enter, "Esc": Esc,
		"CtrlC": CtrlC, "Help": Help, "Quit": Quit, "Tab": Tab,
		"ShiftTab": ShiftTab, "Set": Set, "Unset": Unset,
		"Variables": Variables, "Tools": Tools, "HelpCmd": HelpCmd,
	}
	for name, v := range all {
		if v == "" {
			t.Fatalf("%s is empty", name)
		}
	}
}
