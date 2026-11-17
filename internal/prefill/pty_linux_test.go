//go:build linux

// Package prefill injects the chosen command into the terminal's input
// buffer with the TIOCSTI ioctl, so it appears as if the user typed it.
// Users can review and edit the command before execution.
package prefill

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// =============================================================================
// Pty Pair Setup
// =============================================================================

// newPtyPair opens a master/slave pty pair for injection tests.
func newPtyPair() (master, slave int, err error) {
	master, err = unix.Open("/dev/ptmx", unix.O_RDWR|unix.O_NOCTTY|unix.O_CLOEXEC, 0)
	if err != nil {
		return 0, 0, fmt.Errorf("open ptmx: %w", err)
	}
	// TIOCSPTLCK takes a pointer to int; a null arg fails with EFAULT.
	if err := unix.IoctlSetInt(master, unix.TIOCSPTLCK, 0); err != nil {
		unix.Close(master)
		return 0, 0, fmt.Errorf("unlock pty: %w", err)
	}
	num, err := unix.IoctlGetInt(master, unix.TIOCGPTN)
	if err != nil {
		unix.Close(master)
		return 0, 0, fmt.Errorf("pty number: %w", err)
	}
	slave, err = unix.Open(fmt.Sprintf("/dev/pts/%d", num), unix.O_RDWR|unix.O_NOCTTY, 0)
	if err != nil {
		unix.Close(master)
		return 0, 0, fmt.Errorf("open pty slave: %w", err)
	}
	return master, slave, nil
}
