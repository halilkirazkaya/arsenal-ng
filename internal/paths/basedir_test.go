// Package paths resolves OS-specific configuration directories. It follows
// the XDG Base Directory Specification on Linux, uses ~/Library/Application
// Support on macOS, and ~/.config elsewhere.
package paths

import (
	"path/filepath"
	"strings"
	"testing"
)

// =============================================================================
// Platform Matrix
// =============================================================================

// TestBaseDirLinuxXDG verifies an explicit XDG override wins on Linux.
func TestBaseDirLinuxXDG(t *testing.T) {
	isolate(t)
	t.Setenv("XDG_CONFIG_HOME", "/custom/xdg")
	got, err := baseDir("linux")
	if err != nil {
		t.Fatal(err)
	}
	if got != "/custom/xdg" {
		t.Fatalf("linux xdg = %q", got)
	}
}

// TestBaseDirLinuxDefault verifies empty XDG falls back to ~/.config.
func TestBaseDirLinuxDefault(t *testing.T) {
	tmp := isolate(t)
	t.Setenv("XDG_CONFIG_HOME", "")
	got, err := baseDir("linux")
	if err != nil {
		t.Fatal(err)
	}
	if want := filepath.Join(tmp, ".config"); got != want {
		t.Fatalf("linux default = %q, want %q", got, want)
	}
}

// TestBaseDirDarwin verifies macOS resolves under Application Support.
func TestBaseDirDarwin(t *testing.T) {
	tmp := isolate(t)
	got, err := baseDir("darwin")
	if err != nil {
		t.Fatal(err)
	}
	if want := filepath.Join(tmp, "Library", "Application Support"); got != want {
		t.Fatalf("darwin = %q, want %q", got, want)
	}
}

// TestBaseDirOtherUnix verifies unknown systems use the ~/.config layout.
func TestBaseDirOtherUnix(t *testing.T) {
	tmp := isolate(t)
	got, err := baseDir("freebsd")
	if err != nil {
		t.Fatal(err)
	}
	if want := filepath.Join(tmp, ".config"); got != want {
		t.Fatalf("other = %q, want %q", got, want)
	}
	if !strings.HasPrefix(got, tmp) {
		t.Fatalf("base = %q, want under HOME", got)
	}
}

// TestBaseDirFailsWithoutHome verifies missing HOME fails every platform.
func TestBaseDirFailsWithoutHome(t *testing.T) {
	isolate(t)
	t.Setenv("HOME", "")
	t.Setenv("XDG_CONFIG_HOME", "")
	for _, goos := range []string{"linux", "darwin", "freebsd"} {
		if _, err := baseDir(goos); err == nil {
			t.Fatalf("%s without HOME = nil error", goos)
		}
	}
	if _, err := ConfigDir(); err == nil {
		t.Fatal("ConfigDir without HOME = nil error")
	}
}
