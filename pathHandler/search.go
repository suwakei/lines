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
		if err != nil {
			return err
		}

		if d.IsDir() {
			if contains(ignores["dir"], path) {
				return fp.SkipDir
			}
			return nil
		}

		if contains(ignores["file"], path) {
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

func contains(ignores []string, path string) bool {
	return slices.Contains(ignores, fp.Base(path))
}
