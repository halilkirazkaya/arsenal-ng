// Package main is the entry point for arsenal-ng, a modern pentest command launcher.
package main

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui"
)

// =============================================================================
// Success Path
// =============================================================================

// TestRunSuccess verifies a full run prefills and closes the log.
func TestRunSuccess(t *testing.T) {
	quiet(t)
	cheat := &model.Cheat{Tool: "nmap", Command: "nmap x"}
	s := &stub{cheats: []*model.Cheat{cheat}, app: ui.App{FinalCommand: "nmap x"}}
	if err := run(s); err != nil {
		t.Fatalf("run = %v", err)
	}
	if s.prefilled != "nmap x" {
		t.Fatalf("prefilled = %q", s.prefilled)
	}
	if !s.closed {
		t.Fatal("log file not closed")
	}
}
