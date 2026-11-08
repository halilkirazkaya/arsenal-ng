// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"strings"
	"testing"
)

// =============================================================================
// Set Flow Tests
// =============================================================================

// TestFlowSetVar verifies a valid set command stores the variable.
func TestFlowSetVar(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("set ip=1.2.3.4")
	m, _ := a.Update(keyEnter())
	got := m.(App)
	v, ok := got.globals.Get("ip")
	if !ok || v != "1.2.3.4" {
		t.Fatalf("var not stored: %v %v", v, ok)
	}
	if !strings.Contains(got.search.StatusMsg, "Set ip = 1.2.3.4") {
		t.Fatalf("status = %q", got.search.StatusMsg)
	}
	if got.search.StatusIsError {
		t.Fatal("success must not be error")
	}
}

// TestFlowSetVarInvalidFormat verifies malformed set commands show usage.
func TestFlowSetVarInvalidFormat(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("set novalue")
	m, _ := a.Update(keyEnter())
	got := m.(App)
	if !got.search.StatusIsError || !strings.Contains(got.search.StatusMsg, "Usage") {
		t.Fatalf("status = %q", got.search.StatusMsg)
	}
}

// =============================================================================
// Unset Flow Tests
// =============================================================================

// TestFlowUnsetVar verifies a valid unset command removes the variable.
func TestFlowUnsetVar(t *testing.T) {
	a := newTestApp(t, nil)
	if err := a.globals.Set("ip", "1.2.3.4"); err != nil {
		t.Fatal(err)
	}
	a.search = a.search.SetQuery("unset ip")
	m, _ := a.Update(keyEnter())
	got := m.(App)
	if _, ok := got.globals.Get("ip"); ok {
		t.Fatal("var must be removed")
	}
	if !strings.Contains(got.search.StatusMsg, "Unset ip") {
		t.Fatalf("status = %q", got.search.StatusMsg)
	}
}

// TestFlowUnsetMissing verifies unsetting an unknown variable errors.
func TestFlowUnsetMissing(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("unset ghost")
	m, _ := a.Update(keyEnter())
	got := m.(App)
	if !got.search.StatusIsError || !strings.Contains(got.search.StatusMsg, "not found") {
		t.Fatalf("status = %q", got.search.StatusMsg)
	}
}
