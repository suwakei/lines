package pathHandler

import (
	"io/fs"
	fp "path/filepath"
)

func Search(path string, ignores map[string][]string) ([]string, error) {
	var files []string

	parsedPath, err := Parse(path)
	if err != nil {
		return nil, err
	}

	err = fp.WalkDir(parsedPath, func(path string, d fs.DirEntry, err error) error {
		pathBaseName := fp.Base(path)
		if err != nil {
			return err
		}

		if d.IsDir() {
			if contains(ignores["dir"], pathBaseName) {
				return fp.SkipDir
			}
			return nil
		}

		if !d.IsDir() {
			if contains(ignores["file"], pathBaseName) {
				return nil
			}
		}

		if isInvalidFile(pathBaseName) {
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

var ignoreSet map[string]struct{} = make(map[string]struct{})
func contains(ignores []string, pathBaseName string) bool {
	if ignores == nil {
		return false
	}
	for _, ignore := range ignores {
		ignoreSet[ignore] = struct{}{}
	}
	_, exist := ignoreSet[pathBaseName]
	return exist
}

func isInvalidFile(pathBaseName string) bool {
	ext := fp.Ext(pathBaseName)
	if ext == "" && pathBaseName != "Makefile" && pathBaseName != "Dockerfile" {
		return true
	}
	return false
}
