package handler

import (
	"io"
	"bufio"
	"strings"
)

func MakeIgnoreList(ignoreFile io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(ignoreFile)
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



