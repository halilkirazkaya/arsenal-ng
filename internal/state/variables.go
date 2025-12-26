// Package state manages persistent and session-scoped state for the application.
//
// This file provides a thread-safe global variables store that persists across
// sessions via a JSON file. It supports setting, getting, unsetting variables,
// and applying them to command templates. Variables are saved atomically to disk.
package state

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/halilkirazkaya/arsenal-ng/internal/config"
)

// =============================================================================
// Global Variables Store
// =============================================================================

// Global holds persistent variables that are saved to disk.
// Variables are set via "set name=value" and can be used in commands.
// Variables persist across different shell sessions via variables.json file.
// Thread-safe for concurrent access.
type Global struct {
	mu        sync.RWMutex
	variables map[string]string
	filePath  string // Path to variables.json
}

// NewGlobal creates a new variables store and loads existing variables from disk.
// If the variables.json file doesn't exist, it will be created on first save.
// Returns an error if the config directory cannot be created or accessed.
func NewGlobal() (*Global, error) {
	filePath, err := config.GetVariablesPath()
	if err != nil {
		log.Printf("ERROR: Failed to get variables path: %v", err)
		return nil, fmt.Errorf("failed to get variables path: %w", err)
	}

	g := &Global{
		variables: make(map[string]string),
		filePath:  filePath,
	}

	// Load existing variables from disk (ignore error if file doesn't exist)
	if err := g.LoadFromFile(); err != nil && !os.IsNotExist(err) {
		// Log non-existent file errors silently, but return other errors
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

// =============================================================================
// CRUD Operations
// =============================================================================

// Set stores or updates a variable and saves to disk.
// If a variable with the same name exists, it will be overwritten.
// Returns an error if the save operation fails.
func (g *Global) Set(name, value string) error {
	g.mu.Lock()
	existed := false
	if _, exists := g.variables[name]; exists {
		existed = true
	}
	g.variables[name] = value
	g.mu.Unlock()

	action := "Set"
	if existed {
		action = "Updated"
	}
	log.Printf("%s variable: %s = %s", action, name, value)

	// Save to disk after updating
	if err := g.SaveToFile(); err != nil {
		log.Printf("ERROR: Failed to save variables after Set: %v", err)
		return err
	}

	return nil
}

// Get retrieves a variable's value.
// Returns the value and true if found, empty string and false otherwise.
func (g *Global) Get(name string) (string, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	val, ok := g.variables[name]
	return val, ok
}

// Unset removes a variable and saves to disk.
// Returns true if the variable existed and was removed, false otherwise.
// Returns an error if the save operation fails (but variable is still removed from memory).
func (g *Global) Unset(name string) (bool, error) {
	g.mu.Lock()
	existed := false
	if _, exists := g.variables[name]; exists {
		delete(g.variables, name)
		existed = true
	}
	g.mu.Unlock()

	// Save to disk after removing
	if existed {
		log.Printf("Unset variable: %s", name)
		if err := g.SaveToFile(); err != nil {
			log.Printf("ERROR: Failed to save variables after Unset: %v", err)
			return true, err
		}
		return true, nil
	}
	log.Printf("Attempted to unset non-existent variable: %s", name)
	return false, nil
}

// All returns a copy of all variables.
// Safe for iteration without holding the lock.
func (g *Global) All() map[string]string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	result := make(map[string]string, len(g.variables))
	for k, v := range g.variables {
		result[k] = v
	}
	return result
}

// Count returns the number of stored variables.
func (g *Global) Count() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.variables)
}

// =============================================================================
// Command Integration
// =============================================================================

// ApplyToCommand replaces {{var}} placeholders with stored values.
// Returns the modified command and a list of applied variable names.
func (g *Global) ApplyToCommand(command string) (string, []string) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var applied []string
	result := command

	for name, value := range g.variables {
		placeholder := "{{" + name + "}}"
		if strings.Contains(result, placeholder) {
			result = strings.ReplaceAll(result, placeholder, value)
			applied = append(applied, name)
		}
	}

	return result, applied
}

// =============================================================================
// Display
// =============================================================================

// FormatList returns a formatted string of all variables for display.
func (g *Global) FormatList() string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	if len(g.variables) == 0 {
		return "No variables set"
	}

	// Sort keys for consistent display
	keys := make([]string, 0, len(g.variables))
	for k := range g.variables {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Variables (%d):\n", len(g.variables)))
	for _, k := range keys {
		sb.WriteString(fmt.Sprintf("  %s = %s\n", k, g.variables[k]))
	}
	return sb.String()
}

// =============================================================================
// Persistence
// =============================================================================

// LoadFromFile loads variables from the JSON file on disk.
// If the file doesn't exist, it returns os.ErrNotExist (this is expected on first run).
// If filePath is empty (persistence disabled), returns nil without error.
// Thread-safe.
func (g *Global) LoadFromFile() error {
	if g.filePath == "" {
		// Persistence disabled (fallback mode)
		return nil
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	data, err := os.ReadFile(g.filePath)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		// Empty file is valid (no variables)
		return nil
	}

	var vars map[string]string
	if err := json.Unmarshal(data, &vars); err != nil {
		log.Printf("ERROR: Failed to parse variables.json: %v", err)
		return fmt.Errorf("failed to parse variables.json: %w", err)
	}

	// Replace existing variables with loaded ones
	g.variables = vars
	return nil
}

// SaveToFile saves the current variables to the JSON file on disk.
// Creates the file if it doesn't exist, overwrites if it does.
// If filePath is empty (persistence disabled), returns nil without error.
// Thread-safe.
func (g *Global) SaveToFile() error {
	if g.filePath == "" {
		// Persistence disabled (fallback mode)
		return nil
	}

	g.mu.RLock()
	// Create a copy to avoid holding lock during I/O
	varsCopy := make(map[string]string, len(g.variables))
	for k, v := range g.variables {
		varsCopy[k] = v
	}
	g.mu.RUnlock()

	// Marshal to JSON with indentation for readability
	data, err := json.MarshalIndent(varsCopy, "", "  ")
	if err != nil {
		log.Printf("ERROR: Failed to marshal variables: %v", err)
		return fmt.Errorf("failed to marshal variables: %w", err)
	}

	// Write atomically using a temporary file
	tmpPath := g.filePath + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0644); err != nil {
		log.Printf("ERROR: Failed to write variables file: %v", err)
		return fmt.Errorf("failed to write variables file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tmpPath, g.filePath); err != nil {
		// Clean up temp file on error
		_ = os.Remove(tmpPath)
		log.Printf("ERROR: Failed to rename variables file: %v", err)
		return fmt.Errorf("failed to rename variables file: %w", err)
	}

	log.Printf("Saved %d variable(s) to %s", len(varsCopy), g.filePath)
	return nil
}

// GetFilePath returns the path to the variables.json file.
// Useful for debugging or displaying to the user.
func (g *Global) GetFilePath() string {
	return g.filePath
}

