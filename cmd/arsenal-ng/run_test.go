// Package main is the entry point for arsenal-ng, a modern pentest command launcher.
package main

import (
	"errors"
	"strings"
	"testing"
)

// =============================================================================
// Load Failures
// =============================================================================

// TestRunLoadError verifies a loader failure aborts the run.
func TestRunLoadError(t *testing.T) {
	quiet(t)
	s := &stub{loadErr: errors.New("boom")}
	if err := run(s); err == nil || !strings.Contains(err.Error(), "load cheats") {
		t.Fatalf("run = %v, want load error", err)
	}
}

// TestRunNoCheats verifies an empty cheat list is an error.
func TestRunNoCheats(t *testing.T) {
	quiet(t)
	s := &stub{cheats: nil}
	if err := run(s); err == nil || !strings.Contains(err.Error(), "no cheats") {
		t.Fatalf("run = %v, want no-cheats error", err)
	}
}
