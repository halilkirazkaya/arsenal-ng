// Package config provides application metadata, version info, and UI
// defaults such as terminal dimensions and input constraints.
package config

import (
	"strings"
)

// =============================================================================
// Language Definitions
// =============================================================================

// DefaultLanguage is the fallback language for cheat titles and descriptions.
const DefaultLanguage = "en"

// SupportedLanguages lists the language codes accepted by set language=.
var SupportedLanguages = []string{"en", "tr"}

// LanguageNames maps codes to display names for user feedback.
var LanguageNames = map[string]string{
	"en": "English",
	"tr": "Türkçe",
}

// =============================================================================
// Language Helpers
// =============================================================================

// NormalizeLanguage lowercases the code and strips any region suffix.
// It reports false when the code is not supported (e.g., "xx", "").
func NormalizeLanguage(code string) (string, bool) {
	base, _, _ := strings.Cut(strings.TrimSpace(strings.ToLower(code)), "-")
	base, _, _ = strings.Cut(base, "_")
	if !IsSupported(base) {
		return "", false
	}
	return base, true
}

// IsSupported reports whether the code is a supported language.
func IsSupported(code string) bool {
	for _, supported := range SupportedLanguages {
		if code == supported {
			return true
		}
	}
	return false
}

// LanguageName returns the display name for a supported code.
func LanguageName(code string) string {
	if name, ok := LanguageNames[code]; ok {
		return name
	}
	return code
}
