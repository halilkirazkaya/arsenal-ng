// Package model defines the core data types for arsenal-ng.
//
// This file contains tool summaries extracted from cheat data: unique tool
// names with their command counts, sorted alphabetically.
package model

import (
	"sort"
)

// =============================================================================
// Tool Summary
// =============================================================================

// ToolInfo represents information about a tool.
// It pairs a tool name with its command count.
type ToolInfo struct {
	// Name is the tool name.
	Name string
	// CommandCount is the number of commands for this tool.
	CommandCount int
}

// ToolSummary extracts unique tool names with command counts, sorted.
// Cheats with an empty tool name are skipped.
func ToolSummary(cheats []*Cheat) []ToolInfo {
	toolMap := make(map[string]int)
	for _, cheat := range cheats {
		if cheat.Tool != "" {
			toolMap[cheat.Tool]++
		}
	}

	tools := make([]ToolInfo, 0, len(toolMap))
	for tool, count := range toolMap {
		tools = append(tools, ToolInfo{
			Name:         tool,
			CommandCount: count,
		})
	}

	sort.Slice(tools, func(i, j int) bool { return tools[i].Name < tools[j].Name })

	return tools
}
