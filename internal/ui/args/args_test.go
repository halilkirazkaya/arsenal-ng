// Package args provides the argument input view: a form where users fill
// in command arguments before execution.
package args

import (
	"io"
	"log"
	"os"
	"testing"

	"charm.land/lipgloss/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/theme"
	"github.com/halilkirazkaya/arsenal-ng/internal/variables"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// TestMain silences logs and pins deterministic theme colors for tests.
func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	theme.SetLogoColor(lipgloss.Color("#FFFFFF"))
	os.Exit(m.Run())
}

// testStore builds an isolated variable store backed by a temp dir.
func testStore(t *testing.T, kvs map[string]string) *variables.Store {
	t.Helper()
	tmp := t.TempDir()
	t.Setenv("HOME", tmp)
	t.Setenv("XDG_CONFIG_HOME", tmp)
	g, err := variables.New()
	if err != nil {
		t.Fatalf("variables.New: %v", err)
	}
	for k, v := range kvs {
		if err := g.Set(k, v); err != nil {
			t.Fatalf("Set(%s): %v", k, err)
		}
	}
	return g
}

// plainCheat returns a cheat with one required and one defaulted arg.
func plainCheat() *model.Cheat {
	return &model.Cheat{Tool: "nmap", Title: "Scan", Command: "nmap -sV {{ip}} {{port|80}}"}
}

// =============================================================================
// Prefill Order
// =============================================================================

// TestNewGlobalBeatsDefault verifies globals win over argument defaults.
func TestNewGlobalBeatsDefault(t *testing.T) {
	g := testStore(t, map[string]string{"port": "9999"})
	m := New(plainCheat(), g)
	if m.Args[1].Value != "9999" {
		t.Fatalf("port = %q, want global 9999", m.Args[1].Value)
	}
}

// TestNewDefaultWhenNoGlobal verifies defaults fill in without globals.
func TestNewDefaultWhenNoGlobal(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	if m.Args[1].Value != "80" {
		t.Fatalf("port = %q, want default 80", m.Args[1].Value)
	}
	if m.Args[0].Value != "" {
		t.Fatalf("ip = %q, want empty", m.Args[0].Value)
	}
}

// TestNewFocusFirstEmpty verifies focus lands on the first empty field.
func TestNewFocusFirstEmpty(t *testing.T) {
	g := testStore(t, nil)
	m := New(plainCheat(), g)
	if m.Cursor != 0 {
		t.Fatalf("Cursor = %d, want 0 (first empty)", m.Cursor)
	}
	if !m.Inputs[0].Focused() {
		t.Fatal("first empty input should be focused")
	}
}

// TestNewFocusFirstWhenAllFilled verifies focus falls back to field zero.
func TestNewFocusFirstWhenAllFilled(t *testing.T) {
	g := testStore(t, map[string]string{"ip": "1.2.3.4"})
	m := New(plainCheat(), g)
	if m.Cursor != 0 || !m.Inputs[0].Focused() {
		t.Fatalf("all filled: cursor=%d focused=%v, want 0/true", m.Cursor, m.Inputs[0].Focused())
	}
}
