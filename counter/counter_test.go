package counter

import (
	"sync"
	"testing"
)

var files = []string{
	"../testdata/dummy/dummy.txt",
	"../testdata/dummy/dummy2.txt",
	"../testdata/dummy/dummy3.txt",
	"../testdata/dummy/dummy4.log",
	"../testdata/dummy/Makefile",
	"../testdata/dummy/view.go",
	"../testdata/dummy/view.txt",
	"../testdata/dummy/a.exe",
}

func BenchmarkCount(b *testing.B) {
	inputPath := "../testdata/dummy"

	for i := 0; i < b.N; i++ {
		_, err := Count(files, inputPath)
		if err != nil {
			b.Fatalf("Expected no error, got %v", err)
		}
	}
}

func TestCount(t *testing.T) {
	inputPath := "../testdata/dummy"
	expectedSteps := 39
	expectedBlanks := 4
	expectedComments := 1
	expectedBytes := "1600(1KB)"

	result, err := Count(files, inputPath)
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}

	if result.TotalSteps != expectedSteps {
		t.Errorf("expected %d steps, got %d", expectedSteps, result.TotalSteps)
	}
	if result.TotalBlanks != expectedBlanks {
		t.Errorf("expected %d blanks, got %d", expectedBlanks, result.TotalBlanks)
	}
	if result.TotalComments != expectedComments {
		t.Errorf("expected %d comments, got %d", expectedComments, result.TotalComments)
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
	line := "// this is a comment"
	if !isSingleComment(line) {
		t.Error("expected true, got false")
	}

	line = "this is not a comment"
	if isSingleComment(line) {
		t.Error("expected false, got true")
	}
}

func TestIsBeginBlockComments(t *testing.T) {
	line := "/* this is a block comment */"
	if !isBeginBlockComments(line) {
		t.Error("expected true, got false")
	}

	line = "this is not a block comment"
	if isBeginBlockComments(line) {
		t.Error("expected false, got true")
	}
}

func TestIsEndBlockComments(t *testing.T) {
	line := "*/"
	if !isEndBlockComments(line) {
		t.Error("expected true, got false")
	}

	line = "not an end comment"
	if isEndBlockComments(line) {
		t.Error("expected false, got true")
	}
}
