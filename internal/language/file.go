// Package language persists the UI language preference across sessions
// via a JSON file. The preference is saved atomically like variables.
package language

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// File Persistence
// =============================================================================

// languageFile is the on-disk shape of language.json.
type languageFile struct {
	// Language is the normalized code (e.g., "en", "tr").
	Language string `json:"language"`
}

// loadFromFile loads the language preference from the JSON file on disk.
// If the file doesn't exist, it returns os.ErrNotExist (expected on first run).
// If filePath is empty (persistence disabled), it returns nil without error.
func (g *Store) loadFromFile() error {
	if g.filePath == "" {
		return nil
	}

	data, err := os.ReadFile(g.filePath)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	var file languageFile
	if err := json.Unmarshal(data, &file); err != nil {
		log.Printf("ERROR: Failed to parse language.json: %v", err)
		return fmt.Errorf("failed to parse language.json: %w", err)
	}

	// A saved code from a newer version falls back to the default.
	normalized, ok := config.NormalizeLanguage(file.Language)
	if !ok {
		log.Printf("WARNING: Ignoring unsupported saved language %q", file.Language)
		return nil
	}

	g.mu.Lock()
	g.language = normalized
	g.mu.Unlock()
	return nil
}

// saveToFile saves the current language to the JSON file on disk.
// It creates the file if missing and overwrites it otherwise.
// If filePath is empty (persistence disabled), it returns nil without error.
func (g *Store) saveToFile() error {
	if g.filePath == "" {
		return nil
	}

	g.mu.RLock()
	language := g.language
	g.mu.RUnlock()

	// A struct always marshals, so no error is possible here.
	data, _ := json.MarshalIndent(languageFile{Language: language}, "", "  ")

	// Write atomically via a temp file. The preference holds no secrets,
	// so unlike variables.json it is world-readable. Any leftover temp
	// file is removed first because WriteFile keeps an existing file's mode.
	tmpPath := g.filePath + ".tmp"
	_ = os.Remove(tmpPath)
	if err := os.WriteFile(tmpPath, data, 0o644); err != nil {
		log.Printf("ERROR: Failed to write language file: %v", err)
		return fmt.Errorf("failed to write language file: %w", err)
	}

	if err := os.Rename(tmpPath, g.filePath); err != nil {
		// Clean up temp file on error
		_ = os.Remove(tmpPath)
		log.Printf("ERROR: Failed to rename language file: %v", err)
		return fmt.Errorf("failed to rename language file: %w", err)
	}

	log.Printf("Saved language %s to %s", language, g.filePath)
	return nil
}
