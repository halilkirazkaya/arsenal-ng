// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"strings"
	"testing"
)

// =============================================================================
// Apply Set Tests
// =============================================================================

// TestApplySetSuccessClearsQuery verifies set stores, clears, and resets.
func TestApplySetSuccessClearsQuery(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("set ip=1.1.1.1 extra")
	got := a.applySet("set ip = 10.0.0.1")
	v, ok := got.globals.Get("ip")
	if !ok || v != "10.0.0.1" {
		t.Fatalf("stored = %q,%v", v, ok)
	}
	if !strings.Contains(got.search.StatusMsg, "Set ip = 10.0.0.1") {
		t.Fatalf("status = %q", got.search.StatusMsg)
	}
	if got.search.Input.Value() != "" {
		t.Fatal("query must be cleared after set")
	}
	if len(got.search.Filtered) != 2 {
		t.Fatal("list must reset to full cheats after set")
	}
}

// TestApplySetRejectsEmpty verifies blank names and values are rejected.
func TestApplySetRejectsEmpty(t *testing.T) {
	for _, q := range []string{"set =x", "set name=", "set "} {
		a := newTestApp(t, nil)
		got := a.applySet(q)
		if !got.search.StatusIsError {
			t.Fatalf("%q: must be error", q)
		}
		if _, ok := got.globals.Get("name"); ok {
			t.Fatalf("%q: must not store", q)
		}
	}
}

// =============================================================================
// Apply Unset Tests
// =============================================================================

// TestApplyUnsetSuccess verifies unset removes an existing variable.
func TestApplyUnsetSuccess(t *testing.T) {
	a := newTestApp(t, nil)
	if err := a.globals.Set("ip", "1.1.1.1"); err != nil {
		t.Fatal(err)
	}
	got := a.applyUnset("unset ip")
	if _, ok := got.globals.Get("ip"); ok {
		t.Fatal("var must be gone")
	}
	if !strings.Contains(got.search.StatusMsg, "Unset ip") {
		t.Fatalf("status = %q", got.search.StatusMsg)
	}
}

// TestApplyUnsetMissingAndInvalid verifies missing and malformed unsets error.
func TestApplyUnsetMissingAndInvalid(t *testing.T) {
	a := newTestApp(t, nil)
	missing := a.applyUnset("unset ghost")
	if !missing.search.StatusIsError || !strings.Contains(missing.search.StatusMsg, "not found") {
		t.Fatalf("missing status = %q", missing.search.StatusMsg)
	}
	invalid := a.applyUnset("unset ")
	if !invalid.search.StatusIsError || !strings.Contains(invalid.search.StatusMsg, "Usage") {
		t.Fatalf("invalid status = %q", invalid.search.StatusMsg)
	}
}
