//go:build darwin || linux

// Package prefill injects the chosen command into the terminal's input
// buffer with the TIOCSTI ioctl, so it appears as if the user typed it.
// Users can review and edit the command before execution.
package prefill

import (
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"testing"
)

// =============================================================================
// Test Fixtures
// =============================================================================

// termiosGet returns the platform TCGETS constant for termios assertions.
func termiosGet() uint {
	_, tcgets, _ := platformIoctls()
	return tcgets
}

// =============================================================================
// Platform Ioctls
// =============================================================================

// TestPlatformIoctls verifies the ioctl constants for the active platform.
func TestPlatformIoctls(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)
	tiocsti, tcgets, tcsets := platformIoctls()
	if tiocsti == 0 || tcgets == 0 || tcsets == 0 {
		t.Fatal("ioctl constants must be non-zero")
	}
	if runtime.GOOS == "darwin" {
		if tiocsti != 0x80017472 || tcgets != 0x40487413 || tcsets != 0x80487414 {
			t.Fatalf("darwin ioctls = %x/%x/%x", tiocsti, tcgets, tcsets)
		}
	} else {
		if tiocsti != 0x5412 || tcgets != 0x5401 || tcsets != 0x5402 {
			t.Fatalf("linux ioctls = %x/%x/%x", tiocsti, tcgets, tcsets)
		}
	}
}

// =============================================================================
// Command Entry Point
// =============================================================================

// TestCommandEmptyIsNoop verifies an empty command returns nil silently.
func TestCommandEmptyIsNoop(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)
	if err := Command(""); err != nil {
		t.Fatalf("empty command = %v, want nil", err)
	}
}

// TestCommandWithoutTTY verifies the open failure mentions /dev/tty.
func TestCommandWithoutTTY(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)
	// No controlling terminal in tests: must fail with an open error.
	if err := Command("echo hi"); err == nil {
		t.Fatal("expected /dev/tty error without a terminal")
	} else if !strings.Contains(err.Error(), "/dev/tty") {
		t.Fatalf("error = %q, want /dev/tty mention", err)
	}
}
