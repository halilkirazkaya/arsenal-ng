// Package variables provides a thread-safe global variables store that
// persists across sessions via a JSON file. Variables are saved atomically.
package variables

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/halilkirazkaya/arsenal-ng/internal/paths"
)

// =============================================================================
// Variable Store
// =============================================================================

// Store holds persistent variables that are saved to disk.
// Variables are set via "set name=value" and can be used in commands.
// Variables persist across different shell sessions via variables.json file.
// Thread-safe for concurrent access.
type Store struct {
	mu sync.RWMutex
	// variables maps names to values in memory.
	variables map[string]string
	// filePath is the path to variables.json.
	filePath string
}

// Empty creates an in-memory store with persistence disabled.
// Use it as a fallback when the config directory is unavailable: all
// operations work, but nothing is written to disk.
func Empty() *Store {
	return &Store{variables: make(map[string]string)}
}

// New creates a new variables store and loads existing variables from disk.
// If the variables.json file doesn't exist, it will be created on first save.
// It returns an error if the config directory cannot be created or accessed.
func New() (*Store, error) {
	filePath, err := paths.Variables()
	if err != nil {
		log.Printf("ERROR: Failed to get variables path: %v", err)
		return nil, fmt.Errorf("failed to get variables path: %w", err)
	}

	g := &Store{
		variables: make(map[string]string),
		filePath:  filePath,
	}

	// Load existing variables from disk; a missing file is expected on first run.
	if err := g.loadFromFile(); err != nil && !os.IsNotExist(err) {
		log.Printf("ERROR: Failed to load variables from %s: %v", filePath, err)
		return nil, fmt.Errorf("failed to load variables: %w", err)
	}

	if len(g.variables) > 0 {
		log.Printf("Loaded %d variable(s) from %s", len(g.variables), filePath)
	} else {
		log.Printf("No variables found, starting with empty store (file: %s)", filePath)
	}

	return g, nil
}
