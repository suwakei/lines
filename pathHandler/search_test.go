package pathHandler

import (
    "io/fs"
    "os"
    "path/filepath"
    "reflect"
    "testing"
)

func TestSearch(t *testing.T) {
    // テスト用のディレクトリとファイルを作成
    rootDir := "testdata"
    subDir := filepath.Join(rootDir, "subdir")
    testFiles := []string{
        filepath.Join(rootDir, "file1.txt"),
        filepath.Join(rootDir, "file2.txt"),
        filepath.Join(subDir, "file3.txt"),
        filepath.Join(subDir, "file4.txt"),
    }
    os.MkdirAll(subDir, 0755)
    for _, file := range testFiles {
        os.WriteFile(file, []byte("test content"), 0644)
    }
    defer os.RemoveAll(rootDir)

    ignores := map[string][]string{
        "file": {"file2.txt"},
        "dir":  {"subdir"},
    }

    expected := []string{
        filepath.Join(rootDir, "file1.txt"),
    }

    result, err := Search(rootDir, ignores)
    if err != nil {
        t.Fatalf("Search returned an error: %v", err)
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestContains(t *testing.T) {
    ignores := []string{"ignore1", "ignore2"}
    if !contains(ignores, "ignore1") {
        t.Error("Expected true, got false")
    }
    if contains(ignores, "notignored") {
        t.Error("Expected false, got true")
    }
}

func TestIsInvalidFile(t *testing.T) {
    if !isInvalidFile("invalidfile") {
        t.Error("Expected true, got false")
    }
    if isInvalidFile("validfile.txt") {
        t.Error("Expected false, got true")
    }
    if isInvalidFile("Makefile") {
        t.Error("Expected false, got true")
    }
}