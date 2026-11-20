// Package model defines the core data types for arsenal-ng.
package model

import (
	"testing"

	"go.yaml.in/yaml/v4"
)

// =============================================================================
// Text Tests
// =============================================================================

// TestResolvePrefersRequested verifies exact matches win over fallbacks.
func TestResolvePrefersRequested(t *testing.T) {
	text := Text{"en": "Port scan", "tr": "Port taraması"}
	if got := text.Resolve("tr"); got != "Port taraması" {
		t.Fatalf("tr = %q", got)
	}
	if got := text.Resolve("en"); got != "Port scan" {
		t.Fatalf("en = %q", got)
	}
}

// TestResolveFallsBack verifies the default and last-resort fallbacks.
func TestResolveFallsBack(t *testing.T) {
	text := Text{"en": "Port scan", "tr": "Port taraması"}
	if got := text.Resolve("de"); got != "Port scan" {
		t.Fatalf("unknown lang = %q, want default", got)
	}
	onlyTr := Text{"tr": "Port taraması"}
	if got := onlyTr.Resolve("en"); got != "Port taraması" {
		t.Fatalf("missing default = %q, want any", got)
	}
	var empty Text
	if got := empty.Resolve("en"); got != "" {
		t.Fatalf("empty = %q, want empty", got)
	}
}

// TestUnmarshalScalar verifies plain strings decode as the default language.
func TestUnmarshalScalar(t *testing.T) {
	var action Action
	if err := yaml.Unmarshal([]byte("title: Port scan\ncommand: nmap"), &action); err != nil {
		t.Fatal(err)
	}
	if got := action.Title.Resolve("en"); got != "Port scan" {
		t.Fatalf("title = %q", got)
	}
	// A scalar keeps resolving when the UI switches language.
	if got := action.Title.Resolve("tr"); got != "Port scan" {
		t.Fatalf("fallback = %q", got)
	}
}

// TestUnmarshalMap verifies per-language maps decode every translation.
func TestUnmarshalMap(t *testing.T) {
	var action Action
	data := []byte("title:\n  en: Port scan\n  tr: Port taraması\ncommand: nmap\n")
	if err := yaml.Unmarshal(data, &action); err != nil {
		t.Fatal(err)
	}
	if got := action.Title.Resolve("tr"); got != "Port taraması" {
		t.Fatalf("tr = %q", got)
	}
	if got := action.Title.Resolve("en"); got != "Port scan" {
		t.Fatalf("en = %q", got)
	}
}

// TestUnmarshalMissingDesc verifies absent descriptions resolve empty.
func TestUnmarshalMissingDesc(t *testing.T) {
	var action Action
	if err := yaml.Unmarshal([]byte("title: x\ncommand: y"), &action); err != nil {
		t.Fatal(err)
	}
	if got := action.Desc.Resolve("en"); got != "" {
		t.Fatalf("desc = %q, want empty", got)
	}
}
