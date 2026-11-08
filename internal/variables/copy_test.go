// Package variables provides a thread-safe global variables store that
// persists across sessions via a JSON file. Variables are saved atomically.
package variables

import (
	"testing"
)

// =============================================================================
// Isolation Tests
// =============================================================================

// TestAllReturnsCopy verifies All returns a defensive copy, not internals.
func TestAllReturnsCopy(t *testing.T) {
	isolateHome(t)
	s := mustNew(t)
	if err := s.Set("a", "1"); err != nil {
		t.Fatal(err)
	}
	all := s.All()
	if all["a"] != "1" {
		t.Fatalf("All = %v", all)
	}
	all["a"] = "mutated"
	if got, _ := s.Get("a"); got != "1" {
		t.Fatalf("All exposed internals: %q", got)
	}
}

// TestEmptyFilePathFallback verifies CRUD works with persistence disabled.
func TestEmptyFilePathFallback(t *testing.T) {
	isolateHome(t)
	s := &Store{variables: map[string]string{}, filePath: ""}
	if err := s.Set("k", "v"); err != nil {
		t.Fatalf("Set fallback = %v", err)
	}
	if got, ok := s.Get("k"); !ok || got != "v" {
		t.Fatalf("Get fallback = %q, %v", got, ok)
	}
	if ok, err := s.Unset("k"); !ok || err != nil {
		t.Fatalf("Unset fallback = %v, %v", ok, err)
	}
	if s.Count() != 0 {
		t.Fatalf("Count fallback = %d", s.Count())
	}
}
