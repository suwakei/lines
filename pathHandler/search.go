package pathHandler

import (
	"io/fs"
	//"path/filepath"
	fp "path/filepath"
	//"slices"
	"regexp"
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
			if contains(path, ignores["dir"]) {
				return fp.SkipDir
			}
			return nil
		}

		if !d.IsDir() && isInvalidFile(path) {
			return nil
		}

		if !d.IsDir() && contains(path, ignores["file"]) {
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

func contains(path string, ignores []string) bool {
	// if slices.Contains(ignores, path) {
	// 	return true
	// }
	// pathExt := fp.Ext(path)
	// if pathExt == "" {
	// 	return slices.Contains(ignores, filepath.Base(path))
	// }
	// return slices.Contains(ignores, pathExt)
	for _, pattern := range ignores {
		matched, err := regexp.MatchString(pattern, path)
		if err == nil && matched {
			return true
		}
	}
	return false
}

func isInvalidFile(path string) bool {
	base := fp.Base(path)
	ext := fp.Ext(path)
	if ext == "" && base != "Makefile" && base != "Dockerfile" && base != "LICENSE" {
		return true
	}
	return false
}
