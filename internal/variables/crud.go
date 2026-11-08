// Package variables provides a thread-safe global variables store that
// persists across sessions via a JSON file. Variables are saved atomically.
package variables

import (
	"log"
)

// =============================================================================
// Store Operations
// =============================================================================

// Set stores or updates a variable and saves to disk.
// If a variable with the same name exists, it will be overwritten.
// It returns an error if the save operation fails.
func (g *Store) Set(name, value string) error {
	g.mu.Lock()
	_, existed := g.variables[name]
	g.variables[name] = value
	g.mu.Unlock()

	action := "Set"
	if existed {
		action = "Updated"
	}
	log.Printf("%s variable: %s = %s", action, name, value)

	if err := g.saveToFile(); err != nil {
		log.Printf("ERROR: Failed to save variables after Set: %v", err)
		return err
	}

	return nil
}

// Get retrieves a variable's value.
// It returns the value and true if found, empty string and false otherwise.
func (g *Store) Get(name string) (string, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	val, ok := g.variables[name]
	return val, ok
}

// Unset removes a variable and saves to disk.
// It returns true when the variable existed and was removed, false otherwise,
// plus an error if the save fails (the variable is still removed from memory).
func (g *Store) Unset(name string) (bool, error) {
	g.mu.Lock()
	_, existed := g.variables[name]
	if existed {
		delete(g.variables, name)
	}
	g.mu.Unlock()

	if existed {
		log.Printf("Unset variable: %s", name)
		if err := g.saveToFile(); err != nil {
			log.Printf("ERROR: Failed to save variables after Unset: %v", err)
			return true, err
		}
		return true, nil
	}
	log.Printf("Attempted to unset non-existent variable: %s", name)
	return false, nil
}

// All returns a copy of all variables.
// The copy is safe for iteration without holding the lock.
func (g *Store) All() map[string]string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	result := make(map[string]string, len(g.variables))
	for k, v := range g.variables {
		result[k] = v
	}
	return result
}

// Count returns the number of stored variables.
func (g *Store) Count() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.variables)
}
