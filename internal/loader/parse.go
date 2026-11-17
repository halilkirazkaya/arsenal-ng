// Package loader reads cheat files from the embedded filesystem and
// converts them into Cheat objects.
package loader

import (
	"strings"

	"go.yaml.in/yaml/v4"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// YAML Parsing
// =============================================================================

// isYAMLFile checks if the path has a YAML extension.
// Matching is case-insensitive and covers both .yaml and .yml.
func isYAMLFile(path string) bool {
	lower := strings.ToLower(path)
	return strings.HasSuffix(lower, ".yaml") || strings.HasSuffix(lower, ".yml")
}

// parseCheatData parses YAML content into Cheats, tagging each with path.
// Each action in the file becomes one Cheat sharing the tool and tags.
func parseCheatData(data []byte, path string) ([]*model.Cheat, error) {
	var file model.CheatFile
	if err := yaml.Unmarshal(data, &file); err != nil {
		return nil, err
	}

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
