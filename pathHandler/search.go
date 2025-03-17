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
		pathBase := fp.Base(path)
		if err != nil {
			return err
		}

		if d.IsDir() {
			if contains(pathBase, ignores["dir"]) {
				return fp.SkipDir
			}
			return nil
		}

		if !d.IsDir() && isInvalidFile(pathBase) {
			return nil
		}

		if !d.IsDir() && contains(pathBase, ignores["file"]) {
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


func contains(pathBase string, ignores []string) bool {
	pathExt := fp.Ext(pathBase)
	if pathExt == "" {
		return slices.Contains(ignores, pathBase)
	}
	return slices.Contains(ignores, pathExt)
}

func isInvalidFile(pathBase string) bool {
	ext := fp.Ext(pathBase)
	if ext == "" && pathBase != "Makefile" && pathBase != "Dockerfile" {
		return true
	}
	return false
}
