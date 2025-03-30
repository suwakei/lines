package counter

import (
	"sync"
	"testing"
)

var files = []string{
	"../testdata/linestest/Dockerfile",
	"../testdata/linestest/Makefile",
	"../testdata/linestest/test.go",
	"../testdata/linestest/view.txt",
	"../testdata/linestest/テスト.txt",
}

// func BenchmarkCount(b *testing.B) {
// 	inputPath := "./testdata/bench"

// 	for i := 0; i < b.N; i++ {
// 		_, err := Count(files, inputPath)
// 		if err != nil {
// 			b.Fatalf("Expected no error, got %v", err)
// 		}
// 	}
// }

func TestCount(t *testing.T) {
	inputPath := "../testdata/linestest"
	expectedLines := 90
	expectedBlanks := 12
	expectedComments := 6
	expectedFiles := 5
	expectedBytes := "3460(3.5 KB)"

	result, err := Count(files, inputPath)
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}

	if result.TotalLines != expectedLines {
		t.Errorf("expected %d lines, got %d", expectedLines, result.TotalLines)
	}
	if result.TotalBlanks != expectedBlanks {
		t.Errorf("expected %d blanks, got %d", expectedBlanks, result.TotalBlanks)
	}
	if result.TotalComments != expectedComments {
		t.Errorf("expected %d comments, got %d", expectedComments, result.TotalComments)
	}
	if result.TotalFiles != expectedFiles {
		t.Errorf("expected %d files, got %d", expectedFiles, result.TotalFiles)
	}
	if result.TotalBytes != expectedBytes {
		t.Errorf("expected %s bytes, got %s", expectedBytes, result.TotalBytes)
	}
}

func TestProcessFile(t *testing.T) {
	file := "../testdata/dummy/view.txt"
	bufMap := make(map[string]*FileInfo)
	mu := &sync.Mutex{}

	err := processFile(file, bufMap, mu)
	if err != nil {
		t.Fatalf("processFile failed: %v", err)
	}

	if len(bufMap) == 0 {
		t.Error("bufMap is empty, expected at least one entry")
	}
}

func TestRetFileType(t *testing.T) {
	file := "../testdata/dummy/view.go"
	expected := ".go"

	result := retFileType(file)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	file = "../testdata/dummy/Makefile"
	expected = "Makefile"

	result = retFileType(file)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestIsSingleComment(t *testing.T) {
	tests := []struct {
		fileType string
		line     string
		expected bool
	}{
		{".go", "// This is a comment", true},
		{".go", "fmt.Println(\"Hello, World!\")", false},
		{".cpp", "// Another comment", true},
		{".cpp", "int main() {}", false},
		{".js", "// JavaScript comment", true},
		{".js", "console.log('Hello');", false},
		{".py", "# Python comment", true},
		{".py", "print('Hello')", false},
		{".java", "// Java comment", true},
		{".java", "System.out.println(\"Hello\");", false},
		{".rb", "# Ruby comment", true},
		{".rb", "puts 'Hello'", false},
		{".cs", "// C# comment", true},
		{".cs", "Console.WriteLine(\"Hello\");", false},
		{".html", "<!-- HTML comment -->", false}, // HTML is not single comment.
	}

	for _, test := range tests {
		fi := FileInfo{FileType: test.fileType}
		result := fi.isSingleComment(test.line)
		if result != test.expected {
			t.Errorf("Expected %v for line %q in file type %q, got %v", test.expected, test.line, test.fileType, result)
		}
	}
}

func TestIsBeginBlockComments(t *testing.T) {
	tests := []struct {
		fileType string
		line     string
		expected bool
	}{
		{".go", "/* This is a block comment", true},
		{".cpp", "/* Start of comment", true},
		{".java", "// This is a single line comment", false},
		{".c", "/* Block comment starts here", true},
		{".js", "// Another comment", false},
		{".html", "<!-- This is an HTML comment", true},
		{".py", "# This is a Python comment", false},
		{".rs", "/* Rust block comment", true},
		{".txt", "Just a plain text line", false},
	}

	for _, test := range tests {
		fi := FileInfo{FileType: test.fileType}
		result := fi.isBeginBlockComments(test.line)
		if result != test.expected {
			t.Errorf("For fileType %q and line %q, expected %v but got %v", test.fileType, test.line, test.expected, result)
		}
	}
}

func TestIsEndBlockComments(t *testing.T) {
	tests := []struct {
		fileType string
		line     string
		expected bool
	}{
		{"go", "*/", true},
		{"go", "*/ some comment", true},
		{"go", "not a comment", false},
		{"cpp", "*/", true},
		{"cpp", "*/ some comment", true},
		{"cpp", "not a comment", false},
		{"java", "*/", true},
		{"java", "*/ some comment", true},
		{"java", "not a comment", false},
		{"", "*/", false}, // ファイルタイプが空の場合
	}

	for _, test := range tests {
		fi := FileInfo{FileType: test.fileType}
		result := fi.isEndBlockComments(test.line)
		if result != test.expected {
			t.Errorf("Expected %v for fileType %s and line %q, got %v", test.expected, test.fileType, test.line, result)
		}
	}
}
