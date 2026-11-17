// Package main is the entry point for arsenal-ng, a modern pentest command launcher.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui"
)

// =============================================================================
// Test Doubles
// =============================================================================

// stub wires deps to scripted behavior and records calls for assertions.
type stub struct {
	// logErr forces logPath to fail.
	logErr error
	// setupErr forces setupLog to fail.
	setupErr error
	// cheats is the fixture list returned by load.
	cheats []*model.Cheat
	// loadErr forces load to fail.
	loadErr error
	// app is the fixture state returned by interact.
	app ui.App
	// interactErr forces interact to fail.
	interactErr error
	// prefillErr forces prefill to fail.
	prefillErr error
	// warnings records warn calls in order.
	warnings []string
	// prefilled records the last command passed to prefill.
	prefilled string
	// closed tracks whether the log file was closed.
	closed bool
}

// logPath returns the fixture path or the scripted error.
func (s *stub) logPath() (string, error) {
	if s.logErr != nil {
		return "", s.logErr
	}
	return "/tmp/test.log", nil
}

// stubCloser marks the stub log as closed for assertions.
type stubCloser struct{ s *stub }

// Close records the close call on the parent stub.
func (c stubCloser) Close() error { c.s.closed = true; return nil }

// setupLog returns a recording closer or the scripted error.
func (s *stub) setupLog(string) (io.Closer, error) {
	if s.setupErr != nil {
		return nil, s.setupErr
	}
	return stubCloser{s}, nil
}

// load returns the fixture cheats or the scripted error.
func (s *stub) load() ([]*model.Cheat, error) { return s.cheats, s.loadErr }

// interact returns the fixture app or the scripted error.
func (s *stub) interact([]*model.Cheat) (ui.App, error) { return s.app, s.interactErr }

// prefill records the command and returns the scripted error.
func (s *stub) prefill(command string) error {
	s.prefilled = command
	return s.prefillErr
}

// warn records a formatted warning for later assertions.
func (s *stub) warn(format string, args ...any) {
	s.warnings = append(s.warnings, fmt.Sprintf(format, args...))
}

// quiet silences the standard logger for the duration of a test.
func quiet(t *testing.T) {
	t.Helper()
	log.SetOutput(io.Discard)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })
}
