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
	FileType  string
	FileColor string
	Steps     int
	Blanks    int
	Comments  int
	Files     int
	Bytes     string
	bytesBuf  int
}

type CntResult struct {
	Info        []FileInfo
	InputPath   string
	AllSteps    int
	AllBlanks   int
	AllComments int
	AllFiles    int
	AllBytes    string
}

const (
	maxCapacity          = 10 * 1024 * 1024 // 10MB
	concurrencyThreshold = 6
)

var FileTypeList map[string][]string = map[string][]string{
	".asm":          {"Assembly(.asm)", "Red"},
	".bat":          {"Batch File(.bat)", "Cyan"},
	".c":            {"C(.c)", "Yellow"},
	".cc":           {"C++(.cc)", "Yellow"},
	".cs":           {"C#(.h)", "Cyan"},
	".css":          {"CSS(.css)", "Yellow"},
	".cfg":          {"Configuration File(.cfg)", "HiBlack"},
	".cpp":          {"C++(.cpp)", "Yellow"},
	".clj":          {"Clojure(.clj)", "HiGreen"},
	".coffee":       {"CoffeeScript(.coffee)", "Yellow"},
	".d":            {"D(.d)", "Red"},
	".dart":         {"Dart(.dart)", "Cyan"},
	".dockerfile":   {"Dockerfile(.dockerfile)", "HiBlue"},
	".Dockerfile":   {"Dockerfile(.Dockerfile)", "HiBlue"},
	"Dockerfile":    {"Dockerfile", "HiBlue"},
	".dockerignore": {"Docker ignore file(.dockerignore)", "HiBlue"},
	".fs":           {"F#(.fs)", "Cyan"},
	".f90":          {"Fortran(.f90)", "HiWhite"},
	".erl":          {"Erlang(.erl)", "HiWhite"},
	".ex":           {"Elixir(.ex)", "Magenta"},
	".exs":          {"Elixir Script(.exs)", "Magenta"},
	".go":           {"Go(.go)", "Blue"},
	".groovy":       {"Groovy(.groovy)", "Yellow"},
	".gitignore":    {"Git ignore file(.gitignore)", "HiWhite"},
	".h":            {"C(.h)", "Yellow"},
	".hpp":          {"C++(.hpp)", "Yellow"},
	".html":         {"HTML(.html)", "HiYellow"},
	".ini":          {"Initialization File(.ini)", "HiWhite"},
	".js":           {"JavaScript(.js)", "Yellow"},
	".jsx":          {"JSX(.jsx)", "Yellow"},
	".java":         {"Java(.java)", "Red"},
	".json":         {"JSON File(.json)", "Yellow"},
	".jsonc":        {"JSONC File(.jsonc)", "Yellow"},
	".kt":           {"Kotlin(.kt)", "Green"},
	".lua":          {"Lua(.lua)", "Cyan"},
	".log":          {"Log File(.log)", "HiWhite"},
	".less":         {"Less(.less)", "Cyan"},
	"LICENSE":       {"LICENSE File", "Yellow"},
	".m":            {"Objective-C(.m)", "Yellow"},
	".ml":           {"ML(.ml)", "HiYellow"},
	".md":           {"Markdown(.md)", "Cyan"},
	".mk":           {"Makefile(.mk)", "HiRed"},
	".mod":          {"Go modules file(.mod)", "Blue"},
	"Makefile":      {"Makefile", "HiRed"},
	".nim":          {"Nim(.nim)", "YellowGreen"},
	".py":           {"Python(.py)", "HiCyan"},
	".pl":           {"Perl(.pl)", "Cyan"},
	".pas":          {"Pascal(.pas)", "HiWhite"},
	".php":          {"PHP(.php)", "Blue"},
	".r":            {"R(.r)", "Cyan"},
	".rb":           {"Ruby(.rb)", "Red"},
	".rs":           {"Rust(.rs)", "HiBlack"},
	".rtf":          {"Rich Text Format(.rtf)", "HiWhite"},
	".raku":         {"Raku(.raku)", "HiWhite"},
	".s":            {"Assembly(.s)", "Red"},
	".svg":          {"SVG", "Magenta"},
	".sh":           {"Shell Script(.sh)", "Green"},
	".sql":          {"SQL(.sql)", "Pink"},
	".sum":          {"Go sum file(.sum)", "Blue"},
	".sass":         {"SASS(.sass)", "Red"},
	".scss":         {"SCSS(.scss)", "Red"},
	".swift":        {"Swift(.swift)", "HiYellow"},
	".scala":        {"Scala(.scala)", "Red"},
	".ts":           {"TypeScript(.ts)", "Blue"},
	".tcl":          {"Tcl(.tcl)", "HiWhite"},
	".tsx":          {"TSX(.tsx)", "Blue"},
	".txt":          {"Plain Text(.txt)", "HiWhite"},
	".toml":         {"TOML(.toml)", "HiBlack"},
	".v":            {"Verilog(.v)", "HiWhite"},
	".vue":          {"Vue file(.vue)", "Green"},
	".vhdl":         {"VHDL(.vhdl)", "HiWhite"},
	".wasm":         {"WebAssembly", "Magenta"},
	".xml":          {"XML(.xml)", "Blue"},
	".xsl":          {"XSLT(.xsl)", "HiWhite"},
	".yml":          {"YAML File(.yml)", "Magenta"},
	".yaml":         {"YAML File(.yaml)", "Magenta"},
	"zsh":           {"ZSH", "Green"},
	".zig":          {"Zig(.zig)", "HiYellow"},
}

