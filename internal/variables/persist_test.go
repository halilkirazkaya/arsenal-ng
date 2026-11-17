// Package variables provides a thread-safe global variables store that
// persists across sessions via a JSON file. Variables are saved atomically.
package variables

import (
	"os"
	"testing"
)

// =============================================================================
// Persistence Tests
// =============================================================================

// TestRoundTrip verifies variables survive a reload and unset persists too.
func TestRoundTrip(t *testing.T) {
	isolateHome(t)
	s := mustNew(t)
	if err := s.Set("user", "admin"); err != nil {
		t.Fatal(err)
	}
	if err := s.Set("pw", "s3cret"); err != nil {
		t.Fatal(err)
	}

	reloaded, err := New()
	if err != nil {
		t.Fatalf("New reload = %v", err)
	}
	if reloaded.Count() != 2 {
		t.Fatalf("reloaded Count = %d, want 2", reloaded.Count())
	}
	if got, _ := reloaded.Get("pw"); got != "s3cret" {
		t.Fatalf("reloaded pw = %q", got)
	}

	// Unset persists to disk too.
	if _, err := reloaded.Unset("pw"); err != nil {
		t.Fatal(err)
	}
	third, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := third.Get("pw"); ok {
		t.Fatal("unset variable survived reload")
	}
}

// TestCorruptAndNullJSON verifies corrupt files error while null stays usable.
func TestCorruptAndNullJSON(t *testing.T) {
	isolateHome(t)
	s := mustNew(t)
	path := s.filePath

	if err := os.WriteFile(path, []byte("{not json"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := New(); err == nil {
		t.Fatal("New with corrupt JSON = nil error")
	}

	if err := os.WriteFile(path, []byte("null"), 0600); err != nil {
		t.Fatal(err)
	}
	nullStore, err := New()
	if err != nil {
		t.Fatalf("New with null JSON = %v", err)
	}
	if err := nullStore.Set("k", "v"); err != nil {
		t.Fatalf("Set after null JSON = %v (nil map panic?)", err)
	}

	if err := os.WriteFile(path, []byte(""), 0600); err != nil {
		t.Fatal(err)
	}
	empty, err := New()
	if err != nil {
		t.Fatalf("New with empty file = %v", err)
	}
	if empty.Count() != 0 {
		t.Fatalf("empty file Count = %d", empty.Count())
	}
}
