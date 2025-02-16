package counter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	fp "path/filepath"
)

type CntResult struct {
	Filetype string
	Steps int
	Blanks int
	Comments int
	Bytes int
}

func Count(files []string) ([]CntResult, error) {
	var results []CntResult
	var lenFiles uint = uint(len(files))
	if lenFiles >= 6 {
		var (
			aaa []string
			bbb []string
			ccc []string
			wg sync.WaitGroup
			mu sync.Mutex
		)

		alen := (lenFiles+2) / 3
		blen := (lenFiles+1) / 3
		clen := (lenFiles) / 3

		aaa = make([]string, 0, alen)
		bbb = make([]string, 0, blen)
		ccc = make([]string, 0, clen)

		aaa = append(aaa, files[0:alen-1]...)
		bbb = append(bbb, files[alen:alen+blen-1]...)
		ccc = append(ccc, files[alen+blen:lenFiles-1]...)

		wg.Add(3)
		processFiles := func(files []string) {
			defer wg.Done()
			for _, file := range files {
				r, err := count(file)
				if err != nil {
					fmt.Printf("[ERROR]: failed to countLine! %q\n %q\n", file, err)
					continue
				}
				mu.Lock()
				results = append(results, r)
				mu.Unlock()
			}
		}
		go processFiles(aaa)
		go processFiles(bbb)
		go processFiles(ccc)
		wg.Wait()
	} else {
		for _, file := range files {
			r, err := count(file)
			if err != nil {
				fmt.Printf("[ERROR]: failed to countLine! %q\n %q\n", file, err)
				return nil, err
			}
			results = append(results, r)
		}
	}
	return results, nil
}

func count(file string) (CntResult, error) {
	var result CntResult
	p, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		return CntResult{}, err
	}
	defer p.Close()
	result.Filetype = fp.Ext(file)


	scanner := bufio.NewScanner(p)
	const maxCapacity = 1024 * 1024
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var inBlockComment bool = false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		result.Steps++
		result.Bytes += len(line) + 1 // +1 for the newline character

		if line == "" {
			result.Blanks++
			continue
		}

		if isSingleComment(line) {
			result.Comments++
			continue
		}

		if isBeginBlockComments(line) {
			inBlockComment = true
			result.Comments++
			continue
		}

		if inBlockComment {
			result.Comments++
			if isEndBlockComments(line) {
				inBlockComment = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return CntResult{}, err
	}

	return result, nil
}


func isSingleComment(line string) bool {
	if len(line) == 0 {
		return false
	}
	return strings.HasPrefix(line, "//") ||
	strings.HasPrefix(line, "///") ||
	strings.HasPrefix(line, "#") ||
strings.HasPrefix(line, "!") ||
	strings.HasPrefix(line, "--") ||
	strings.HasPrefix(line, "%") ||
	strings.HasPrefix(line, ";") ||
	strings.HasPrefix(line, "#;") ||
	strings.HasPrefix(line, "⍝") ||
strings.HasPrefix(line, "rem ") ||
strings.HasPrefix(line, "::") ||
strings.HasPrefix(line, ":  ") ||
strings.HasPrefix(line, "'") ||
}

func isBeginBlockComments(line string) bool {
	if len(line) == 0 {
		return false
	}
	return strings.HasPrefix(line, "/*") ||
	strings.HasPrefix(line, "/**") ||
	strings.HasPrefix(line, "--") ||
	strings.HasPrefix(line, "<!--") ||
strings.HasPrefix(line, "<%--") ||
	strings.HasPrefix(line, "////") ||
	strings.HasPrefix(line, "/+") ||
	strings.HasPrefix(line, "/++") ||
	strings.HasPrefix(line, "(*") ||
	strings.HasPrefix(line, "{-") ||
	strings.HasPrefix(line, "\"\"\"") ||
strings.HasPrefix(line, "'''") ||
	strings.HasPrefix(line, "#=") ||
	strings.HasPrefix(line, "--[[") ||
	strings.HasPrefix(line, "%{") ||
	strings.HasPrefix(line, "#[") ||
	strings.HasPrefix(line, "=pod") ||
strings.HasPrefix(line, "=comment") ||
strings.HasPrefix(line, "=begin") ||
	strings.HasPrefix(line, "<#") ||
	strings.HasPrefix(line, "#|")
}

func isEndBlockComments(line string) bool {
	return strings.HasSuffix(line, "*/") ||
	strings.HasSuffix(line, "**/") ||
	strings.HasSuffix(line, "-->") ||
strings.HasSuffix(line, "--%>") ||
	strings.HasSuffix(line, "--") ||
	strings.HasSuffix(line, "+/") ||
	strings.HasSuffix(line, "*)") ||
	strings.HasSuffix(line, "-}") ||
	strings.HasSuffix(line, "%}") ||
	strings.HasSuffix(line, "=#") ||
	strings.HasSuffix(line, "=cut") ||
strings.HasPrefix(line, "=end") ||
	strings.HasSuffix(line, "--]]") ||
	strings.HasSuffix(line, "]#") ||
	strings.HasSuffix(line, "*)") ||
	strings.HasSuffix(line, "#>") ||
	strings.HasSuffix(line, "\"\"\"") ||
strings.HasPrefix(line, "'''") ||
	strings.HasSuffix(line, "|#")
}
