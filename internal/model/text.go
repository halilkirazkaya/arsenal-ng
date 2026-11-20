// Package model defines the core data types for arsenal-ng.
//
// This file contains the localizable Text type used for cheat titles and
// descriptions. YAML accepts either a plain string (the default language)
// or a per-language map, so existing files keep working unchanged.
package model

import (
	"sort"

	"go.yaml.in/yaml/v4"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Localizable Text
// =============================================================================

// Text holds a string in one or more languages, keyed by language code.
type Text map[string]string

// NewText builds a single-language Text in the default language.
func NewText(s string) Text {
	if s == "" {
		return nil
	}
	return Text{config.DefaultLanguage: s}
}

// =============================================================================
// Language Resolution
// =============================================================================

// Resolve returns the text in the requested language.
// It falls back to the default language, then to any available translation,
// and finally to the empty string when the map holds nothing.
func (t Text) Resolve(lang string) string {
	if s, ok := t[lang]; ok && s != "" {
		return s
	}
	if s, ok := t[config.DefaultLanguage]; ok && s != "" {
		return s
	}
	// Sorted keys keep the last-resort pick deterministic.
	keys := make([]string, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if t[k] != "" {
			return t[k]
		}
	}
	return ""
}

// =============================================================================
// YAML Decoding
// =============================================================================

// UnmarshalYAML decodes either a plain string or a per-language map:
//
//	title: Port scan
//	title:
//	  en: Port scan
//	  tr: Port taraması
func (t *Text) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err == nil {
		*t = NewText(s)
		return nil
	}
	var m map[string]string
	if err := value.Decode(&m); err != nil {
		return err
	}
	*t = m
	return nil
}
