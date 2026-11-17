// Package loader reads cheat files from the embedded filesystem and
// converts them into Cheat objects.
package loader

import (
	"strings"
	"testing"
)

// =============================================================================
// Optional Markers
// =============================================================================

// TestOptionalPlaceholdersKept verifies documented optional args stay optional.
// A past refactor silently dropped the "|" marker, turning valid no-arg
// commands into blocked required-arg forms.
func TestOptionalPlaceholdersKept(t *testing.T) {
	cases := []struct {
		file string
		want []string
	}{
		{"cheat-files/npm.yaml", []string{"package_name|", "force_flag|"}},
		{"cheat-files/impacket-goldenpac.yaml", []string{"binary_arguments|"}},
		{"cheat-files/impacket-mssqlclient.yaml", []string{"domain|"}},
	}
	for _, tc := range cases {
		cheats, err := parseCheatFile(embeddedCheats, tc.file)
		if err != nil {
			t.Fatal(err)
		}
		joined := ""
		for _, c := range cheats {
			joined += c.Command + "\n"
		}
		for _, want := range tc.want {
			if !strings.Contains(joined, "{{"+want+"}}") {
				t.Errorf("%s: missing optional {{%s}}", tc.file, want)
			}
		}
	}
}
