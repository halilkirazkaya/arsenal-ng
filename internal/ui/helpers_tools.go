// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains tool list utility functions that extract and organize
// tool information from cheat data. It provides functions to get unique tools
// with their command counts and sort them alphabetically.
package ui

// =============================================================================
// Tool List Utilities
// =============================================================================

// ToolInfo represents information about a tool.
type ToolInfo struct {
	Name         string
	CommandCount int
}

// getToolInfoList extracts unique tool names with their command counts and returns them sorted.
func (m App) getToolInfoList() []ToolInfo {
	toolMap := make(map[string]int)
	for _, cheat := range m.cheats {
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

	// Sort tools alphabetically
	for i := 0; i < len(tools)-1; i++ {
		for j := i + 1; j < len(tools); j++ {
			if tools[i].Name > tools[j].Name {
				tools[i], tools[j] = tools[j], tools[i]
			}
		}
	}

	return tools
}

// getUniqueTools extracts unique tool names from all cheats and returns them sorted.
func (m App) getUniqueTools() []string {
	toolInfo := m.getToolInfoList()
	tools := make([]string, len(toolInfo))
	for i, info := range toolInfo {
		tools[i] = info.Name
	}
	return tools
}

