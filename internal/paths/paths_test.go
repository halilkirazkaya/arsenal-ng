// Package paths resolves OS-specific configuration directories. It follows
// the XDG Base Directory Specification on Linux, uses ~/Library/Application
// Support on macOS, and ~/.config elsewhere.
package paths

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// isolate redirects HOME and XDG paths into a temp dir for the test.
func isolate(t *testing.T) string {
	t.Helper()
	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(tmp, ".config"))
	return tmp
}

// =============================================================================
// Directory Resolution
// =============================================================================

// TestConfigDirCreatesDir verifies ConfigDir creates a 0700 app directory.
func TestConfigDirCreatesDir(t *testing.T) {
	tmp := isolate(t)
	dir, err := ConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(dir, tmp) {
		t.Fatalf("dir = %q, want under %q", dir, tmp)
	}
	if !strings.HasSuffix(dir, config.AppName) {
		t.Fatalf("dir = %q, want app suffix", dir)
	}
	if st, err := os.Stat(dir); err != nil || !st.IsDir() {
		t.Fatalf("dir not created: %v", err)
	}
	if st, _ := os.Stat(dir); st.Mode().Perm() != 0o700 {
		t.Fatalf("mode = %o, want 700", st.Mode().Perm())
	}
}

// TestVariablesAndLogPaths verifies sibling file names share one directory.
func TestVariablesAndLogPaths(t *testing.T) {
	isolate(t)
	variables, err := Variables()
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Base(variables) != "variables.json" {
		t.Fatalf("variables = %q", variables)
	}
	log, err := Log()
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Base(log) != "debug.log" {
		t.Fatalf("log = %q", log)
	}
	if filepath.Dir(variables) != filepath.Dir(log) {
		t.Fatal("variables and log must share a directory")
	}
}

// TestConfigDirFailsWhenHomeReadonly verifies unreadable HOME surfaces errors.
func TestConfigDirFailsWhenHomeReadonly(t *testing.T) {
	isolate(t)
	tmp := t.TempDir()
	if err := os.Chmod(tmp, 0o555); err != nil {
		t.Fatal(err)
	}
	// Restore write access so temp cleanup can remove the directory.
	t.Cleanup(func() { _ = os.Chmod(tmp, 0o755) })
	t.Setenv("HOME", tmp)

	if _, err := ConfigDir(); err == nil {
		t.Fatal("read-only HOME = nil error")
	}
	if _, err := Variables(); err == nil {
		t.Fatal("Variables with bad HOME = nil error")
	}
	if _, err := Log(); err == nil {
		t.Fatal("Log with bad HOME = nil error")
	}
}
