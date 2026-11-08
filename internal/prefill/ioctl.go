//go:build darwin || linux

// Package prefill injects the chosen command into the terminal's input
// buffer with the TIOCSTI ioctl, so it appears as if the user typed it.
// Users can review and edit the command before execution.
package prefill

import (
	"runtime"
)

// =============================================================================
// Platform Constants
// =============================================================================

const (
	// ioctlTIOCSTILinux injects characters into the Linux terminal input buffer.
	ioctlTIOCSTILinux = 0x5412
	// ioctlTIOCSTIDarwin injects characters into the Darwin terminal input buffer.
	ioctlTIOCSTIDarwin = 0x80017472

	// ioctlTCGETSLinux reads Linux terminal attributes.
	ioctlTCGETSLinux = 0x5401
	// ioctlTCSETSLinux writes Linux terminal attributes.
	ioctlTCSETSLinux = 0x5402
	// ioctlTCGETSDarwin reads Darwin terminal attributes.
	ioctlTCGETSDarwin = 0x40487413
	// ioctlTCSETSDarwin writes Darwin terminal attributes.
	ioctlTCSETSDarwin = 0x80487414
)

// =============================================================================
// Platform Selection
// =============================================================================

// platformIoctls returns the correct ioctl constants for the current OS.
func platformIoctls() (tiocsti, tcgets, tcsets uint) {
	if runtime.GOOS == "darwin" {
		return ioctlTIOCSTIDarwin, ioctlTCGETSDarwin, ioctlTCSETSDarwin
	}
	return ioctlTIOCSTILinux, ioctlTCGETSLinux, ioctlTCSETSLinux
}
