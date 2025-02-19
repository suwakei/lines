package pathHandler

import (
	"io/fs"
	fp "path/filepath"
	"slices"
)

func Search(path string, ignores map[string][]string) ([]string, error) {
	var files []string

	parsedPath, err := Parse(path)
	if err != nil {
		return nil, err
	}

	err = fp.WalkDir(parsedPath, func(path string, d fs.DirEntry, err error) error {
		base := fp.Base(path)
		if err != nil {
			return err
		}

		if d.IsDir() {
			if contains(ignores["dir"], base) {
				return fp.SkipDir
			}
			return nil
		}

		if !d.IsDir() {
			if contains(ignores["file"], base) {
				return nil
			}
		}

		if isInvalidFile(base) {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func contains(ignores []string, base string) bool {
	return slices.Contains(ignores, base)
}

func isInvalidFile(base string) bool {
	ext := fp.Ext(base)
	if ext == "" {
		if base != "Makefile" || base != "Dockerfile" {
				return true
			}
	}
	return false
}

