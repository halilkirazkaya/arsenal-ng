// Package ui provides the terminal user interface for arsenal-ng.
//
// This file defines special command prefixes (set, unset, show) and common
// variable suggestions used for command hints in the search view.
package ui

// =============================================================================
// Special Command Prefixes
// =============================================================================

const (
	cmdSet   = "set"
	cmdUnset = "unset"
	cmdShow  = "show"
)

// =============================================================================
// Common Variable Suggestions
// =============================================================================

// variableHint represents a variable suggestion for the set command hints.
type variableHint struct {
	Name string
	Desc string
}

// commonVariables contains common variable suggestions for set command hints.
var commonVariables = []variableHint{
	{Name: "domain", Desc: "Target domain (e.g., corp.local)"},
	{Name: "hash", Desc: "NTLM hash"},
	{Name: "ip", Desc: "Target IP address (e.g., 10.10.10.10)"},
	{Name: "lhost", Desc: "Local host (your IP)"},
	{Name: "lport", Desc: "Local port (for reverse shell)"},
	{Name: "output", Desc: "Output file name"},
	{Name: "pass", Desc: "Password"},
	{Name: "port", Desc: "Port number (e.g., 445)"},
	{Name: "url", Desc: "Target URL"},
	{Name: "user", Desc: "Username (e.g., admin)"},
	{Name: "wordlist", Desc: "Wordlist path"},
}

