// Package search ranks cheats against a query.
package search

import (
	"testing"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

// =============================================================================
// Ranking Tests
// =============================================================================

// TestQueryRanksToolFirst verifies tool-name matches outrank description hits.
func TestQueryRanksToolFirst(t *testing.T) {
	// "nmap" matches nmap's tool name but only asnmap's description:
	// the tool match must outrank regardless of input order.
	cheats := []*model.Cheat{testCheats()[1], testCheats()[0]}
	got := Query(cheats, "nmap")
	if len(got) != 2 || got[0].Tool != "nmap" || got[1].Tool != "asnmap" {
		t.Fatalf("order = %v, want [nmap asnmap]", tools(got))
	}
}

// TestQueryRanksExactOverPrefix verifies exact matches beat prefix matches.
func TestQueryRanksExactOverPrefix(t *testing.T) {
	cheats := []*model.Cheat{
		{Tool: "nmapautomator", Title: "auto", Command: "run"},
		{Tool: "nmap", Title: "scan", Command: "run"},
	}
	got := Query(cheats, "nmap")
	if len(got) != 2 || got[0].Tool != "nmap" {
		t.Fatalf("order = %v, want nmap first", tools(got))
	}
}

// TestQuerySumsMultiFieldMatches verifies matches in two fields outrank one.
func TestQuerySumsMultiFieldMatches(t *testing.T) {
	// Both match "dns" in tags, but dnsrecon also matches in tool name.
	cheats := []*model.Cheat{
		{Tool: "amass", Tags: []string{"dns"}, Title: "enum", Command: "run"},
		{Tool: "dnsrecon", Tags: []string{"dns"}, Title: "enum", Command: "run"},
	}
	got := Query(cheats, "dns")
	if len(got) != 2 || got[0].Tool != "dnsrecon" {
		t.Fatalf("order = %v, want dnsrecon first", tools(got))
	}
}

// TestQueryStableTies verifies identical scores keep input order.
func TestQueryStableTies(t *testing.T) {
	// Identical scores keep input order.
	cheats := []*model.Cheat{
		{Tool: "b-tool", Title: "scan", Command: "run"},
		{Tool: "a-tool", Title: "scan", Command: "run"},
	}
	got := Query(cheats, "scan")
	if len(got) != 2 || got[0].Tool != "b-tool" || got[1].Tool != "a-tool" {
		t.Fatalf("order = %v, want input order", tools(got))
	}
}

// TestQueryMatchesAllFields verifies every searchable field can match.
func TestQueryMatchesAllFields(t *testing.T) {
	cheats := []*model.Cheat{
		{Tool: "t1", Tags: []string{"needle"}, Title: "x", Command: "run"},
		{Tool: "t2", Title: "needle here", Command: "run"},
		{Tool: "t3", Title: "x", Command: "run needle"},
		{Tool: "t4", Title: "x", Command: "run", Desc: "has needle inside"},
		{Tool: "t5", Title: "x", Command: "run"},
	}
	got := Query(cheats, "needle")
	if len(got) != 4 {
		t.Fatalf("matches = %v, want 4 (tag/title/command/desc)", tools(got))
	}
	// Field weights order the results: tag > title > command > desc.
	want := []string{"t1", "t2", "t3", "t4"}
	for i := range want {
		if got[i].Tool != want[i] {
			t.Fatalf("order = %v, want %v", tools(got), want)
		}
	}
}
