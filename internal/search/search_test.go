// Package search ranks cheats against a query.
package search

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Search Fixtures
// =============================================================================

// testCheats builds a fixed cheat list shared by search tests.
// Tool names are chosen so "nmap" also matches "asnmap" loosely.
func testCheats() []*model.Cheat {
	return []*model.Cheat{
		{Tool: "nmap", Tags: []string{"scan"}, Title: "Port scan", Command: "nmap {{ip}}", Desc: "Scan ports"},
		{Tool: "asnmap", Tags: []string{"asn"}, Title: "ASN lookup", Command: "asnmap -d {{domain}}", Desc: "Query ASN data via nmap-style output"},
		{Tool: "ffuf", Tags: []string{"web", "fuzz"}, Title: "Directory fuzz", Command: "ffuf -u {{url}}", Desc: "Fuzz web paths"},
	}
}

// tools extracts tool names from results for compact assertions.
func tools(res []*model.Cheat) []string {
	got := make([]string, len(res))
	for i, c := range res {
		got[i] = c.Tool
	}
	return got
}

// =============================================================================
// Query Tests
// =============================================================================

// TestQueryEmptyReturnsAll verifies an empty query returns everything.
func TestQueryEmptyReturnsAll(t *testing.T) {
	cheats := testCheats()
	if got := Query(cheats, ""); len(got) != len(cheats) {
		t.Fatalf("empty query = %d, want %d", len(got), len(cheats))
	}
}

// TestQueryRequiresAllTerms verifies multi-word queries need every term.
func TestQueryRequiresAllTerms(t *testing.T) {
	cheats := testCheats()
	if got := Query(cheats, "nmap ffuf"); len(got) != 0 {
		t.Fatalf("disjoint terms = %v, want empty", tools(got))
	}
	if got := Query(cheats, "nmap scan"); len(got) != 1 || got[0].Tool != "nmap" {
		t.Fatalf("multi-term = %v, want [nmap]", tools(got))
	}
}

// TestQueryCaseInsensitive verifies queries match regardless of case.
func TestQueryCaseInsensitive(t *testing.T) {
	got := Query(testCheats(), "NMAP")
	if len(got) != 2 {
		t.Fatalf("NMAP = %v, want 2 results", tools(got))
	}
}
