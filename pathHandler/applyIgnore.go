package pathHandler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MakeIgnoreList[eOri string | []string](ignoreFile eOri) (map[string][]string, error) {
	ignoreListMap := make(map[string][]string, 2)

	switch v := any(ignoreFile).(type) {
	case string:
		if v == "" {
			ignoreListMap["file"] = nil
			ignoreListMap["dir"] = nil
			return ignoreListMap, nil
		}

		p, err := Parse(v)
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
			line := scanner.Text()
			if strings.HasPrefix(line, "#") || line == "" {
				continue
			}
			if IsDir(line) {
				ignoreListMap["dir"] = append(ignoreListMap["dir"], line)
			} else {
				ignoreListMap["file"] = append(ignoreListMap["file"], line)
			}
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return ignoreListMap, nil

	case []string:
		if len(v) == 0 {
			return nil, fmt.Errorf("[ERROR]: no extension specified")
		}
		for _, ignore := range v {
			i, err := Parse(ignore)
			if err != nil {
				return nil, err
			}
			if IsDir(i) {
				ignoreListMap["dir"] = append(ignoreListMap["dir"], ignore)
			} else {
				ignoreListMap["file"] = append(ignoreListMap["file"], ignore)
			}
		}
		return ignoreListMap, nil

	default:
		return nil, fmt.Errorf("[ERROR]: type of ignores is not valid")
	}
}


func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
