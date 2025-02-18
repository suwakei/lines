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

const concurrencyThreshold = 6

func Count(files []string) ([]CntResult, error) {
	var results []CntResult
	var lenFiles uint = uint(len(files))
	if lenFiles >= concurrencyThreshold {
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

		aaa = append(aaa, files[0:alen]...)
		bbb = append(bbb, files[alen:alen+blen]...)
		ccc = append(ccc, files[alen+blen:lenFiles]...)

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
	
	result.Filetype = retFileType(file)

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

// Efficiency of searching comment prefixes from O(n) to O(1)
var singleCommentPrefixes map[string]struct{} = map[string]struct{}{
	"//": {},
	"///": {},
	"#": {},
	"!": {},
	"--": {},
	"%": {},
	";": {},
	"#;": {},
	"⍝": {},
	"rem ": {},
	"::": {},
	":  ": {},
	"'": {},
}


func isSingleComment(line string) bool {
	lineLen := len(line)
    if lineLen == 0 {
        return false
    }

	// conpare prefix to line
    for prefix := range singleCommentPrefixes {
		prefLen := len(prefix)
        if lineLen >= prefLen && line[:prefLen] == prefix {
            return true
        }
    }

    return false
}


var blockCommentPrefixes map[string]struct{} = map[string]struct{}{
	"/*": {},
	"/**": {},
	"--": {},
	"<!--": {},
	"<%--": {},
	"////": {},
	"/+": {},
	"/++": {},
	"(*": {},
	"{-": {},
	"\"\"\"": {},
	"'''": {},
	"#=": {},
	"--[[": {},
	"%{": {},
	"#[": {},
	"=pod": {},
	"=comment": {},
	"=begin": {},
	"<#": {},
	"#|": {},
}

func isBeginBlockComments(line string) bool {
	lineLen := len(line)

    if lineLen == 0 {
        return false
    }

    for prefix := range blockCommentPrefixes {
		prefLen := len(prefix)
        if lineLen >= prefLen && line[:prefLen] == prefix {
            return true
        }
    }

    return false
}


var blockCommentSuffixes map[string]struct{} = map[string]struct{}{
	"*/": {}, 
	"**/": {},
	"-->": {},
	"--%>": {},
	"--": {},
	"+/": {},
	"*)": {},
	"-}": {},
	"%}": {},
	"=#": {},
	"=cut": {},
	"=end": {},
	"--]]": {},
	"]#": {},
	"#>": {},
	"\"\"\"": {},
	"'''": {},
	"|#": {},
}

func isEndBlockComments(line string) bool {
    if len(line) == 0 {
        return false
    }

	// confirm suffix directly
    _, exists := blockCommentSuffixes[line[len(line)-2:]] // confirm last two chars
    if exists {
        return exists
    }

	_, exists = blockCommentSuffixes[line[len(line)-3:]] // confirm last three chars
    return exists
}

func retFileType(file string) string {
	if fp.Ext(file) == "" {
		b := fp.Base(file)
		return b
	} else {
		return fp.Ext(file)
	}

}