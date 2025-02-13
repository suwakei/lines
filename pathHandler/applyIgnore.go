package pathHandler

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	fp "path/filepath"
)

func MakeIgnoreList[eOri string | []string] (ignores eOri) (map[string][]string, error) {
	var ignoreListMap map[string][]string = make(map[string][]string, 2)

	if ignoreFilePath, ok := any(ignores).(string); ok{
		if ignoreFilePath != ".gitignore" {
			abs, _ := Parse(ignoreFilePath)
			return nil, fmt.Errorf("[INFO]: ignore file must be .gitignore\n not exist %s", abs)
		}
		p, err := Parse(ignoreFilePath)
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(p); os.IsNotExist(err) {
			return nil, err
		}

		f, err := os.Open(p)
		if err != nil {
			return nil, err
		}

		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "#") {
				continue
			} else if scanner.Text() == "" {
				continue
			}
			if IsFile(scanner.Text()) {
				ignoreListMap["file"] = append(ignoreListMap["file"], scanner.Text())
			} else {
				ignoreListMap["dir"] = append(ignoreListMap["dir"], scanner.Text())
			}
		}
		return ignoreListMap, nil
	} else {
		ignoreSlice, ok := any(ignores).([]string)
		if !ok {
			return nil, fmt.Errorf("[ERROR]: type of ignores is not []string")
		}
		if len(ignoreSlice) == 0 {
			return nil, fmt.Errorf("[ERROR]: no extension specified")
		}
		for _, ignore := range ignoreSlice {
			i := fp.Base(ignore)
			if IsFile(ignore) {
				ignoreListMap["file"] = append(ignoreListMap["file"], i)
			} else {
				ignoreListMap["dir"] = append(ignoreListMap["dir"], i)
			}
		}
		return ignoreListMap, nil
	}
}

func IsFile(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        return false
    }
    return !info.IsDir()
}
