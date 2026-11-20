// Package language persists the UI language preference across sessions
// via a JSON file. The preference is saved atomically like variables.
package language

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
	"github.com/halilkirazkaya/arsenal-ng/internal/paths"
)

// =============================================================================
// Language Store
// =============================================================================

// Store holds the UI language preference saved to disk.
// The language is set via "set language=tr" and read back on startup.
// Thread-safe for concurrent access.
type Store struct {
	mu sync.RWMutex
	// language is the normalized code (e.g., "en", "tr").
	language string
	// filePath is the path to language.json.
	filePath string
}

// Empty creates an in-memory store with persistence disabled.
// Use it as a fallback when the config directory is unavailable: all
// operations work, but nothing is written to disk.
func Empty() *Store {
	return &Store{language: config.DefaultLanguage}
}

// New creates a language store and loads the preference from disk.
// A missing file means first run and yields the default language.
// It returns an error if the config directory cannot be created or accessed.
func New() (*Store, error) {
	filePath, err := paths.Language()
	if err != nil {
		log.Printf("ERROR: Failed to get language path: %v", err)
		return nil, fmt.Errorf("failed to get language path: %w", err)
	}

	g := &Store{
		language: config.DefaultLanguage,
		filePath: filePath,
	}

	// Load the saved preference; a missing file is expected on first run.
	if err := g.loadFromFile(); err != nil && !os.IsNotExist(err) {
		log.Printf("ERROR: Failed to load language from %s: %v", filePath, err)
		return nil, fmt.Errorf("failed to load language: %w", err)
	}

	log.Printf("Language loaded: %s (file: %s)", g.language, filePath)
	return g, nil
}

// =============================================================================
// Store Operations
// =============================================================================

// Get returns the current language code.
func (g *Store) Get() string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.language
}

// Set validates, stores, and saves the language code.
// It returns an error for unsupported codes without changing the store.
func (g *Store) Set(code string) error {
	normalized, ok := config.NormalizeLanguage(code)
	if !ok {
		return fmt.Errorf("unsupported language %q (supported: %s)",
			code, strings.Join(config.SupportedLanguages, ", "))
	}

	g.mu.Lock()
	g.language = normalized
	g.mu.Unlock()

	log.Printf("Set language: %s", normalized)
	if err := g.saveToFile(); err != nil {
		log.Printf("ERROR: Failed to save language after Set: %v", err)
		return err
	}
	return nil
}
