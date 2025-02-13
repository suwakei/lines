package pathHandler

import (
	"os"
	"fmt"
	fp "path/filepath"
)


func Parse(path string) (string, error) {
	path, err := fp.Abs(path)
	if err != nil {
		return "", err
	}
	if !Exists(path) {
		return "", fmt.Errorf("[INFO]: path does not exist \" %s \"", path)
	}
	return fp.Clean(path), nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}