// Package loader reads cheat files from the embedded filesystem and
// converts them into Cheat objects.
package loader

import (
	"io"
	"io/fs"
	"log"
	"os"
	"testing"
	"testing/fstest"
)

// =============================================================================
// Filesystem Walk Tests
// =============================================================================

// quiet silences log output for the duration of a test.
func quiet(t *testing.T) {
	t.Helper()
	log.SetOutput(io.Discard)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })
}

// TestLoadFromFS verifies valid files load while others are skipped.
func TestLoadFromFS(t *testing.T) {
	quiet(t)
	fsys := fstest.MapFS{
		"cheats/ok.yaml":  {Data: []byte("tool: t\ntags: [a]\nactions:\n  - title: hi\n    command: run\n")},
		"cheats/skip.md":  {Data: []byte("not yaml")},
		"cheats/bad.yaml": {Data: []byte("{broken")},
	}
	cheats, err := loadFromFS(fsys, "cheats")
	if err != nil {
		t.Fatal(err)
	}
	// Valid file loads, non-YAML skipped, broken YAML skipped.
	if len(cheats) != 1 || cheats[0].Tool != "t" {
		t.Fatalf("cheats = %+v, want 1 valid", cheats)
	}
}

// TestLoadFromFSMissingRoot verifies a missing root returns an error.
func TestLoadFromFSMissingRoot(t *testing.T) {
	quiet(t)
	fsys := fstest.MapFS{}
	if _, err := loadFromFS(fsys, "nope"); err == nil {
		t.Fatal("missing root = nil error")
	}
}

// TestLoadFromFSWalkError verifies filesystem walk failures abort the load.
func TestLoadFromFSWalkError(t *testing.T) {
	quiet(t)
	// A MapFS that fails reads surfaces the walk error.
	fsys := failFS{}
	if _, err := loadFromFS(fsys, "."); err == nil {
		t.Fatal("failing fs = nil error")
	}
}

// failFS always fails directory reads.
// It forces the walk-error branch of loadFromFS.
type failFS struct{}

// Open always returns a permission error for any path.
func (failFS) Open(name string) (fs.File, error) {
	return nil, fs.ErrPermission
}
