// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"testing"
)

// =============================================================================
// Transition Tests
// =============================================================================

// TestShowVariablesResetsQuery verifies entering variables clears the query.
func TestShowVariablesResetsQuery(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("nmap")
	got := a.showVariables()
	if got.view != viewVariables {
		t.Fatal("must enter variables view")
	}
	if got.search.Input.Value() != "" {
		t.Fatal("transition must clear search query")
	}
}

// TestShowToolsResetsQuery verifies entering tools clears the query.
func TestShowToolsResetsQuery(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("nmap")
	got := a.showTools()
	if got.view != viewTools {
		t.Fatal("must enter tools view")
	}
	if got.search.Input.Value() != "" {
		t.Fatal("transition must clear search query")
	}
}

// TestShowHelpResetsQuery verifies entering help clears the query.
func TestShowHelpResetsQuery(t *testing.T) {
	a := newTestApp(t, nil)
	a.search = a.search.SetQuery("nmap")
	got := a.showHelp()
	if got.view != viewHelp || got.search.Input.Value() != "" {
		t.Fatal("help must activate with cleared query")
	}
}

// TestBackToSearchFocuses verifies returning to search refocuses input.
func TestBackToSearchFocuses(t *testing.T) {
	a := newTestApp(t, nil).showHelp()
	got := a.backToSearch()
	if got.view != viewSearch {
		t.Fatal("must return to search")
	}
	if !got.search.Input.Focused() {
		t.Fatal("search input must be focused after back")
	}
}
