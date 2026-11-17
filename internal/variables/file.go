// Package variables provides a thread-safe global variables store that
// persists across sessions via a JSON file. Variables are saved atomically.
package variables

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// =============================================================================
// File Persistence
// =============================================================================

// loadFromFile loads variables from the JSON file on disk.
// If the file doesn't exist, it returns os.ErrNotExist (expected on first run).
// If filePath is empty (persistence disabled), it returns nil without error.
func (g *Store) loadFromFile() error {
	if g.filePath == "" {
		return nil
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	data, err := os.ReadFile(g.filePath)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	var variables map[string]string
	if err := json.Unmarshal(data, &variables); err != nil {
		log.Printf("ERROR: Failed to parse variables.json: %v", err)
		return fmt.Errorf("failed to parse variables.json: %w", err)
	}

	// A file containing "null" decodes to a nil map, and Set would panic on it.
	if variables == nil {
		variables = make(map[string]string)
	}

	g.variables = variables
	return nil
}

// saveToFile saves the current variables to the JSON file on disk.
// It creates the file if missing and overwrites it otherwise.
// If filePath is empty (persistence disabled), it returns nil without error.
func (g *Store) saveToFile() error {
	if g.filePath == "" {
		return nil
	}

	g.mu.RLock()
	// Copy under the lock so disk I/O runs unlocked.
	varsCopy := make(map[string]string, len(g.variables))
	for k, v := range g.variables {
		varsCopy[k] = v
	}
	g.mu.RUnlock()

	// A map[string]string always marshals, so no error is possible here.
	data, _ := json.MarshalIndent(varsCopy, "", "  ")

	// Write atomically via a temp file. Variables often hold passwords, so the
	// file is readable by its owner only. Any leftover temp file is removed
	// first because WriteFile keeps the mode of an existing file.
	tmpPath := g.filePath + ".tmp"
	_ = os.Remove(tmpPath)
	if err := os.WriteFile(tmpPath, data, 0600); err != nil {
		log.Printf("ERROR: Failed to write variables file: %v", err)
		return fmt.Errorf("failed to write variables file: %w", err)
	}

	if err := os.Rename(tmpPath, g.filePath); err != nil {
		// Clean up temp file on error
		_ = os.Remove(tmpPath)
		log.Printf("ERROR: Failed to rename variables file: %v", err)
		return fmt.Errorf("failed to rename variables file: %w", err)
	}

	log.Printf("Saved %d variable(s) to %s", len(varsCopy), g.filePath)
	return nil
}
