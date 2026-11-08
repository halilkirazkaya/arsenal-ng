// Package model defines the core data types for arsenal-ng.
package model

import (
	"reflect"
	"testing"
)

// =============================================================================
// Tool Summary Tests
// =============================================================================

// TestToolSummary verifies sorted unique tool extraction with counts.
func TestToolSummary(t *testing.T) {
	cheats := []*Cheat{
		{Tool: "nmap"}, {Tool: "ffuf"}, {Tool: "nmap"},
		{Tool: ""},
	}
	got := ToolSummary(cheats)
	want := []ToolInfo{{Name: "ffuf", CommandCount: 1}, {Name: "nmap", CommandCount: 2}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ToolSummary = %+v, want %+v", got, want)
	}
	if got := ToolSummary(nil); len(got) != 0 {
		t.Fatalf("ToolSummary(nil) = %+v, want empty", got)
	}
}
