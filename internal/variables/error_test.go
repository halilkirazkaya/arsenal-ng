// Package variables provides a thread-safe global variables store that
// persists across sessions via a JSON file. Variables are saved atomically.
package variables

import (
	"os"
	"path/filepath"
	"testing"
)

// =============================================================================
// Error Path Tests
// =============================================================================

// TestSaveErrorReadonlyDir verifies disk failures surface while memory updates.
func TestSaveErrorReadonlyDir(t *testing.T) {
	isolateHome(t)
	s := mustNew(t)

	// Make the config dir read-only so the atomic write fails.
	dir := filepath.Dir(s.filePath)
	if err := os.Chmod(dir, 0o555); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(dir, 0o755) })

	if err := s.Set("k", "v"); err == nil {
		t.Fatal("Set on read-only dir = nil error")
	}
	// Memory still updated even though disk save failed.
	if got, ok := s.Get("k"); !ok || got != "v" {
		t.Fatalf("Get after failed save = %q, %v", got, ok)
	}

	if _, err := s.Unset("missing"); err != nil {
		t.Fatalf("Unset missing on read-only dir = %v", err)
	}
	if ok, err := s.Unset("k"); !ok || err == nil {
		t.Fatalf("Unset on read-only dir = %v, %v; want removed+error", ok, err)
	}
}

// TestNewFailsWhenConfigDirUnwritable verifies New errors without a config dir.
func TestNewFailsWhenConfigDirUnwritable(t *testing.T) {
	isolateHome(t)
	// HOME itself read-only so no config dir can be made. Linux prefers
	// XDG_CONFIG_HOME, so point it inside the read-only dir as well.
	tmp := t.TempDir()
	if err := os.Chmod(tmp, 0o555); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(tmp, 0o755) })
	t.Setenv("HOME", tmp)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(tmp, ".config"))

	if _, err := New(); err == nil {
		t.Fatal("New with unwritable config dir = nil error")
	}
}

// TestLoadFallbackEmptyPath verifies loading with no path is a no-op success.
func TestLoadFallbackEmptyPath(t *testing.T) {
	s := &Store{variables: map[string]string{}, filePath: ""}
	if err := s.loadFromFile(); err != nil {
		t.Fatalf("load fallback = %v", err)
	}
}

// TestSaveFailsWhenTargetIsDir verifies rename over a dir reports an error.
func TestSaveFailsWhenTargetIsDir(t *testing.T) {
	isolateHome(t)
	s := mustNew(t)
	// Point the store at an existing directory: tmp writes fine,
	// but renaming over a non-empty dir fails.
	dir := t.TempDir()
	sub := filepath.Join(dir, "sub")
	if err := os.Mkdir(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	s.filePath = dir
	if err := s.Set("k", "v"); err == nil {
		t.Fatal("Set with dir target = nil error")
	}
}
