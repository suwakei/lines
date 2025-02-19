package counter

import (
	"bufio"
	"log"
	"os"
	fp "path/filepath"
	"strings"
	"sync"
)

type FileInfo struct {
	Filetype string
	Steps int
	Blanks int
	Comments int
	Bytes int
}

type CntResult struct {
    info []FileInfo
    AllSteps int
    AllBlanks int
    AllComments int
    AllBytes int64
}

const (
	maxCapacity = 1024 * 1024
	concurrencyThreshold = 6
)

func Count(files []string) (CntResult, error) {
	var(
		result CntResult
		bufMap map[string]*FileInfo = make(map[string]*FileInfo)
		lenFiles uint = uint(len(files))
		mu sync.Mutex
		wg sync.WaitGroup
	)

	if lenFiles >= concurrencyThreshold {
		var (
			aaa []string
			bbb []string
			ccc []string
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
		go processFiles(aaa, &wg, &mu, bufMap)
		go processFiles(bbb, &wg, &mu, bufMap)
		go processFiles(ccc, &wg, &mu, bufMap)
		wg.Wait()
	} else {
		for _, file := range files {
			if err := processFile(file, bufMap, &mu); err != nil {
				log.Fatalf("processFile failed %q", err)
			}
		}
	}

	for _, m := range bufMap {
		result.info = append(result.info, *m)
	}
	result.assignAlls()
	return result, nil
}

func count(file string) (FileInfo, error) {
	var info FileInfo
	p, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("openFile failed %q", err)
	}
	defer p.Close()
	
	info.Filetype = retFileType(p.Name())

	scanner := bufio.NewScanner(p)
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var inBlockComment bool = false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		info.Steps++
		info.Bytes += len(line) + 1 // +1 for the newline character

		if line == "" {
			info.Blanks++
			continue
		}

		if isSingleComment(line) {
			info.Comments++
			continue
		}

		if isBeginBlockComments(line) {
			inBlockComment = true
			info.Comments++
			continue
		}

		if inBlockComment {
			info.Comments++
			if isEndBlockComments(line) {
				inBlockComment = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scan failed %q", err)
	}
	return info, nil
}

func processFiles(files []string, wg *sync.WaitGroup, mu *sync.Mutex, bufMap map[string]*FileInfo) {
	defer wg.Done()
	for _, file := range files {
		if err := processFile(file, bufMap, mu); err != nil {
			log.Fatalf("[ERROR]: failed to countLine! %q\n %q\n", file, err)
		}
	}
}

func processFile(file string, bufMap map[string]*FileInfo, mu *sync.Mutex) error {
	i, err := count(file)
	if err != nil {
		log.Fatalf("[ERROR]: count function failed %q", err)
	}
	mu.Lock()
	defer mu.Unlock()
	if existingMap, found := bufMap[i.Filetype]; found {
		existingMap.Steps += i.Steps
		existingMap.Blanks += i.Blanks
		existingMap.Comments += i.Comments
		existingMap.Bytes += i.Bytes
	} else {
		bufMap[i.Filetype] = &i
	}
	return nil
}

func retFileType(file string) string {
	if fp.Ext(file) == "" {
		b := fp.Base(file)
		return b
	} else {
		return fp.Ext(file)
	}
}

func (r *CntResult) assignAlls() {
	for _, i := range r.info {
		r.AllSteps += i.Steps
		r.AllBlanks += i.Blanks
		r.AllComments += i.Comments
		r.AllBytes += int64(i.Bytes)
	}
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