// Package main is the entry point for arsenal-ng, a modern pentest command launcher.
package main

import (
	"fmt"
	"io"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/loader"
	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/paths"
	"github.com/halilkirazkaya/arsenal-ng/internal/prefill"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui"
)

// =============================================================================
// Dependency Injection
// =============================================================================

// deps abstracts the side effects of run for testing.
type deps interface {
	// logPath returns the debug log path, or an error when unavailable.
	logPath() (string, error)
	// setupLog redirects the standard logger to the path.
	setupLog(path string) (io.Closer, error)
	// load reads all cheat files.
	load() ([]*model.Cheat, error)
	// interact runs the TUI and returns the final app state.
	interact(cheats []*model.Cheat) (ui.App, error)
	// prefill writes the command to the terminal input buffer.
	prefill(command string) error
	// warn prints a non-fatal warning to stderr.
	warn(format string, args ...any)
}

// =============================================================================
// Production Wiring
// =============================================================================

// production wires deps to the real implementations.
type production struct{}

// logPath resolves the real debug log path.
func (production) logPath() (string, error) { return paths.Log() }

// setupLog redirects bubbletea logging to the given file.
func (production) setupLog(path string) (io.Closer, error) {
	return tea.LogToFile(path, config.AppName)
}

// load reads all cheat files from disk.
func (production) load() ([]*model.Cheat, error) { return loader.Load() }

// interact runs the real fullscreen TUI and returns its final state.
func (production) interact(cheats []*model.Cheat) (ui.App, error) {
	app := ui.New(cheats)
	program := tea.NewProgram(app)
	result, err := program.Run()
	if err != nil {
		return ui.App{}, err
	}
	return result.(ui.App), nil
}

// prefill writes the chosen command into the terminal input buffer.
func (production) prefill(command string) error { return prefill.Command(command) }

// warn prints a non-fatal warning to stderr.
func (production) warn(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}
