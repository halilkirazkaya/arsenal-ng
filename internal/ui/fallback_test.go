// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"os"
	"path/filepath"
	"testing"
)

// =============================================================================
// Fallback Tests
// =============================================================================

// TestNewFallsBackToEmptyStore verifies an unwritable HOME still boots.
func TestNewFallsBackToEmptyStore(t *testing.T) {
	tmp := t.TempDir()
	if err := os.Chmod(tmp, 0o555); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(tmp, 0o755) })
	t.Setenv("HOME", tmp)

	a := New(nil)
	if a.globals == nil || a.globals.Count() != 0 {
		t.Fatal("fallback store must be empty and usable")
	}
	if a.view != viewSearch {
		t.Fatalf("view = %v, want search", a.view)
	}
}

// TestUpdateUnknownView verifies unknown views stay put with no command.
func TestUpdateUnknownView(t *testing.T) {
	a := newTestApp(t, nil)
	a.view = view(99)
	m, cmd := a.Update(keyRune("x"))
	if m.(App).view != view(99) || cmd != nil {
		t.Fatal("unknown view must stay with nil cmd")
	}
	if a.render() != "" {
		t.Fatal("unknown view must render empty")
	}
}

// TestSetUnsetDiskError verifies failed saves surface as error statuses.
func TestSetUnsetDiskError(t *testing.T) {
	a := newTestApp(t, nil)
	if err := a.globals.Set("k", "v"); err != nil {
		t.Fatal(err)
	}
	// Freeze the config dir so subsequent saves fail.
	cfg := filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "arsenal-ng")
	if err := os.Chmod(cfg, 0o555); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(cfg, 0o755) })

	got := a.applySet("set k2=v2")
	if !got.search.StatusIsError {
		t.Fatal("failed save must set error status")
	}
	got = got.applyUnset("unset k")
	if !got.search.StatusIsError {
		t.Fatal("failed unset save must set error status")
	}
}
