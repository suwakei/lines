package pathHandler

import (
	"testing"
)

func TestSearch(t *testing.T) {
	// テスト用のパスと無視リストを設定
	testPath := "test_directory"
	ignores := map[string][]string{
		"dir":  {"ignored_dir"},
		"file": {"ignored_file.txt"},
	}

	// Search関数を呼び出し、結果を検証
	files, err := Search(testPath, ignores)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// 期待されるファイルリストを検証
	expectedFiles := []string{"file1.txt", "file2.txt"} // 期待されるファイル名を設定
	if len(files) != len(expectedFiles) {
		t.Fatalf("expected %v files, got %v", len(expectedFiles), len(files))
	}
	// 追加の検証を行うことができます
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