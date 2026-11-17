// Package paths resolves OS-specific configuration directories. It follows
// the XDG Base Directory Specification on Linux, uses ~/Library/Application
// Support on macOS, and ~/.config elsewhere.
package paths

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Public Locations
// =============================================================================

// ConfigDir returns the app config directory, creating it when needed.
func ConfigDir() (string, error) {
	base, err := baseDir(runtime.GOOS)
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(base, config.AppName)
	// MkdirAll keeps a missing parent chain from failing first runs.
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", err
	}

	return configDir, nil
}

// Variables returns the path to the user variable store file.
func Variables() (string, error) {
	configDir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "variables.json"), nil
}

// Log returns the path to the debug log beside the variable store.
func Log() (string, error) {
	configDir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "debug.log"), nil
}

// =============================================================================
// Platform Roots
// =============================================================================

// baseDir returns the parent directory holding the app config dir.
func baseDir(goos string) (string, error) {
	switch goos {
	case "linux":
		// XDG override wins when set; otherwise default below.
		if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
			return xdgConfig, nil
		}
		// Explicit XDG value absent, so fall back to ~/.config.
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".config"), nil

	case "darwin":
		// macOS convention stores app data under Application Support.
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "Library", "Application Support"), nil

	default:
		// Remaining systems share the plain ~/.config layout.
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".config"), nil
	}
}
