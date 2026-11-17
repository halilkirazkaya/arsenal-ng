// Package model defines the core data types for arsenal-ng.
package model

import (
	"reflect"
	"testing"
)

// =============================================================================
// Validation Tests
// =============================================================================

// TestEmptyArgs verifies required-argument emptiness detection.
func TestEmptyArgs(t *testing.T) {
	required := []Argument{{Name: "ip", Value: ""}}
	optional := []Argument{{Name: "d", Value: "", Optional: true}}
	filled := []Argument{{Name: "ip", Value: "x"}}

	if !HasEmptyArgs(required) {
		t.Fatal("HasEmptyArgs(required empty) = false")
	}
	if HasEmptyArgs(optional) {
		t.Fatal("HasEmptyArgs(optional empty) = true")
	}
	if HasEmptyArgs(filled) {
		t.Fatal("HasEmptyArgs(filled) = true")
	}
	if HasEmptyArgs(nil) {
		t.Fatal("HasEmptyArgs(nil) = true")
	}
	if got := GetEmptyArgs([]Argument{
		{Name: "ip", Value: ""},
		{Name: "d", Value: "", Optional: true},
		{Name: "p", Value: "80"},
	}); !reflect.DeepEqual(got, []string{"ip"}) {
		t.Fatalf("GetEmptyArgs = %v, want [ip]", got)
	}
	if got := GetEmptyArgs(filled); len(got) != 0 {
		t.Fatalf("GetEmptyArgs(filled) = %v, want empty", got)
	}
}
