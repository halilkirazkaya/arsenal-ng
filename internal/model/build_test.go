// Package model defines the core data types for arsenal-ng.
package model

import (
	"testing"
)

// =============================================================================
// Command Building Tests
// =============================================================================

// TestBuildCommand verifies placeholder substitution and single-pass behavior.
func TestBuildCommand(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
		args []Argument
		want string
	}{
		{"plain substitution", "nmap {{ip}}", []Argument{{Name: "ip", Value: "1.2.3.4"}}, "nmap 1.2.3.4"},
		{"single pass no recursion", "run {{a}}", []Argument{{Name: "a", Value: "{{b}}"}}, "run {{b}}"},
		{"unknown placeholder kept", "nmap {{ip}}", nil, "nmap {{ip}}"},
		{"default value used", "ffuf {{w|admin}}", []Argument{{Name: "w", Value: "admin"}}, "ffuf admin"},
		{"no placeholders", "ls -la", []Argument{{Name: "x", Value: "y"}}, "ls -la"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := BuildCommand(tc.cmd, tc.args); got != tc.want {
				t.Fatalf("BuildCommand(%q) = %q, want %q", tc.cmd, got, tc.want)
			}
		})
	}
}
