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
		{".c", "//", true},
		{".c", "// some comment", true},
		{".c", "not a comment", false},
		{".cc", "//", true},
		{".cc", "// some comment", true},
		{".cc", "not a comment", false},
		{".cs", "//", true},
		{".cs", "// some comment", true},
		{".cs", "not a comment", false},
		{".cpp", "//", true},
		{".cpp", "// some comment", true},
		{".cpp", "not a comment", false},
		{".h", "//", true},
		{".h", "// some comment", true},
		{".h", "not a comment", false},
		{".hpp", "//", true},
		{".hpp", "// some comment", true},
		{".hpp", "not a comment", false},
		{".d", "//", true},
		{".d", "// some comment", true},
		{".d", "not a comment", false},
		{".dart", "//", true},
		{".dart", "// some comment", true},
		{".dart", "not a comment", false},
		{".groovy", "//", true},
		{".groovy", "// some comment", true},
		{".groovy", "not a comment", false},
		{".js", "//", true},
		{".js", "// some comment", true},
		{".js", "not a comment", false},
		{".ts", "//", true},
		{".ts", "// some comment", true},
		{".ts", "not a comment", false},
		{".jsx", "//", true},
		{".jsx", "// some comment", true},
		{".jsx", "not a comment", false},
		{".tsx", "//", true},
		{".tsx", "// some comment", true},
		{".tsx", "not a comment", false},
		{".java", "//", true},
		{".java", "// some comment", true},
		{".java", "not a comment", false},
		{".jsonc", "//", true},
		{".jsonc", "// some comment", true},
		{".jsonc", "not a comment", false},
		{".kt", "//", true},
		{".kt", "// some comment", true},
		{".kt", "not a comment", false},
		{".m", "//", true},
		{".m", "// some comment", true},
		{".m", "not a comment", false},
		{".php", "//", true},
		{".php", "// some comment", true},
		{".php", "not a comment", false},
		{".rs", "not a comment", false},
		{".rs", "//", true},
		{".rs", "// some comment", true},
		{"", "//", false},
		{".d", "not a comment", false},
		{".d", "///", true},
		{".d", "/// some comment", true},
		{".dart", "not a comment", false},
		{".dart", "///", true},
		{".dart", "/// some comment", true},
		{".rs", "/// some comment", true},
		{".rs", "///", true},
		{"", "///", false},
		{".rs", "//!", true},
		{".rs", "//! some comment", true},
		{"", "//!", false},
		{".bash", "#", true},
		{".bash", "# some comment", true},
		{".bash", "not a comment", false},
		{".cfg", "#", true},
		{".cfg", "# some comment", true},
		{".cfg", "not a comment", false},
		{".coffee", "#", true},
		{".coffee", "# some comment", true},
		{".coffee", "not a comment", false},
		{".dockerfile", "not a comment", false},
		{".dockerfile", "#", true},
		{".dockerfile", "# some comment", true},
		{".Dockerfile", "not a comment", false},
		{".Dockerfile", "#", true},
		{".Dockerfile", "# some comment", true},
		{"Dockerfile", "not a comment", false},
		{"Dockerfile", "#", true},
		{"Dockerfile", "# some comment", true},
		{".dockerignore", "not a comment", false},
		{".dockerignore", "#", true},
		{".dockerignore", "# some comment", true},
		{".ex", "not a comment", false},
		{".ex", "#", true},
		{".ex", "# some comment", true},
		{".gitignore", "not a comment", false},
		{".gitignore", "#", true},
		{".gitignore", "# some comment", true},
		{".mk", "not a comment", false},
		{".mk", "#", true},
		{".mk", "# some comment", true},
		{"Makefile", "not a comment", false},
		{"Makefile", "#", true},
		{"Makefile", "# some comment", true},
		{".py", "not a comment", false},
		{".py", "#", true},
		{".py", "# some comment", true},
		{".pl", "not a comment", false},
		{".pl", "#", true},
		{".pl", "# some comment", true},
		{".php", "not a comment", false},
		{".php", "#", true},
		{".php", "# some comment", true},
		{".rb", "not a comment", false},
		{".rb", "#", true},
		{".rb", "# some comment", true},
		{"", "#", false},
		{".nim", "not a comment", false},
		{".nim", "##", true},
		{".nim", "## some comment", true},
		{"", "##", false},
		{".f90", "not a comment", false},
		{".f90", "!", true},
		{".f90", "! some comment", true},
		{"", "!", false},
		{".lua", "not a comment", false},
		{".lua", "--", true},
		{".lua", "-- some comment", true},
		{"", "--", false},
		{".erl", "not a comment", false},
		{".erl", "%", true},
		{".erl", "% some comment", true},
		{"", "%", false},
		{".asm", "not a comment", false},
		{".asm", ";", true},
		{".asm", "; some comment", true},
		{".clj", "not a comment", false},
		{".clj", ";", true},
		{".clj", "; some comment", true},
		{".ini", "not a comment", false},
		{".ini", ";", true},
		{".ini", "; some comment", true},
		{"", ";", false},
		{".bat", "not a comment", false},
		{".bat", "rem", true},
		{".bat", "rem some comment", true},
		{"", "rem", false},
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
		{".c", "/*", true},
		{".c", "/* some comment", true},
		{".c", "not a comment", false},
		{".cc", "/*", true},
		{".cc", "/* some comment", true},
		{".cc", "not a comment", false},
		{".cs", "/*", true},
		{".cs", "/* some comment", true},
		{".cs", "not a comment", false},
		{".cpp", "/*", true},
		{".cpp", "/* some comment", true},
		{".cpp", "not a comment", false},
		{".h", "/*", true},
		{".h", "/* some comment", true},
		{".h", "not a comment", false},
		{".hpp", "/*", true},
		{".hpp", "/* some comment", true},
		{".hpp", "not a comment", false},
		{".d", "/*", true},
		{".d", "/* some comment", true},
		{".d", "not a comment", false},
		{".dart", "/*", true},
		{".dart", "/* some comment", true},
		{".dart", "not a comment", false},
		{".groovy", "/*", true},
		{".groovy", "/* some comment", true},
		{".groovy", "not a comment", false},
		{".js", "/*", true},
		{".js", "/* some comment", true},
		{".js", "not a comment", false},
		{".ts", "/*", true},
		{".ts", "/* some comment", true},
		{".ts", "not a comment", false},
		{".jsx", "/*", true},
		{".jsx", "/* some comment", true},
		{".jsx", "not a comment", false},
		{".tsx", "/*", true},
		{".tsx", "/* some comment", true},
		{".tsx", "not a comment", false},
		{".java", "/*", true},
		{".java", "/* some comment", true},
		{".java", "not a comment", false},
		{".jsonc", "/*", true},
		{".jsonc", "/* some comment", true},
		{".jsonc", "not a comment", false},
		{".kt", "/*", true},
		{".kt", "/* some comment", true},
		{".kt", "not a comment", false},
		{".m", "/*", true},
		{".m", "/* some comment", true},
		{".m", "not a comment", false},
		{".php", "/*", true},
		{".php", "/* some comment", true},
		{".php", "not a comment", false},
		{".rs", "not a comment", false},
		{".rs", "/*", true},
		{".rs", "/* some comment", true},
		{"", "/*", false},
		{".d", "/++", true},
		{".d", "/++ some comment", true},
		{"", "/++", false},
		{".rs", "/*!", true},
		{".rs", "/*! some comment", true},
		{".d", "not a comment", false},
		{".d", "/**", true},
		{".d", "/** some comment", true},
		{".kt", "not a comment", false},
		{".kt", "/**", true},
		{".kt", "/** some comment", true},
		{".m", "not a comment", false},
		{".m", "/**", true},
		{".m", "/** some comment", true},
		{"", "/**", false},
		{".pas", "{", true},
		{".pas", "{ some comment", true},
		{".pas", "not a comment", false},
		{"", "{", false},
		{".html", "<!--", true},
		{".html", "<!-- some comment", true},
		{".html", "not a comment", false},
		{".xml", "<!--", true},
		{".xml", "<!-- some comment", true},
		{".xml", "not a comment", false},
		{".md", "<!--", true},
		{".md", "<!-- some comment", true},
		{".md", "not a comment", false},
		{"", "<!--", false},
		{".pl", "=pod", true},
		{".pl", "=pod some comment", true},
		{".pl", "not a comment", false},
		{"", "=pod", false},
		{".rb", "=begin", true},
		{".rb", "=begin some comment", true},
		{".rb", "not a comment", false},
		{"", "=begin", false},
		{".lua", "--[[", true},
		{".lua", "--[[ some comment", true},
		{".lua", "not a comment", false},
		{"", "--[[", false},
		{".ex", "\"\"\"", true},
		{".ex", "\"\"\" some comment", true},
		{".ex", "not a comment", false},
		{".py", "\"\"\"", true},
		{".py", "\"\"\" some comment", true},
		{".py", "not a comment", false},
		{"", "\"\"\"", false},
		{".coffee", "###", true},
		{".coffee", "### some comment", true},
		{".coffee", "not a comment", false},
		{"", "#)", false},
		{".fs", "(#", true},
		{".fs", "(# some comment", true},
		{".fs", "not a comment", false},
		{"", "(#", false},
	}

	for _, test := range tests {
		fi := FileInfo{FileType: test.fileType}
		result := fi.isBeginBlockComments(test.line)
		if result != test.expected {
			t.Errorf("Expected %v for fileType %s and line %q, got %v", test.expected, test.fileType, test.line, result)
		}
	}
}

