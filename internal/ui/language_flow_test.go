// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"strings"
	"testing"
)

// =============================================================================
// Language Flow Tests
// =============================================================================

// TestFlowSetLanguage verifies set language= switches translations.
func TestFlowSetLanguage(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("set language=tr")
	m, _ := a.Update(keyEnter())
	got := m.(App)
	if got.lang != "tr" || got.search.Lang() != "tr" {
		t.Fatalf("lang = %q/%q, want tr/tr", got.lang, got.search.Lang())
	}
	if !strings.Contains(got.search.StatusMsg, "Türkçe (tr)") {
		t.Fatalf("status = %q", got.search.StatusMsg)
	}
	if got.search.StatusIsError {
		t.Fatal("success must not be error")
	}
	if _, ok := got.globals.Get("language"); ok {
		t.Fatal("language must not leak into variables")
	}
}

// TestFlowSetLanguageInvalid verifies unsupported codes error cleanly.
func TestFlowSetLanguageInvalid(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("set language=xx")
	m, _ := a.Update(keyEnter())
	got := m.(App)
	if !got.search.StatusIsError || !strings.Contains(got.search.StatusMsg, "unsupported") {
		t.Fatalf("status = %q", got.search.StatusMsg)
	}
	if got.lang != "en" {
		t.Fatalf("lang = %q, want unchanged en", got.lang)
	}
}

// TestFlowSetLanguagePersists verifies the choice survives restarts.
func TestFlowSetLanguagePersists(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("set language=tr")
	if _, _ = a.Update(keyEnter()); a.lang != "en" {
		t.Fatal("Update must not mutate the sender")
	}
	restarted := New(fixtureCheats())
	if restarted.lang != "tr" || restarted.search.Lang() != "tr" {
		t.Fatalf("restarted lang = %q/%q, want tr/tr", restarted.lang, restarted.search.Lang())
	}
}
