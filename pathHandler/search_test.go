package pathHandler

import (
	"fmt"
	"path/filepath"
	"slices"
	"testing"
)

func TestSearch(t *testing.T) {
	// configure path for testing
	testPathAbs, _ := filepath.Abs("../testdata/linestest")
	testPath := filepath.Clean(testPathAbs)
	ignores := map[string][]string{
		"dir":  {filepath.Join(testPath, "ignore_dir")},
		"file": {filepath.Join(testPath, "ignore.txt")},
	}

	// call Search function and verify result
	files, err := Search(testPath, ignores)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// verify expected fileList
	expectedFiles := mapString([]string{"Dockerfile", "Makefile", "test.go", "view.txt", "テスト.txt"})
	for _, f := range files {
		if !slices.Contains(expectedFiles, f) {
			t.Errorf("%s is not expected file", f)
		}
	}
}

func TestContains(t *testing.T) {
	ignores := []string{"ignored.txt", ".log"}

	if !contains("ignored.txt", ignores) {
		t.Errorf("expected true for ignored.txt")
	}
	if !contains("file.log", ignores) {
		t.Errorf("expected true for file.log")
	}
	if contains("file.txt", ignores) {
		t.Errorf("expected false for file.txt")
	}
}

func TestIsInvalidFile(t *testing.T) {
	if !isInvalidFile("invalid_file.txt") {
		t.Errorf("expected true for invalid_file.txt")
	}
	if isInvalidFile("Makefile") {
		t.Errorf("expected false for Makefile")
	}
	if isInvalidFile("Dockerfile") {
		t.Errorf("expected false for Dockerfile")
	}
	if isInvalidFile("LICENSE") {
		t.Errorf("expected false for LICENSE")
	}
}

func mapString(s []string) []string {
	resS := make([]string, 0, len(s))

	for _, S := range s {
		abs, _ := filepath.Abs(fmt.Sprintf("../testdata/linestest/%s", S))
		resS = append(resS, filepath.Clean(abs))
	}
	return resS
}