package counter

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	fp "path/filepath"
	"strings"
	"sync"
)

type FileInfo struct {
	FileType  string
	FileColor string
	Lines     int
	Blanks    int
	Comments  int
	Files     int
	Bytes     string
	bytesBuf  int
}

type CntResult struct {
	Info          []FileInfo
	InputPath     string
	TotalLines    int
	TotalBlanks   int
	TotalComments int
	TotalFiles    int
	TotalBytes    string
}

const (
	maxCapacity          = 10 * 1024 * 1024 // 10MB
	concurrencyThreshold = 6
)

var FileTypeList map[string][]string = map[string][]string{
	".asm":          {"Assembly(.asm)", "Red"},
	".bat":          {"Batch File(.bat)", "Cyan"},
	".bash":         {"BASH(.bash)", "HiGreen"},
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
	".md":           {"Markdown(.md)", "HiBlue"},
	".mk":           {"Makefile(.mk)", "HiRed"},
	".mod":          {"Go modules file(.mod)", "Blue"},
	"Makefile":      {"Makefile", "HiRed"},
	".nim":          {"Nim(.nim)", "YellowGreen"},
	".py":           {"Python(.py)", "HiBlue"},
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
	".zsh":          {"ZSH", "Green"},
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
		m.Bytes = fmt.Sprintf("%d(%s)", m.bytesBuf, readableBytes(uint64(m.bytesBuf)))
		result.Info = append(result.Info, *m)
	}
	result.assignTotals()
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
		if err := scanner.Err(); err != nil {
			log.Fatalf("scan failed %q", err)
		}

		info.Lines++
		line := strings.TrimSpace(scanner.Text())
		info.bytesBuf += len(line) + 1 // +1 for the newline character

		if line == "" {
			info.Blanks++
			continue
		}

		if info.isBeginBlockComments(line) {
			inBlockComment = true
			info.Comments++
			continue
		}

		if inBlockComment {
			info.Comments++
			if info.isEndBlockComments(line) {
				inBlockComment = false
			}
			continue
		}

		if info.isSingleComment(line) {
			info.Comments++
			continue
		}
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
		existingMap.Lines += i.Lines
		existingMap.Blanks += i.Blanks
		existingMap.Comments += i.Comments
		existingMap.Files += 1
		existingMap.bytesBuf += i.bytesBuf
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
	ext := fp.Ext(file)
	if ext == "" {
		base := fp.Base(file)
		return base
	} else {
		return ext
	}
}

func (r *CntResult) assignTotals() {
	var TotalBytesBuf int
	for _, i := range r.Info {
		r.TotalLines += i.Lines
		r.TotalBlanks += i.Blanks
		r.TotalComments += i.Comments
		r.TotalFiles += i.Files
		TotalBytesBuf += i.bytesBuf
	}
	r.TotalBytes = fmt.Sprintf("%d(%s)", TotalBytesBuf, readableBytes(uint64(TotalBytesBuf)))
}

// Efficiency of searching comment prefixes from O(n) to O(1)
var singleCommentPrefixes map[string][]string = map[string][]string{
	"//": {
		".c",
		".cc",
		".cs",
		".cpp",
		".go",
		".h",
		".hpp",
		".d",
		".dart",
		".fs",
		".go",
		".groovy",
		".js",
		".ts",
		".jsx",
		".tsx",
		".java",
		".jsonc",
		".kt",
		".m",
		".php",
		".rs",
	},
	"///": {".d", ".dart", ".rs"},
	"//!": {".rs"},
	"/**": {".dart"},
	"#": {
		".bash",
		".cfg",
		".coffee",
		".dockerfile",
		".Dockerfile",
		"Dockerfile",
		".dockerignore",
		".ex",
		".gitignore",
		".mk",
		"Makefile",
		".py",
		".pl",
		".php",
		".rb",
	},
	"##":  {".nim"},
	"!":   {".f90"},
	"--":  {".lua"},
	"%":   {".erl"},
	";":   {".asm", ".clj", ".ini"},
	"rem": {".bat"},
}

func (fi FileInfo) isSingleComment(line string) bool {
	lineLen := len(strings.TrimSpace(line))
	if lineLen == 0 {
		return false
	}

	for prefix, extensions := range singleCommentPrefixes {
		for _, ext := range extensions {
			if ext == fi.FileType {
				prefLen := len(prefix)
				if lineLen >= prefLen && string(line[:prefLen]) == prefix {
					return true
				}
			}
		}
	}

	return false
}

var blockCommentPrefixes map[string][]string = map[string][]string{
	"/*": {
		".c",
		".cc",
		".cs",
		".cpp",
		".css",
		".h",
		".hpp",
		".d",
		".dart",
		".go",
		".groovy",
		".js",
		".ts",
		".jsx",
		".tsx",
		".java",
		".jsonc",
		".kt",
		".m",
		".php",
		".rs",
	},
	"/*!":    {".rs"},
	"/**":    {".d", ".kt", ".m", ".rs"},
	"{":      {".pas"},
	"<!--":   {".html", ".xml", ".md"},
	"/++":    {".d"},
	"\"\"\"": {".ex", ".py"},
	"--[[":   {".lua"},
	"=pod":   {".pl"},
	"=begin": {".rb"},
	"###":    {".coffee"},
	"(#":     {".fs"},
}

func (fi FileInfo) isBeginBlockComments(line string) bool {
	lineLen := len(strings.TrimSpace(line))
	if lineLen == 0 {
		return false
	}

	for blockPrefix, extensions := range blockCommentPrefixes {
		for _, ext := range extensions {
			if ext == fi.FileType {
				prefLen := len(blockPrefix)
				if lineLen >= prefLen && string(line[:prefLen]) == blockPrefix {
					return true
				}
			}
		}
	}

	return false
}

var blockCommentSuffixes map[string][]string = map[string][]string{
	"*/": {
		".c",
		".cc",
		".cs",
		".cpp",
		".css",
		".h",
		".hpp",
		".d",
		".dart",
		".go",
		".groovy",
		".js",
		".ts",
		".jsx",
		".tsx",
		".java",
		".jsonc",
		".kt",
		".m",
		".php",
		".rs",
	},
	"+/":     {".d"},
	"}":      {".pas"},
	"-->":    {".html", ".xml", ".md"},
	"=cut":   {".pl"},
	"=end":   {".rb"},
	"]]":     {".lua"},
	"\"\"\"": {".ex", ".py"},
	"###":    {".coffee"},
	"#)":     {".fs"},
}

func (fi FileInfo) isEndBlockComments(line string) bool {
	lineLen := len(strings.TrimSpace(line))
	if lineLen == 0 {
		return false
	}

	for blockSuffix, extensions := range blockCommentSuffixes {
		for _, ext := range extensions {
			if ext == fi.FileType {
				suffLen := len(blockSuffix)
				if lineLen >= suffLen && string(line[:suffLen]) == blockSuffix {
					return true
				}
			}
		}
	}

	return false
}

func readableBytes(s uint64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	var base float64 = 1000

	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}

	logn := func(n, b float64) float64 {
		return math.Log(n) / math.Log(b)
	}

	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %s"
	if val < 10 {
		f = "%.1f %s"
	}
	return fmt.Sprintf(f, val, suffix)
}
