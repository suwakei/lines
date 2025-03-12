package pathHandler

import (
	"fmt"
	"os"
	fp "path/filepath"
)

func Parse(path string) (string, error) {
	absPath, err := fp.Abs(path)
	if err != nil {
		return "", err
	}
	if !Exists(absPath) {
		return "", fmt.Errorf("[INFO]: path does not exist \" %s \"", absPath)
	}
	return fp.Clean(absPath), nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
