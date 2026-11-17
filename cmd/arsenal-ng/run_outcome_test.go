// Package main is the entry point for arsenal-ng, a modern pentest command launcher.
package main

import (
	"errors"
	"strings"
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui"
)

// =============================================================================
// Interact Outcomes
// =============================================================================

// TestRunInteractError verifies a TUI failure aborts the run.
func TestRunInteractError(t *testing.T) {
	quiet(t)
	cheat := &model.Cheat{Tool: "nmap", Command: "nmap x"}
	s := &stub{cheats: []*model.Cheat{cheat}, interactErr: errors.New("tty gone")}
	if err := run(s); err == nil || !strings.Contains(err.Error(), "TUI error") {
		t.Fatalf("run = %v, want TUI error", err)
	}
}

// TestRunCancelled verifies cancellation skips the prefill step.
func TestRunCancelled(t *testing.T) {
	quiet(t)
	cheat := &model.Cheat{Tool: "nmap", Command: "nmap x"}
	s := &stub{cheats: []*model.Cheat{cheat}, app: ui.App{Cancelled: true}}
	if err := run(s); err != nil {
		t.Fatalf("run = %v", err)
	}
	if s.prefilled != "" {
		t.Fatalf("cancelled run prefilled %q", s.prefilled)
	}
}

// TestRunNoCommand verifies an empty selection skips the prefill step.
func TestRunNoCommand(t *testing.T) {
	quiet(t)
	cheat := &model.Cheat{Tool: "nmap", Command: "nmap x"}
	s := &stub{cheats: []*model.Cheat{cheat}, app: ui.App{}}
	if err := run(s); err != nil {
		t.Fatalf("run = %v", err)
	}
	if s.prefilled != "" {
		t.Fatalf("empty run prefilled %q", s.prefilled)
	}
}

// =============================================================================
// Prefill And Logging
// =============================================================================

// TestRunPrefillErrorWarns verifies prefill failures only warn, not fail.
func TestRunPrefillErrorWarns(t *testing.T) {
	quiet(t)
	cheat := &model.Cheat{Tool: "nmap", Command: "nmap x"}
	s := &stub{cheats: []*model.Cheat{cheat}, app: ui.App{FinalCommand: "nmap x"},
		prefillErr: errors.New("EIO")}
	if err := run(s); err != nil {
		t.Fatalf("run = %v, prefill error must not fail", err)
	}
	joined := strings.Join(s.warnings, "\n")
	if !strings.Contains(joined, "could not prefill") {
		t.Fatalf("warnings = %q, want prefill hint", joined)
	}
	if !strings.Contains(joined, "nmap x") {
		t.Fatalf("warnings = %q, want fallback command", joined)
	}
}

// TestRunLoggingOptional verifies broken logging only warns and continues.
func TestRunLoggingOptional(t *testing.T) {
	quiet(t)
	cheat := &model.Cheat{Tool: "nmap", Command: "nmap x"}
	for name, s := range map[string]*stub{
		"log path error": {logErr: errors.New("no home"), cheats: []*model.Cheat{cheat},
			app: ui.App{FinalCommand: "nmap x"}},
		"setup error": {setupErr: errors.New("no file"), cheats: []*model.Cheat{cheat},
			app: ui.App{FinalCommand: "nmap x"}},
	} {
		if err := run(s); err != nil {
			t.Fatalf("%s: run = %v, logging must be optional", name, err)
		}
		if s.prefilled != "nmap x" {
			t.Fatalf("%s: prefilled = %q", name, s.prefilled)
		}
		if len(s.warnings) == 0 {
			t.Fatalf("%s: expected a warning", name)
		}
	}
}
