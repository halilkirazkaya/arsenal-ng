// Package ui provides the terminal user interface for arsenal-ng.
//
// App implements tea.Model and composes one sub-model per view. Views
// report outcomes; the parent owns cross-view transitions and shared data.
package ui

import (
	"fmt"
	"log"
	"strings"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/keys"
	"github.com/halilkirazkaya/arsenal-ng/internal/ui/finder"
	"github.com/halilkirazkaya/arsenal-ng/internal/variables"
)

// =============================================================================
// Set And Unset Flows
// =============================================================================

// applySet parses "set name=value" and stores the variable with feedback.
func (m App) applySet(query string) App {
	parts := strings.SplitN(query[4:], "=", 2)
	if len(parts) == 2 {
		name := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if name != "" && value != "" {
			// language is reserved: it switches translations, not a variable.
			if strings.EqualFold(name, keys.Language) {
				return m.applyLanguage(value)
			}
			m.search = applySetVar(m.search, m.globals, name, value)
			m.search = m.search.SetQuery("")
			m.search = m.search.ResetList(m.cheats)
			return m
		}
	}
	log.Printf("Invalid set command format")
	m.search.StatusMsg = "✗ Usage: set name=value"
	m.search.StatusIsError = true
	return m
}

// applyLanguage switches the UI language and persists the choice.
func (m App) applyLanguage(value string) App {
	if err := m.langs.Set(value); err != nil {
		log.Printf("Invalid language %q", value)
		m.search.StatusMsg = fmt.Sprintf("✗ %v", err)
		m.search.StatusIsError = true
		return m
	}
	m.lang = m.langs.Get()
	m.search = m.search.SetLang(m.lang)
	m.search = m.search.SetQuery("")
	m.search = m.search.ResetList(m.cheats)
	m.search.StatusMsg = fmt.Sprintf("✓ Language set to %s (%s)",
		config.LanguageName(m.lang), m.lang)
	return m
}

// applyUnset parses "unset name" and removes the variable with feedback.
func (m App) applyUnset(query string) App {
	name := strings.TrimSpace(query[6:])
	if name != "" {
		m.search = applyUnsetVar(m.search, m.globals, name)
		m.search = m.search.SetQuery("")
		m.search = m.search.ResetList(m.cheats)
		return m
	}
	log.Printf("Invalid unset command format")
	m.search.StatusMsg = "✗ Usage: unset name"
	m.search.StatusIsError = true
	return m
}

// applySetVar persists the variable and reports the result as a status.
func applySetVar(m finder.Model, globals *variables.Store, name, value string) finder.Model {
	if err := globals.Set(name, value); err != nil {
		log.Printf("ERROR: Failed to set variable %s: %v", name, err)
		m.StatusMsg = fmt.Sprintf("✗ Failed to save variable: %v", err)
		m.StatusIsError = true
	} else {
		// Name-only: values may be secrets visible on screen recordings.
		m.StatusMsg = fmt.Sprintf("✓ Set %s", name)
	}
	return m
}

// applyUnsetVar removes the variable and reports the result as a status.
func applyUnsetVar(m finder.Model, globals *variables.Store, name string) finder.Model {
	existed, err := globals.Unset(name)
	if err != nil {
		log.Printf("ERROR: Failed to unset variable %s: %v", name, err)
		m.StatusMsg = fmt.Sprintf("✗ Failed to save changes: %v", err)
		m.StatusIsError = true
	} else if existed {
		m.StatusMsg = fmt.Sprintf("✓ Unset %s", name)
	} else {
		log.Printf("Attempted to unset non-existent variable: %s", name)
		m.StatusMsg = fmt.Sprintf("✗ Variable '%s' not found", name)
		m.StatusIsError = true
	}
	return m
}
