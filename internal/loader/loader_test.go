// Package loader reads cheat files from the embedded filesystem and
// converts them into Cheat objects.
package loader

import (
	"io"
	"log"
	"os"
	"testing"
)

// =============================================================================
// Loader Tests
// =============================================================================

// TestIsYAMLFile verifies YAML extension detection, including case variants.
func TestIsYAMLFile(t *testing.T) {
	for path, want := range map[string]bool{
		"nmap.yaml": true, "tool.YML": true, "a/b.Yaml": true,
		"readme.md": false, "tool.yaml.bak": false, "": false,
	} {
		if got := isYAMLFile(path); got != want {
			t.Errorf("isYAMLFile(%q) = %v, want %v", path, got, want)
		}
	}
}

// TestParseCheatFile verifies a real embedded file parses into nmap cheats.
func TestParseCheatFile(t *testing.T) {
	cheats, err := parseCheatFile(embeddedCheats, "cheat-files/nmap.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if len(cheats) == 0 {
		t.Fatal("nmap.yaml produced no cheats")
	}
	for _, c := range cheats {
		if c.Tool != "nmap" {
			t.Fatalf("tool = %q, want nmap", c.Tool)
		}
		if c.Command == "" || c.Filename != "cheat-files/nmap.yaml" {
			t.Fatalf("cheat = %+v", c)
		}
	}
}

// TestParseCheatFileMissing verifies a missing path returns an error.
func TestParseCheatFileMissing(t *testing.T) {
	if _, err := parseCheatFile(embeddedCheats, "cheat-files/does-not-exist.yaml"); err == nil {
		t.Fatal("missing file = nil error")
	}
}

// TestLoad verifies the embedded tree loads complete, non-empty cheats.
func TestLoad(t *testing.T) {
	log.SetOutput(io.Discard)
	defer log.SetOutput(os.Stderr)

	cheats, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if len(cheats) == 0 {
		t.Fatal("Load produced no cheats")
	}
	seen := map[string]bool{}
	for _, c := range cheats {
		if c.Tool == "" || c.Command == "" || c.Title == "" {
			t.Fatalf("incomplete cheat: %+v", c)
		}
		seen[c.Tool] = true
	}
	for _, tool := range []string{"nmap", "ffuf"} {
		if !seen[tool] {
			t.Fatalf("tool %q missing from load", tool)
		}
	}
}
