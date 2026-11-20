// Package language persists the UI language preference across sessions
// via a JSON file. The preference is saved atomically like variables.
package language

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/paths"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// isolate redirects HOME and XDG paths into a temp dir for the test.
func isolate(t *testing.T) {
	t.Helper()
	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	t.Setenv("XDG_CONFIG_HOME", tmp)
	log.SetOutput(io.Discard)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })
}

// =============================================================================
// Store Tests
// =============================================================================

// TestNewDefaultsToEnglish verifies first run yields the default language.
func TestNewDefaultsToEnglish(t *testing.T) {
	isolate(t)
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if got := g.Get(); got != config.DefaultLanguage {
		t.Fatalf("language = %q, want %q", got, config.DefaultLanguage)
	}
}

// TestSetPersistsAcrossInstances verifies the choice survives reloads.
func TestSetPersistsAcrossInstances(t *testing.T) {
	isolate(t)
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if err := g.Set("TR"); err != nil {
		t.Fatal(err)
	}
	reloaded, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if got := reloaded.Get(); got != "tr" {
		t.Fatalf("reloaded = %q, want tr", got)
	}
}

// TestSetRejectsUnsupported verifies bad codes error without changing state.
func TestSetRejectsUnsupported(t *testing.T) {
	isolate(t)
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if err := g.Set("xx"); err == nil {
		t.Fatal("unsupported code = nil error")
	}
	if got := g.Get(); got != config.DefaultLanguage {
		t.Fatalf("language = %q after failed set", got)
	}
}

// TestLoadIgnoresUnsupportedSaved verifies unknown saved codes fall back.
func TestLoadIgnoresUnsupportedSaved(t *testing.T) {
	isolate(t)
	path, err := paths.Language()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(`{"language": "de"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if got := g.Get(); got != config.DefaultLanguage {
		t.Fatalf("language = %q, want default", got)
	}
}

// TestLoadRejectsCorruptFile verifies invalid JSON surfaces an error.
func TestLoadRejectsCorruptFile(t *testing.T) {
	isolate(t)
	path, err := paths.Language()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("{nope"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := New(); err == nil {
		t.Fatal("corrupt file = nil error")
	}
}

// TestEmptyDisablesPersistence verifies the fallback store stays in memory.
func TestEmptyDisablesPersistence(t *testing.T) {
	g := Empty()
	if err := g.Set("tr"); err != nil {
		t.Fatal(err)
	}
	if got := g.Get(); got != "tr" {
		t.Fatalf("language = %q, want tr", got)
	}
}
