// Package config provides application configuration and constants.
//
// This file contains path management functions for OS-specific configuration
// directories. It follows XDG Base Directory Specification on Linux, uses
// ~/Library/Application Support on macOS, and provides functions to get
// the config directory and variables.json file path.
package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// =============================================================================
// Config Path Management
// =============================================================================

// GetConfigDir returns the OS-specific configuration directory for arsenal-ng.
// Follows XDG Base Directory Specification on Linux, uses ~/Library/Application Support on macOS,
// and ~/.config on other Unix systems.

func GetConfigDir() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "linux":
		// XDG_CONFIG_HOME takes precedence
		if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
			configDir = filepath.Join(xdgConfig, AppName)
		} else {
			// Default: ~/.config/arsenal-ng
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			configDir = filepath.Join(home, ".config", AppName)
		}

	case "darwin":
		// macOS: ~/Library/Application Support/arsenal-ng
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configDir = filepath.Join(home, "Library", "Application Support", AppName)

	default:
		// Other Unix systems: ~/.config/arsenal-ng
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configDir = filepath.Join(home, ".config", AppName)
	}

	// Ensure directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

// GetVariablesPath returns the full path to the variables.json file.
func GetVariablesPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "variables.json"), nil
}

// GetLogPath returns the full path to the debug.log file.
// Logs are stored in the same directory as variables.json.
func GetLogPath() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "debug.log"), nil
}