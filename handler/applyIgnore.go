package handler

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	fp "path/filepath"
)

func MakeIgnoreList(ignoreFile string) ([]string, error) {
	if ignoreFile != ".gitignore" {
		abs, _ := fp.Abs(ignoreFile)
		return nil, fmt.Errorf("[INFO]: ignore file must be .gitignore\n not exist %s", abs)
	}

	p, err := fp.Abs(ignoreFile)
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
	var ignoreList []string
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		} else if scanner.Text() == "" {
			continue
		}
		ignoreList = append(ignoreList, scanner.Text())
	}
	return ignoreList, nil
}



