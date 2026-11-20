// Package loader reads cheat files from the embedded filesystem and
// converts them into Cheat objects.
package loader

import (
	"testing"
)

// =============================================================================
// Cheat Data Tests
// =============================================================================

// TestParseCheatData verifies YAML actions map to tagged Cheat values.
func TestParseCheatData(t *testing.T) {
	data := []byte(`tool: nmap
tags: [scan, recon]
actions:
  - title: Port scan
    desc: Scan ports
    command: nmap {{ip}}
  - title: Ping scan
    command: nmap -sn {{net}}
`)
	cheats, err := parseCheatData(data, "nmap.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if len(cheats) != 2 {
		t.Fatalf("cheats = %d, want 2", len(cheats))
	}
	if cheats[0].Tool != "nmap" || cheats[0].Title.Resolve("en") != "Port scan" {
		t.Fatalf("first = %+v", cheats[0])
	}
	if len(cheats[0].Tags) != 2 || cheats[1].Filename != "nmap.yaml" {
		t.Fatalf("tags/filename = %+v", cheats[0])
	}
	if got := cheats[1].Desc.Resolve("en"); got != "" {
		t.Fatalf("missing desc = %q, want empty", got)
	}
}

// TestParseCheatDataTranslated verifies per-language maps survive parsing.
func TestParseCheatDataTranslated(t *testing.T) {
	data := []byte(`tool: nmap
actions:
  - title:
      en: Port scan
      tr: Port taraması
    desc:
      en: Scan ports
      tr: Portları tara
    command: nmap {{ip}}
`)
	cheats, err := parseCheatData(data, "nmap.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if len(cheats) != 1 {
		t.Fatalf("cheats = %d, want 1", len(cheats))
	}
	if got := cheats[0].Title.Resolve("tr"); got != "Port taraması" {
		t.Fatalf("title tr = %q", got)
	}
	if got := cheats[0].Desc.Resolve("en"); got != "Scan ports" {
		t.Fatalf("desc en = %q", got)
	}
}

// TestParseCheatDataInvalid verifies malformed YAML errors but empty passes.
func TestParseCheatDataInvalid(t *testing.T) {
	if _, err := parseCheatData([]byte("{not yaml"), "bad.yaml"); err == nil {
		t.Fatal("invalid yaml = nil error")
	}
	if _, err := parseCheatData([]byte(""), "empty.yaml"); err != nil {
		t.Fatalf("empty yaml = %v", err)
	}
}
