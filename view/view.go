package view

import (
	"fmt"

	"github.com/suwakei/steps/counter"
)

var fileTypeList map[string][]string = map[string][]string{
	".c": {"C", "yellow"},
	".h": {"C", "yellow"},
	".cpp": {"C++", "yellow"},
	".cc": {"C++", "yellow"},
	".hpp": {"C++", "yellow"},
	".cs": {"C#", "cyan"},
	".html": {"HTML", "orange"},
	".css": {"CSS", "yellow"},
	".scss": {"SCSS", "red"},
	".sass": {"SASS", "red"},
	".py": {"Python", "white"},
	".java": {"Java", "red"},
	".js": {"JavaScript", "yellow"},
	".ts": {"TypeScript", "blue"},
	".rb": {"Ruby", "red"},
	".rs": {"Rust", "gray"},
	".zig": {"Zig", "orange"},
	".go": {"Go", "blue"},
	".php": {"PHP", "blue"},
	".xml": {"XML", "blue"},
	".json": {"JSON File", "yellow"},
	".jsonc": {"JSONC File", "yellow"},
	".yaml": {"YAML File", "purple"},
	".yml": {"YAML File", "purple"},
	".toml": {"TOML", "gray"},
	".md": {"Markdown", "cyan"},
	".txt": {"Plain Text", "white"},
	".sql": {"SQL", "pink"},
	".sh": {"Shell Script", "green"},
	".bat": {"Batch File", "cyan"},
	".pl": {"Perl", "cyan"},
	".swift": {"Swift", "orange"},
	".r": {"R", "cyan"},
	".scala": {"Scala", "red"},
	".kt": {"Kotlin", "orange"},
	".dart": {"Dart", "cyan"},
	".asm": {"Assembly", "red"},
	".lua": {"Lua", "cyan"},
	".clj": {"Clojure", "yellowgreen"},
	".coffee": {"CoffeeScript", "yellow"},
	".f90": {"Fortran", "white"},
	".groovy": {"Groovy", "yellow"},
	".v": {"Verilog", "white"},
	".vhdl": {"VHDL", "white"},
	".d": {"D", "red"},
	".nim": {"Nim", "yellowgreen"},
	".pas": {"Pascal", "white"},
	".tcl": {"Tcl", "white"},
	".raku": {"Raku", "white"},
	".erl": {"Erlang", "white"},
	".ex": {"Elixir", "purple"},
	".exs": {"Elixir Script", "purple"},
	".fs": {"F#", "cyan"},
	".ml": {"ML", "orange"},
	".m": {"Objective-C", "yellow"},
	".s": {"Assembly", "red"},
	".xsl": {"XSLT", "white"},
	".less": {"Less", "cyan"},
	".log": {"Log File", "white"},
	".ini": {"Initialization File", "white"},
	".cfg": {"Configuration File", "gray"},
	".rtf": {"Rich Text Format", "white"},
	".doc": {"Microsoft Word Document", "cyan"},
	".docx": {"Microsoft Word Document (Open XML)", "cyan"},
	".pdf": {"Portable Document Format", "red"},
	".epub": {"Electronic Publication", "white"},
}

func Write(cntResult counter.CntResult, distPaths []string, ignoreListMap map[string][]string) {
	pathsLen := len(distPaths)
	if pathsLen == 0 {
		PrintTable(cntResult, ignoreListMap)
	}
}

func b2kb(Bytes int) int {
	return Bytes / 1024
}

func largest(cntResult counter.CntResult, fileTypeList map[string][]string) (
	largestFileType int,
	largestSteps int,
	largestBlanks int,
	largestComments int,
	largestFiles int,
	largestBytes int,
	) {
	var fileTypeNum int
	info := cntResult.Info
	for _, i := range info {
		if _, found := fileTypeList[i.Filetype]; found {
			fileTypeNum = len(fmt.Sprintf("%s(%s)", fileTypeList[i.Filetype][0], i.Filetype))
		} else {
			fileTypeNum = len(fmt.Sprint(i.Filetype))
		}
		stepsNum := len(fmt.Sprint(i.Steps))
		blanksNum := len(fmt.Sprint(i.Blanks))
		commentsNum := len(fmt.Sprint(i.Comments))
		filesNum := len(fmt.Sprint(i.Files))
		bytesNum := len(fmt.Sprint(i.Bytes))

		if largestFileType < fileTypeNum {
			largestFileType = fileTypeNum
		}

		if largestSteps < stepsNum {
			largestSteps = stepsNum
		}

		if largestBlanks < blanksNum {
			largestBlanks = blanksNum
		}

		if largestComments < commentsNum {
			largestComments = commentsNum
		}

		if largestFiles < filesNum {
			largestFiles = filesNum
		}

		if largestBytes < bytesNum {
			largestBytes = bytesNum
		}
	}
	return 
}