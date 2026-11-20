// Package config provides application metadata, version info, and UI
// defaults such as terminal dimensions and input constraints.
package config

import (
	"testing"
)

// =============================================================================
// Language Tests
// =============================================================================

// TestNormalizeLanguage verifies codes normalize and unsupported ones fail.
func TestNormalizeLanguage(t *testing.T) {
	for code, want := range map[string]string{
		"en": "en", "tr": "tr", "TR": "tr", " tr ": "tr",
		"tr-TR": "tr", "en_US": "en",
	} {
		if got, ok := NormalizeLanguage(code); !ok || got != want {
			t.Errorf("NormalizeLanguage(%q) = %q, %v; want %q, true", code, got, ok, want)
		}
	}
	for _, code := range []string{"", "xx", "de", "english"} {
		if got, ok := NormalizeLanguage(code); ok {
			t.Errorf("NormalizeLanguage(%q) = %q, want unsupported", code, got)
		}
	}
}

// TestLanguageName verifies display names resolve with code fallback.
func TestLanguageName(t *testing.T) {
	if got := LanguageName("tr"); got != "Türkçe" {
		t.Fatalf("name = %q", got)
	}
	if got := LanguageName("xx"); got != "xx" {
		t.Fatalf("unknown name = %q, want code", got)
	}
}

// TestSupportedLanguagesStartWithDefault verifies the default is supported.
func TestSupportedLanguagesStartWithDefault(t *testing.T) {
	if !IsSupported(DefaultLanguage) {
		t.Fatal("default language must be supported")
	}
	if SupportedLanguages[0] != DefaultLanguage {
		t.Fatal("default language must come first")
	}
}
