// Package loader handles loading and searching cheat files.
//
// This file contains functions to load cheat files from the embedded filesystem,
// parse YAML files, and convert them into Cheat objects. It uses Go's embed
// feature to include cheat files directly in the binary.
package loader

import (
	"embed"
	"io/fs"
	"log"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Embedded Cheat Files
// =============================================================================

//go:embed cheat-files
var embeddedCheats embed.FS

// =============================================================================
// Loading
// =============================================================================

// Load reads all cheat files from the embedded filesystem.
// Returns a flat list of Cheat objects, one per action.

func Load() ([]*model.Cheat, error) {
	var cheats []*model.Cheat
	fileCount := 0
	errorCount := 0

	err := fs.WalkDir(embeddedCheats, "cheat-files", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		// Only process YAML files
		if !isYAMLFile(path) {
			return nil
		}

		fileCount++
		// Parse the file
		fileCheats, err := parseCheatFile(path)
		if err != nil {
			// Skip invalid files but count the error!
			errorCount++
			return nil
		}

		cheats = append(cheats, fileCheats...)
		return nil
	})

	if err != nil {
		log.Printf("ERROR: Failed to walk cheat files directory: %v", err)
		return cheats, err
	}

	// Log summary only
	if errorCount > 0 {
		log.Printf("Cheat loading complete: %d files processed, %d cheats loaded, %d error(s) occurred", fileCount, len(cheats), errorCount)
	} else {
		log.Printf("Cheat loading complete: %d files processed, %d cheats loaded, 0 errors", fileCount, len(cheats))
	}
	return cheats, err
}

// isYAMLFile checks if the path has a YAML extension.
func isYAMLFile(path string) bool {
	lower := strings.ToLower(path)
	return strings.HasSuffix(lower, ".yaml") || strings.HasSuffix(lower, ".yml")
}

// parseCheatFile reads and parses a single cheat file.
func parseCheatFile(path string) ([]*model.Cheat, error) {
	data, err := embeddedCheats.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var file model.CheatFile
	if err := yaml.Unmarshal(data, &file); err != nil {
		return nil, err
	}

	// Convert each action to a Cheat
	cheats := make([]*model.Cheat, 0, len(file.Actions))
	for _, action := range file.Actions {
		cheat := &model.Cheat{
			Tool:     file.Tool,
			Tags:     file.Tags,
			Title:    action.Title,
			Desc:     action.Desc,
			Command:  action.Command,
			Filename: path,
		}
		cheats = append(cheats, cheat)
	}

	return cheats, nil
}
