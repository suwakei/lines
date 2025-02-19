package pathHandler

import (
	"io/fs"
	fp "path/filepath"
	"slices"
)

func Search(path string, ignores map[string][]string) ([]string, error) {
	var files []string
  ignoreSetDir map[string]struct{} = make(map[string]struct{}, len(ignores["dir"]))

        for _, ignore := range ignores["dir"] {
                ignoreSetDir[ignore] = struct{}{}
        }

ignoreSetFile map[string]struct{} = make(map[string]struct{}, len(ignores["file"]))

for _, ignore := range ignores["file"] {
                ignoreSetFile[ignore] = struct{}{}
        }


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

func contains(ignores []string, pathBaseName string) bool {
	if ignores == nil {
		return false
	}
	
	_, exist := ignoreSet[pathBaseName]
	return exist
}

func isInvalidFile(pathBaseName string) bool {
	ext := fp.Ext(pathBaseName)
	if ext == "" {
		if pathBaseName != "Makefile" && pathBaseName != "Dockerfile" {
				return true
			}
	}
	return false
}

