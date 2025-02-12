package handler

import (
	"io/fs"
	fp "path/filepath"
)

func Search(path string, ignores []string) ([]string, error) {
	var files []string
	path, err := Parse(path)
	if err != nil {
		return nil, err
	}

	fp.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	return files, nil
}
