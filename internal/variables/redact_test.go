// Package variables provides a thread-safe global variables store that
// persists across sessions via a JSON file. Variables are saved atomically.
package variables

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

// =============================================================================
// Secret Redaction
// =============================================================================

// TestSetRedactsValue verifies values never reach the debug log.
func TestSetRedactsValue(t *testing.T) {
	isolateHome(t)
	s := mustNew(t)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	if err := s.Set("password", "s3cr3t-value"); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(buf.String(), "s3cr3t-value") {
		t.Fatalf("log leaks value: %q", buf.String())
	}
	if !strings.Contains(buf.String(), "password") {
		t.Fatalf("log missing name: %q", buf.String())
	}
}

// =============================================================================
// Empty Fallback
// =============================================================================

// TestEmptyStoreUsable verifies the fallback store never panics on write.
func TestEmptyStoreUsable(t *testing.T) {
	s := Empty()
	if err := s.Set("k", "v"); err != nil {
		t.Fatalf("Empty.Set = %v", err)
	}
	if got, ok := s.Get("k"); !ok || got != "v" {
		t.Fatalf("Empty.Get = %q, %v", got, ok)
	}
	if ok, err := s.Unset("k"); !ok || err != nil {
		t.Fatalf("Empty.Unset = %v, %v", ok, err)
	}
	if s.Count() != 0 {
		t.Fatalf("Empty.Count = %d", s.Count())
	}
}
