// Package loader reads cheat files from the embedded filesystem and
// converts them into Cheat objects.
package loader

import (
	"embed"
	"io/fs"
	"log"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Embedded Filesystem
// =============================================================================

//go:embed cheat-files
var embeddedCheats embed.FS

// =============================================================================
// Cheat Loading
// =============================================================================

// Load reads all cheat files from the embedded filesystem.
// It returns a flat list of Cheat objects, one per action.
func Load() ([]*model.Cheat, error) {
	return loadFromFS(embeddedCheats, "cheat-files")
}

// loadFromFS reads all YAML cheat files under root in the given filesystem.
// Invalid files are skipped and counted, while walk errors abort the load.
func loadFromFS(fsys fs.FS, root string) ([]*model.Cheat, error) {
	var cheats []*model.Cheat
	fileCount := 0
	errorCount := 0

	err := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		if !isYAMLFile(path) {
			return nil
		}

		fileCount++
		fileCheats, err := parseCheatFile(fsys, path)
		if err != nil {
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

	if errorCount > 0 {
		log.Printf("Cheat loading complete: %d files processed, %d cheats loaded, %d error(s) occurred", fileCount, len(cheats), errorCount)
	} else {
		log.Printf("Cheat loading complete: %d files processed, %d cheats loaded, 0 errors", fileCount, len(cheats))
	}
	return cheats, err
}

// parseCheatFile reads and parses a single cheat file.
// It tags every resulting Cheat with the file path.
func parseCheatFile(fsys fs.FS, path string) ([]*model.Cheat, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, err
	}
	return parseCheatData(data, path)
}
