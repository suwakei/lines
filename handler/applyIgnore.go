package handler

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func MakeIgnoreList(ignoreFile string) ([]string, error) {
	if ignoreFile != ".gitignore" {
		abs, _ := Parse(ignoreFile)
		return nil, fmt.Errorf("[INFO]: ignore file must be .gitignore\n not exist %s", abs)
	}
	p, err := Parse(ignoreFile)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		return nil, err
	}

	var ignoreList []string
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
}