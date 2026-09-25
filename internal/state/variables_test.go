package state

import (
	"os"
	"path/filepath"
	"testing"
)

func newTestGlobal(t *testing.T) *Global {
	t.Helper()
	return &Global{
		variables: make(map[string]string),
		filePath:  filepath.Join(t.TempDir(), "variables.json"),
	}
}

func assertOwnerOnly(t *testing.T, path string) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if mode := info.Mode().Perm(); mode != 0600 {
		t.Fatalf("%s has mode %o, want 600", path, mode)
	}
}

func TestSaveWritesVariablesReadableByOwnerOnly(t *testing.T) {
	g := newTestGlobal(t)

	if err := g.Set("password", "Winter2026!"); err != nil {
		t.Fatal(err)
	}

	assertOwnerOnly(t, g.filePath)
}

func TestSaveTightensAnOldWorldReadableFile(t *testing.T) {
	g := newTestGlobal(t)
	if err := os.WriteFile(g.filePath, []byte(`{"ip":"10.0.0.1"}`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(g.filePath+".tmp", []byte("stale"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := g.Set("password", "Winter2026!"); err != nil {
		t.Fatal(err)
	}

	assertOwnerOnly(t, g.filePath)
}

func TestLoadingANullFileLeavesAUsableStore(t *testing.T) {
	g := newTestGlobal(t)
	if err := os.WriteFile(g.filePath, []byte("null"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := g.LoadFromFile(); err != nil {
		t.Fatal(err)
	}

	if err := g.Set("ip", "10.0.0.1"); err != nil {
		t.Fatal(err)
	}

	if v, ok := g.Get("ip"); !ok || v != "10.0.0.1" {
		t.Fatalf("Get(ip) = %q, %v", v, ok)
	}
}
