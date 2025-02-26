package counter

import (
    "sync"
    "testing"
)

func TestCount(t *testing.T) {
    files := []string{"testfile1.go", "testfile2.go"}
    inputPath := "testdata"
    expectedSteps := 10
    expectedBlanks := 2
    expectedComments := 3
    expectedBytes := int64(100)

    result, err := Count(files, inputPath)
    if err != nil {
        t.Fatalf("Count failed: %v", err)
    }

    if result.AllSteps != expectedSteps {
        t.Errorf("expected %d steps, got %d", expectedSteps, result.AllSteps)
    }
    if result.AllBlanks != expectedBlanks {
        t.Errorf("expected %d blanks, got %d", expectedBlanks, result.AllBlanks)
    }
    if result.AllComments != expectedComments {
        t.Errorf("expected %d comments, got %d", expectedComments, result.AllComments)
    }
    if result.AllBytes != expectedBytes {
        t.Errorf("expected %d bytes, got %d", expectedBytes, result.AllBytes)
    }
}

func TestProcessFile(t *testing.T) {
    file := "testfile.go"
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
    file := "testfile.go"
    expected := ".go"

    result := retFileType(file)
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