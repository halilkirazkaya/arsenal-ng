//go:build darwin

// Package prefill injects the chosen command into the terminal's input
// buffer with the TIOCSTI ioctl, so it appears as if the user typed it.
// Users can review and edit the command before execution.
package prefill

import (
	"fmt"
	"unsafe"

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
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(master), unix.TIOCPTYGRANT, 0); errno != 0 {
		unix.Close(master)
		return 0, 0, fmt.Errorf("grant pty: %w", errno)
	}
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(master), unix.TIOCPTYUNLK, 0); errno != 0 {
		unix.Close(master)
		return 0, 0, fmt.Errorf("unlock pty: %w", errno)
	}
	var name [128]byte
	if _, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(master), unix.TIOCPTYGNAME, uintptr(unsafe.Pointer(&name[0]))); errno != 0 {
		unix.Close(master)
		return 0, 0, fmt.Errorf("pty name: %w", errno)
	}
	slave, err = unix.Open(nullTerminated(name[:]), unix.O_RDWR|unix.O_NOCTTY, 0)
	if err != nil {
		unix.Close(master)
		return 0, 0, fmt.Errorf("open pty slave: %w", err)
	}
	return master, slave, nil
}

// nullTerminated returns the C string prefix of the buffer.
func nullTerminated(buf []byte) string {
	for i, b := range buf {
		if b == 0 {
			return string(buf[:i])
		}
	}
	return string(buf)
}
