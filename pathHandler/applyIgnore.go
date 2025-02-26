package pathHandler

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func MakeIgnoreList[eOri string | []string] (ignores eOri) (map[string][]string, error) {
	ignoreListMap := make(map[string][]string, 2)

	switch v := any(ignores).(type) {
	case string:
		if v == "" {
			ignoreListMap["file"] = []string{}
			ignoreListMap["dir"] = []string{}
			return ignoreListMap, nil
		}

		if v != ".gitignore" {
			abs, _ := Parse(v)
			return nil, fmt.Errorf("[INFO]: ignore file must be .gitignore\n not exist %s", abs)
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
			if IsFile(line) {
				ignoreListMap["file"] = append(ignoreListMap["file"], line)
			} else {
				ignoreListMap["dir"] = append(ignoreListMap["dir"], line)
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
			if IsFile(i) {
				ignoreListMap["file"] = append(ignoreListMap["file"], ignore)
			} else {
				ignoreListMap["dir"] = append(ignoreListMap["dir"], ignore)
			}
		}
		return ignoreListMap, nil

	default:
		return nil, fmt.Errorf("[ERROR]: type of ignores is not valid")
	}
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}