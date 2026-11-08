// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"testing"
)

// =============================================================================
// Form Submission
// =============================================================================

// TestSubmitBuildsCommand verifies submit assembles the final command.
func TestSubmitBuildsCommand(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	m.Inputs[0].SetValue("10.0.0.1")
	m.Inputs[1].SetValue("443")
	_, _, oc, cmd := m.submit()
	if oc != Submit || cmd != "nmap -sV 10.0.0.1 443" {
		t.Fatalf("submit = %v/%q", oc, cmd)
	}
}

// TestSubmitBlockedWhenRequiredEmpty verifies empty required args block.
func TestSubmitBlockedWhenRequiredEmpty(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	// Required ip stays empty so submission must return no command.
	m.Inputs[0].SetValue("")
	_, _, oc, cmd := m.submit()
	if oc != Submit || cmd != "" {
		t.Fatalf("blocked submit = %v/%q, want Submit/empty", oc, cmd)
	}
}