func Count(files []string, inputPath string) (CntResult, error) {
	var (
		result   CntResult
		bufMap   map[string]*FileInfo = make(map[string]*FileInfo)
		lenFiles uint                 = uint(len(files))
		mu       sync.Mutex
		wg       sync.WaitGroup
	)

	result.InputPath = inputPath

	if lenFiles >= concurrencyThreshold {
		var (
			aaa []string
			bbb []string
			ccc []string
		)

		alen := (lenFiles + 2) / 3
		blen := (lenFiles + 1) / 3
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
		existingMap.Steps += i.Steps
		existingMap.Blanks += i.Blanks
		existingMap.Comments += i.Comments
		existingMap.bytesBuf += i.bytesBuf
		existingMap.Files += 1
	} else {
		bufMap[i.FileType] = &i
		bufMap[i.FileType].Files += 1
		if _, typeFound := FileTypeList[i.FileType]; typeFound {
			bufMap[i.FileType].FileColor = FileTypeList[i.FileType][1]
		} else {
			bufMap[i.FileType].FileColor = "White"
		}
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
	"//":   {},
	"///":  {},
	"#":    {},
	"!":    {},
	"--":   {},
	"%":    {},
	";":    {},
	"#;":   {},
	"⍝":    {},
	"rem ": {},
	"::":   {},
	":  ":  {},
	"'":    {},
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
	"/*":       {},
	"/**":      {},
	"--":       {},
	"<!--":     {},
	"<%--":     {},
	"////":     {},
	"/+":       {},
	"/++":      {},
	"(*":       {},
	"{-":       {},
	"\"\"\"":   {},
	"'''":      {},
	"#=":       {},
	"--[[":     {},
	"%{":       {},
	"#[":       {},
	"=pod":     {},
	"=comment": {},
	"=begin":   {},
	"<#":       {},
	"#|":       {},
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
	"*/":     {},
	"**/":    {},
	"-->":    {},
	"--%>":   {},
	"--":     {},
	"+/":     {},
	"*)":     {},
	"-}":     {},
	"%}":     {},
	"=#":     {},
	"=cut":   {},
	"=end":   {},
	"--]]":   {},
	"]#":     {},
	"#>":     {},
	"\"\"\"": {},
	"'''":    {},
	"|#":     {},
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
