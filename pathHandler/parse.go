package pathHandler

import (
	fp "path/filepath"
)

func Parse(path string) (string, error) {
	absPath, err := fp.Abs(path)
	if err != nil {
		return "", err
	}
	return fp.Clean(absPath), nil
}
