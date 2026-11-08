// Package ui provides the terminal user interface for arsenal-ng.
package ui

import (
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/ui/finder"
)

// =============================================================================
// Resize Tests
// =============================================================================

// toArgsView moves the app into the args view for routing tests.
func toArgsView(a App) App {
	m, _ := a.pickCheat(finder.Pick{Cheat: argCheat()})
	return m.(App)
}

// TestResizeApplies verifies window resizes update dimensions and input.
func TestResizeApplies(t *testing.T) {
	a := newTestApp(t, nil)
	m, _ := a.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	got := m.(App)
	if got.width != 120 || got.height != 40 {
		t.Fatalf("dims = %dx%d", got.width, got.height)
	}
	if got.search.Input.Width() != 110 {
		t.Fatalf("input width = %d, want 110", got.search.Input.Width())
	}
}

// TestResizeZeroIgnored verifies zero-size events leave dimensions alone.
func TestResizeZeroIgnored(t *testing.T) {
	a := newTestApp(t, nil)
	m, _ := a.Update(tea.WindowSizeMsg{Width: 0, Height: 0})
	got := m.(App)
	if got.width != a.width || got.height != a.height {
		t.Fatal("zero-size event must be ignored")
	}
}

// TestResizeRebuildsToolsVars verifies resizes rebuild table-backed views.
func TestResizeRebuildsToolsVars(t *testing.T) {
	a := newTestApp(t, nil)
	if err := a.globals.Set("target", "1.1.1.1"); err != nil {
		t.Fatal(err)
	}
	for _, v := range []view{viewTools, viewVariables} {
		a.view = v
		if v == viewTools {
			a = a.showTools()
		} else {
			a = a.showVariables()
		}
		m, _ := a.Update(tea.WindowSizeMsg{Width: 100, Height: 30})
		if _, ok := m.(App); !ok {
			t.Fatalf("%v: update must keep App type", v)
		}
	}
}

// =============================================================================
// View Routing
// =============================================================================

// TestRoutesStayInView verifies plain keys keep the app in each view.
func TestRoutesStayInView(t *testing.T) {
	for _, tc := range []struct {
		name string
		mk   func(App) App
		msg  tea.Msg
	}{
		{"search", func(a App) App { return a }, keyRune("n")},
		{"args", toArgsView, keyRune("x")},
		{"variables-view", func(a App) App { return a.showVariables() }, keyRune("x")},
		{"tools", func(a App) App { return a.showTools() }, keyRune("x")},
		{"help", func(a App) App { return a.showHelp() }, keyRune("x")},
	} {
		a := tc.mk(newTestApp(t, nil))
		want := a.view
		m, _ := a.Update(tc.msg)
		if m.(App).view != want {
			t.Fatalf("%s: left view %v", tc.name, want)
		}
	}
}

// TestHelpBackReturnsToSearch verifies the help back outcome returns home.
func TestHelpBackReturnsToSearch(t *testing.T) {
	a := newTestApp(t, nil).showHelp()
	m, _ := a.Update(keyEsc())
	if got := m.(App); got.view != viewSearch {
		t.Fatalf("view = %v, want search", got.view)
	}
}
