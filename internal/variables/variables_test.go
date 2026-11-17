// Package variables provides a thread-safe global variables store that
// persists across sessions via a JSON file. Variables are saved atomically.
package variables

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// isolateHome redirects HOME and config paths to a temp dir for the test.
// Log output is silenced so passing tests stay quiet.
func isolateHome(t *testing.T) {
	t.Helper()
	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(tmp, ".config"))
	log.SetOutput(io.Discard)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })
}

// mustNew creates a fresh empty store, failing the test on error.
// The store starts empty because isolateHome points at a temp dir.
func mustNew(t *testing.T) *Store {
	t.Helper()
	s, err := New()
	if err != nil {
		t.Fatalf("New() = %v", err)
	}
	if s.Count() != 0 {
		t.Fatalf("fresh store Count = %d, want 0", s.Count())
	}
	return s
}

// =============================================================================
// CRUD Tests
// =============================================================================

// TestSetGetUnset verifies the full set, overwrite, and unset lifecycle.
func TestSetGetUnset(t *testing.T) {
	isolateHome(t)
	s := mustNew(t)

	if err := s.Set("target", "10.0.0.1"); err != nil {
		t.Fatalf("Set = %v", err)
	}
	if got, ok := s.Get("target"); !ok || got != "10.0.0.1" {
		t.Fatalf("Get = %q, %v", got, ok)
	}
	if s.Count() != 1 {
		t.Fatalf("Count = %d, want 1", s.Count())
	}

	// Overwrite keeps count at 1.
	if err := s.Set("target", "10.0.0.2"); err != nil {
		t.Fatalf("Set overwrite = %v", err)
	}
	if got, _ := s.Get("target"); got != "10.0.0.2" {
		t.Fatalf("Get after overwrite = %q", got)
	}
	if s.Count() != 1 {
		t.Fatalf("Count after overwrite = %d, want 1", s.Count())
	}

	ok, err := s.Unset("target")
	if !ok || err != nil {
		t.Fatalf("Unset = %v, %v", ok, err)
	}
	if _, ok := s.Get("target"); ok {
		t.Fatal("Get after Unset found deleted key")
	}

	ok, err = s.Unset("missing")
	if ok || err != nil {
		t.Fatalf("Unset missing = %v, %v", ok, err)
	}
	if _, ok := s.Get("missing"); ok {
		t.Fatal("Get missing = found")
	}
}
