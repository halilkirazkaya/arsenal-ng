// Package config provides application metadata, version info, and UI
// defaults such as terminal dimensions and input constraints.
package config

import (
	"strings"
	"testing"
)

// =============================================================================
// Metadata Tests
// =============================================================================

// TestVersionInfo verifies the formatted version matches the raw version.
func TestVersionInfo(t *testing.T) {
	if got := GetVersionInfo(); got != "v"+Version {
		t.Fatalf("version = %q", got)
	}
	if Version == "" {
		t.Fatal("Version must not be empty")
	}
}

// TestAppIdentity verifies the app name and logo banner are present.
func TestAppIdentity(t *testing.T) {
	if AppName == "" {
		t.Fatal("AppName must not be empty")
	}
	if !strings.Contains(Logo, "█") {
		t.Fatal("Logo must contain block art")
	}
}

// TestLimitsSane verifies terminal defaults and input limits are positive.
func TestLimitsSane(t *testing.T) {
	if DefaultWidth <= 0 || DefaultHeight <= 0 {
		t.Fatal("terminal defaults must be positive")
	}
	if SearchCharLimit <= 0 || ArgCharLimit <= 0 {
		t.Fatal("char limits must be positive")
	}
	if SearchPlaceholder == "" {
		t.Fatal("placeholder must not be empty")
	}
}