func TestIsEndBlockComments(t *testing.T) {
	tests := []struct {
		fileType string
		line     string
		expected bool
	}{
		{".c", "*/", true},
		{".c", "*/ some comment", true},
		{".c", "not a comment", false},
		{".cc", "*/", true},
		{".cc", "*/ some comment", true},
		{".cc", "not a comment", false},
		{".cs", "*/", true},
		{".cs", "*/ some comment", true},
		{".cs", "not a comment", false},
		{".cpp", "*/", true},
		{".cpp", "*/ some comment", true},
		{".cpp", "not a comment", false},
		{".h", "*/", true},
		{".h", "*/ some comment", true},
		{".h", "not a comment", false},
		{".hpp", "*/", true},
		{".hpp", "*/ some comment", true},
		{".hpp", "not a comment", false},
		{".d", "*/", true},
		{".d", "*/ some comment", true},
		{".d", "not a comment", false},
		{".dart", "*/", true},
		{".dart", "*/ some comment", true},
		{".dart", "not a comment", false},
		{".groovy", "*/", true},
		{".groovy", "*/ some comment", true},
		{".groovy", "not a comment", false},
		{".js", "*/", true},
		{".js", "*/ some comment", true},
		{".js", "not a comment", false},
		{".ts", "*/", true},
		{".ts", "*/ some comment", true},
		{".ts", "not a comment", false},
		{".jsx", "*/", true},
		{".jsx", "*/ some comment", true},
		{".jsx", "not a comment", false},
		{".tsx", "*/", true},
		{".tsx", "*/ some comment", true},
		{".tsx", "not a comment", false},
		{".java", "*/", true},
		{".java", "*/ some comment", true},
		{".java", "not a comment", false},
		{".jsonc", "*/", true},
		{".jsonc", "*/ some comment", true},
		{".jsonc", "not a comment", false},
		{".kt", "*/", true},
		{".kt", "*/ some comment", true},
		{".kt", "not a comment", false},
		{".m", "*/", true},
		{".m", "*/ some comment", true},
		{".m", "not a comment", false},
		{".php", "*/", true},
		{".php", "*/ some comment", true},
		{".php", "not a comment", false},
		{".rs", "not a comment", false},
		{".rs", "*/", true},
		{".rs", "*/ some comment", true},
		{"", "*/", false},
		{".d", "+/", true},
		{".d", "+/ some comment", true},
		{"", "+/", false},
		{".pas", "}", true},
		{".pas", "} some comment", true},
		{".pas", "not a comment", false},
		{"", "}", false},
		{".html", "-->", true},
		{".html", "--> some comment", true},
		{".html", "not a comment", false},
		{".xml", "-->", true},
		{".xml", "--> some comment", true},
		{".xml", "not a comment", false},
		{".md", "-->", true},
		{".md", "--> some comment", true},
		{".md", "not a comment", false},
		{"", "-->", false},
		{".pl", "=cut", true},
		{".pl", "=cut some comment", true},
		{".pl", "not a comment", false},
		{"", "=cut", false},
		{".rb", "=end", true},
		{".rb", "=end some comment", true},
		{".rb", "not a comment", false},
		{"", "=end", false},
		{".lua", "]]", true},
		{".lua", "]] some comment", true},
		{".lua", "not a comment", false},
		{"", "]]", false},
		{".ex", "\"\"\"", true},
		{".ex", "\"\"\" some comment", true},
		{".ex", "not a comment", false},
		{".py", "\"\"\"", true},
		{".py", "\"\"\" some comment", true},
		{".py", "not a comment", false},
		{"", "\"\"\"", false},
		{".coffee", "###", true},
		{".coffee", "### some comment", true},
		{".coffee", "not a comment", false},
		{"", "#)", false},
		{".fs", "#)", true},
		{".fs", "#) some comment", true},
		{".fs", "not a comment", false},
		{"", "#)", false},
	}

	for _, test := range tests {
		fi := FileInfo{FileType: test.fileType}
		result := fi.isEndBlockComments(test.line)
		if result != test.expected {
			t.Errorf("Expected %v for fileType %s and line %q, got %v", test.expected, test.fileType, test.line, result)
		}
	}
}
