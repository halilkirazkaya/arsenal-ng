// Package model defines the core data types for arsenal-ng.
//
// This file contains argument parsing and command building functionality.
// It handles {{arg}} and {{arg|default}} placeholders in command templates,
// extracts arguments, validates them, and builds final executable commands.
package model

import (
	"testing"
)

// =============================================================================
// Argument Parsing Tests
// =============================================================================

// TestParseArguments verifies ordered deduplicated extraction of placeholders.
func TestParseArguments(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
		want []Argument
	}{
		{"no args", "nmap -sV", nil},
		{"single required", "nmap {{ip}}", []Argument{
			{Name: "ip", Value: "", Optional: false, Position: 5},
		}},
		{"single with default", "ffuf {{word|admin}}", []Argument{
			{Name: "word", DefaultValue: "admin", Value: "admin", Optional: true, Position: 5},
		}},
		{"empty default is optional", "nmap {{domain|}}", []Argument{
			{Name: "domain", DefaultValue: "", Value: "", Optional: true, Position: 5},
		}},
		{"order of first appearance", "a {{b}} {{a}} {{b}}", []Argument{
			{Name: "b", Value: "", Position: 2},
			{Name: "a", Value: "", Position: 8},
		}},
		{"duplicate keeps first", "x {{ip}} {{ip}}", []Argument{
			{Name: "ip", Value: "", Position: 2},
		}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ParseArguments(tc.cmd)
			if len(got) != len(tc.want) {
				t.Fatalf("ParseArguments(%q) = %+v, want %+v", tc.cmd, got, tc.want)
			}
			for i := range got {
				if got[i].Name != tc.want[i].Name ||
					got[i].DefaultValue != tc.want[i].DefaultValue ||
					got[i].Value != tc.want[i].Value ||
					got[i].Optional != tc.want[i].Optional ||
					got[i].Position != tc.want[i].Position {
					t.Fatalf("ParseArguments(%q)[%d] = %+v, want %+v", tc.cmd, i, got[i], tc.want[i])
				}
			}
		})
	}
}
