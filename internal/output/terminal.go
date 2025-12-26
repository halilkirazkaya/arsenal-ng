//go:build darwin || linux

// Package output provides terminal output functionality.
//
// This file uses TIOCSTI ioctl to inject commands into the terminal's input
// buffer, making them appear as if the user typed them. This allows users to
// review and edit commands before execution. Supports both Linux and macOS.
package output

import (
	"log"
	"os"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// =============================================================================
// Platform-specific ioctl Constants
// =============================================================================

const (
	// TIOCSTI - Terminal I/O Control Simulate Terminal Input
	// Injects characters into terminal input buffer
	ioctlTIOCSTI_Linux  = 0x5412
	ioctlTIOCSTI_Darwin = 0x80017472

	// TCGETS/TCSETS - Get/Set terminal attributes
	ioctlTCGETS_Linux  = 0x5401
	ioctlTCSETS_Linux  = 0x5402
	ioctlTCGETS_Darwin = 0x40487413
	ioctlTCSETS_Darwin = 0x80487414
)

// =============================================================================
// Terminal Prefill
// =============================================================================

// ToTerminal writes a command to the terminal's input buffer.
// The command appears as if the user typed it, ready for editing.
// This allows users to review and modify the command before execution.
//
// Note: On Linux kernel 6.2+, this requires:
//
//	sysctl -w dev.tty.legacy_tiocsti=1

func ToTerminal(command string) {
	if len(command) == 0 {
		log.Printf("WARNING: Attempted to output empty command to terminal")
		return
	}

	log.Printf("Opening /dev/tty for command injection")
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		log.Printf("ERROR: Failed to open /dev/tty: %v", err)
		return
	}
	defer tty.Close()

	fd := int(tty.Fd())
	tiocsti, tcgets, tcsets := getPlatformIoctls()

	// Save current terminal settings
	oldTermios, err := unix.IoctlGetTermios(fd, tcgets)
	if err != nil {
		log.Printf("ERROR: Failed to get terminal settings: %v", err)
		return
	}

	// Temporarily disable echo and canonical mode
	newTermios := *oldTermios
	newTermios.Lflag &^= unix.ECHO
	newTermios.Lflag &^= unix.ICANON
	if err := unix.IoctlSetTermios(fd, tcsets, &newTermios); err != nil {
		log.Printf("ERROR: Failed to set terminal settings: %v", err)
		return
	}

	// Inject each character into the terminal input buffer
	for _, char := range []byte(command) {
		_, _, _ = syscall.Syscall(
			syscall.SYS_IOCTL,
			uintptr(fd),
			uintptr(tiocsti),
			uintptr(unsafe.Pointer(&char)),
		)
	}

	// Restore terminal settings
	if err := unix.IoctlSetTermios(fd, tcsets, oldTermios); err != nil {
		log.Printf("WARNING: Failed to restore terminal settings: %v", err)
	}

	log.Printf("Command successfully injected to terminal (%d bytes)", len(command))
}

// getPlatformIoctls returns the correct ioctl constants for the current OS.
func getPlatformIoctls() (tiocsti, tcgets, tcsets uint) {
	if runtime.GOOS == "darwin" {
		return ioctlTIOCSTI_Darwin, ioctlTCGETS_Darwin, ioctlTCSETS_Darwin
	}
	return ioctlTIOCSTI_Linux, ioctlTCGETS_Linux, ioctlTCSETS_Linux
}
