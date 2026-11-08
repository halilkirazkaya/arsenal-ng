// Package finder provides the main search view: cheat filtering, list
// navigation, special commands, and selection of a cheat to execute.
package finder

// =============================================================================
// Suggestion Data
// =============================================================================

// Hint represents a variable suggestion for the set command hints.
type Hint struct {
	// Variable name shown after "set ".
	Name string
	// One-line description of the variable.
	Desc string
}

// CommonHints contains common variable suggestions for set command hints.
var CommonHints = []Hint{
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
