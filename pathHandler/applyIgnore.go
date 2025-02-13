package pathHandler

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func MakeIgnoreList[eOri string | []string] (ignores eOri) ([]string, error) {
	var ignoreList []string
	ignoreList = append(ignoreList, "*.exe", "*.com", "*.dll", "*.so", "*.dylib", "*.xls", "*.xlsx", "*.pdf", "*.doc", "*.docx", "*.ppt", "*.pptx")
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
			ignoreList = append(ignoreList, scanner.Text())
		}
		return ignoreList, nil
	} else {
		ignoreSlice, ok := any(ignores).([]string)
		if !ok {
			return nil, fmt.Errorf("[ERROR]: type of ignores is not []string")
		}
		if len(ignoreSlice) == 0 {
			return nil, fmt.Errorf("[ERROR]: no extension specified")
		}
		ignoreList = append(ignoreList, ignoreSlice...)
		return ignoreList, nil
		}
	}