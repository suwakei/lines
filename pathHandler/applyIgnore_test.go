package pathHandler

import (
    "os"
    "reflect"
    "testing"
)

func TestMakeIgnoreList_String(t *testing.T) {
    // テスト用の.gitignoreファイルを作成
    ignoreFileContent := "testfile.txt\nignoredir/\n# This is a comment\n"
    ignoreFilePath := "./testdata/.gitignore"
    err := os.WriteFile(ignoreFilePath, []byte(ignoreFileContent), 0644)
    if err != nil {
        t.Fatalf("Failed to create test ./testdata/.gitignore file: %v", err)
    }
    defer os.Remove(ignoreFilePath)

    expected := map[string][]string{
        "file": {"testfile.txt"},
        "dir":  {"ignoredir/"},
    }

    result, err := MakeIgnoreList(ignoreFilePath)
    if err != nil {
        t.Errorf("MakeIgnoreList returned an error: %v", err)
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestMakeIgnoreList_String_Invalid(t *testing.T) {
    _, err := MakeIgnoreList("invalidfile")
    if err == nil {
        t.Error("Expected an error for invalid file name, but got none")
    }
}

func TestMakeIgnoreList_Slice(t *testing.T) {
    ignores := []string{"testfile.txt", "ignoredir/"}
    expected := map[string][]string{
        "file": {"testfile.txt"},
        "dir":  {"ignoredir/"},
    }

    result, err := MakeIgnoreList(ignores)
    if err != nil {
        t.Errorf("MakeIgnoreList returned an error: %v", err)
    }

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

func TestMakeIgnoreList_Slice_Empty(t *testing.T) {
    _, err := MakeIgnoreList([]string{})
    if err == nil {
        t.Error("Expected an error for empty slice, but got none")
    }
}

func TestIsFile(t *testing.T) {
    // テスト用のファイルを作成
    fileName := "./testdata/testfile.txt"
    err := os.WriteFile(fileName, []byte("test content"), 0644)
    if err != nil {
        t.Fatalf("Failed to create test file: %v", err)
    }
    defer os.Remove(fileName)

    if !IsFile(fileName) {
        t.Errorf("Expected %s to be a file, but it was not", fileName)
    }

    // ディレクトリをテスト
    dirName := "./testdata/testdir"
    err = os.Mkdir(dirName, 0755)
    if err != nil {
        t.Fatalf("Failed to create test directory: %v", err)
    }
    defer os.Remove(dirName)

    if IsFile(dirName) {
        t.Errorf("Expected %s to be a directory, but it was not", dirName)
    }
}