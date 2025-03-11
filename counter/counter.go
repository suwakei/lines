package counter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	fp "path/filepath"
	"strings"
	"sync"
)

type FileInfo struct {
	FileType string
	FileColor string
	Steps int
	Blanks int
	Comments int
	Files int
	Bytes string
	bytesBuf int
}

type CntResult struct {
    Info []FileInfo
	InputPath string
    AllSteps int
    AllBlanks int
    AllComments int
	AllFiles int
    AllBytes string
}

const (
	maxCapacity = 64 * 1024
	concurrencyThreshold = 6
)

var FileTypeList map[string][]string = map[string][]string{
	".c": {"C(.c)", "Yellow"},
	".h": {"C(.h)", "Yellow"},
	".cpp": {"C++(.cpp)", "Yellow"},
	".cc": {"C++(.cc)", "Yellow"},
	".hpp": {"C++(.hpp)", "Yellow"},
	".cs": {"C#(.h)", "Cyan"},
	".html": {"HTML(.html)", "orange"},
	".css": {"CSS(.css)", "Yellow"},
	".scss": {"SCSS(.scss)", "Red"},
	".sass": {"SASS(.sass)", "Red"},
	".py": {"Python(.py)", "HiCyan"},
	".java": {"Java(.java)", "Red"},
	".js": {"JavaScript(.js)", "Yellow"},
	".jsx": {"JSX(.jsx)", "Yellow"},
	".ts": {"TypeScript(.ts)", "Blue"},
	".tsx": {"TSX(.tsx)", "Blue"},
	".rb": {"Ruby(.rb)", "Red"},
	".rs": {"Rust(.rs)", "Gray"},
	".zig": {"Zig(.zig)", "orange"},
	".go": {"Go(.go)", "Blue"},
	".php": {"PHP(.php)", "Blue"},
	".xml": {"XML(.xml)", "Blue"},
	".json": {"JSON File(.json)", "Yellow"},
	".jsonc": {"JSONC File(.jsonc)", "Yellow"},
	".yaml": {"YAML File(.yaml)", "Magenta"},
	".yml": {"YAML File(.yml)", "Magenta"},
	".toml": {"TOML(.toml)", "Gray"},
	".md": {"Markdown(.md)", "Cyan"},
	".txt": {"Plain Text(.txt)", "White"},
	".sql": {"SQL(.sql)", "pink"},
	".sh": {"Shell Script(.sh)", "Green"},
	".bat": {"Batch File(.bat)", "Cyan"},
	".pl": {"Perl(.pl)", "Cyan"},
	".swift": {"Swift(.swift)", "orange"},
	".r": {"R(.r)", "Cyan"},
	".scala": {"Scala(.scala)", "Red"},
	".kt": {"Kotlin(.kt)", "orange"},
	".dart": {"Dart(.dart)", "Cyan"},
	".asm": {"Assembly(.asm)", "Red"},
	".lua": {"Lua(.lua)", "Cyan"},
	".clj": {"Clojure(.clj)", "HiGreen"},
	".coffee": {"CoffeeScript(.coffee)", "Yellow"},
	".f90": {"Fortran(.f90)", "White"},
	".groovy": {"Groovy(.groovy)", "Yellow"},
	".v": {"Verilog(.v)", "White"},
	".vhdl": {"VHDL(.vhdl)", "White"},
	".d": {"D(.d)", "Red"},
	".nim": {"Nim(.nim)", "YellowGreen"},
	".pas": {"Pascal(.pas)", "White"},
	".tcl": {"Tcl(.tcl)", "White"},
	".raku": {"Raku(.raku)", "White"},
	".erl": {"Erlang(.erl)", "White"},
	".ex": {"Elixir(.ex)", "Magenta"},
	".exs": {"Elixir Script(.exs)", "Magenta"},
	".fs": {"F#(.fs)", "Cyan"},
	".ml": {"ML(.ml)", "orange"},
	".m": {"Objective-C(.m)", "Yellow"},
	".s": {"Assembly(.s)", "Red"},
	".xsl": {"XSLT(.xsl)", "White"},
	".less": {"Less(.less)", "Cyan"},
	".log": {"Log File(.log)", "White"},
	".ini": {"Initialization File(.ini)", "White"},
	".cfg": {"Configuration File(.cfg)", "Gray"},
	".rtf": {"Rich Text Format(.rtf)", "White"},
	".doc": {"Microsoft Word Document(.doc)", "Cyan"},
	".docx": {"Microsoft Word Document (Open XML)(.docx)", "Cyan"},
	".pdf": {"Portable Document Format(.pdf)", "Red"},
	".epub": {"Electronic Publication(.epub)", "White"},
}

func Count(files []string, inputPath string) (CntResult, error) {
	var(
		result CntResult
		bufMap map[string]*FileInfo = make(map[string]*FileInfo)
		lenFiles uint = uint(len(files))
		mu sync.Mutex
		wg sync.WaitGroup
	)

	result.InputPath = inputPath

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
		m.Bytes = fmt.Sprintf("%d(%dKB)", m.bytesBuf, b2kb(m.bytesBuf))
		result.Info = append(result.Info, *m)
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
	
	info.FileType = retFileType(p.Name())

	scanner := bufio.NewScanner(p)
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanLines)

	var inBlockComment bool = false
	for scanner.Scan() {
		info.Steps++
		line := strings.TrimSpace(scanner.Text())
		info.bytesBuf += len(line) + 1 // +1 for the newline character

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
	if existingMap, found := bufMap[i.FileType]; found {
		if existingMap.FileColor == "" {
			existingMap.FileColor = FileTypeList[existingMap.FileType][1]
		}
		existingMap.Steps += i.Steps
		existingMap.Blanks += i.Blanks
		existingMap.Comments += i.Comments
		existingMap.bytesBuf += i.bytesBuf
		existingMap.Files += 1
	} else {
		bufMap[i.FileType] = &i
		bufMap[i.FileType].Files += 1
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
	var AllBytesBuf int
	for _, i := range r.Info {
		r.AllSteps += i.Steps
		r.AllBlanks += i.Blanks
		r.AllComments += i.Comments
        r.AllFiles += i.Files
		AllBytesBuf += i.bytesBuf
	}
	r.AllBytes = fmt.Sprintf("%d(%dKB)", AllBytesBuf, b2kb(AllBytesBuf))
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
    if len(line) >= 2 {
        if _, exists := blockCommentSuffixes[line[len(line)-2:]]; exists {
            return true
        }
    }

    if len(line) >= 3 {
        if _, exists := blockCommentSuffixes[line[len(line)-3:]]; exists {
            return true
        }
    }

	return false
}

func b2kb(Bytes int) int {
	return Bytes / 1024
}