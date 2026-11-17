//go:build darwin || linux

// Package prefill injects the chosen command into the terminal's input
// buffer with the TIOCSTI ioctl, so it appears as if the user typed it.
// Users can review and edit the command before execution.
package prefill

import (
	"io"
	"log"
	"os"
	"testing"

	"golang.org/x/sys/unix"
)

// =============================================================================
// Termios Round Trip
// =============================================================================

// TestPtyTermiosRoundTrip verifies the echo-clear plus restore dance on a pty.
func TestPtyTermiosRoundTrip(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)

	master, slave, err := newPtyPair()
	if err != nil {
		t.Skipf("no pty: %v", err)
	}
	defer unix.Close(master)
	defer unix.Close(slave)

	_, _, tcsets := platformIoctls()
	before, err := unix.IoctlGetTermios(slave, termiosGet())
	if err != nil {
		t.Fatal(err)
	}

	// Mirror inject's termios dance: disable echo/canonical, then restore.
	changed := *before
	changed.Lflag &^= unix.ECHO
	changed.Lflag &^= unix.ICANON
	if err := unix.IoctlSetTermios(slave, tcsets, &changed); err != nil {
		t.Fatal(err)
	}
	mid, err := unix.IoctlGetTermios(slave, termiosGet())
	if err != nil {
		t.Fatal(err)
	}
	if mid.Lflag&unix.ECHO != 0 || mid.Lflag&unix.ICANON != 0 {
		t.Fatal("termios flags not disabled")
	}
	if err := unix.IoctlSetTermios(slave, tcsets, before); err != nil {
		t.Fatal(err)
	}
	after, err := unix.IoctlGetTermios(slave, termiosGet())
	if err != nil {
		t.Fatal(err)
	}
	// The kernel may normalize other bits; only ECHO/ICANON must match.
	const mask = unix.ECHO | unix.ICANON
	if after.Lflag&mask != before.Lflag&mask {
		t.Fatalf("flags not restored: got %x, want %x", after.Lflag&mask, before.Lflag&mask)
	}
}

// =============================================================================
// Inject Guards
// =============================================================================

// TestInjectNonTerminal verifies a regular file fd fails termios lookup.
func TestInjectNonTerminal(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)

	// A regular file fd is not a terminal: termios get must fail.
	f, err := os.CreateTemp("", "prefill")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	if err := inject(int(f.Fd()), "echo hi"); err == nil {
		t.Fatal("expected termios error for non-terminal fd")
	}
}
