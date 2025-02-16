package pathHandler

import (
	"io/fs"
	fp "path/filepath"
	"slices"
)

func Search(path string, ignores map[string][]string) ([]string, error) {
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
			if contains(ignores["dir"], path) {
				return fp.SkipDir
			} else {
				return nil
			}
		}

		if !d.IsDir() && contains(ignores["file"], path) {
			return nil
		}

		files = append(files, path)
		return nil
	})
	return files, nil
}

func contains(ignores []string, path string) bool {
	return slices.Contains(ignores, fp.Base(path))
}

