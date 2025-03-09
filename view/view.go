package view

import (
	"fmt"

	"github.com/suwakei/steps/counter"
)

var fileTypeList map[string][]string = map[string][]string{
	".c": {"C(.c)", "yellow"},
	".h": {"C(.h)", "yellow"},
	".cpp": {"C++(.cpp)", "yellow"},
	".cc": {"C++(.cc)", "yellow"},
	".hpp": {"C++(.hpp)", "yellow"},
	".cs": {"C#(.h)", "cyan"},
	".html": {"HTML(.html)", "orange"},
	".css": {"CSS(.css)", "yellow"},
	".scss": {"SCSS(.scss)", "red"},
	".sass": {"SASS(.sass)", "red"},
	".py": {"Python(.py)", "white"},
	".java": {"Java(.java)", "red"},
	".js": {"JavaScript(.js)", "yellow"},
	".jsx": {"JSX(.jsx)", "yellow"},
	".ts": {"TypeScript(.ts)", "blue"},
	".tsx": {"TSX(.tsx)", "blue"},
	".rb": {"Ruby(.rb)", "red"},
	".rs": {"Rust(.rs)", "gray"},
	".zig": {"Zig(.zig)", "orange"},
	".go": {"Go(.go)", "blue"},
	".php": {"PHP(.php)", "blue"},
	".xml": {"XML(.xml)", "blue"},
	".json": {"JSON File(.json)", "yellow"},
	".jsonc": {"JSONC File(.jsonc)", "yellow"},
	".yaml": {"YAML File(.yaml)", "purple"},
	".yml": {"YAML File(.yml)", "purple"},
	".toml": {"TOML(.toml)", "gray"},
	".md": {"Markdown(.md)", "cyan"},
	".txt": {"Plain Text(.txt)", "white"},
	".sql": {"SQL(.sql)", "pink"},
	".sh": {"Shell Script(.sh)", "green"},
	".bat": {"Batch File(.bat)", "cyan"},
	".pl": {"Perl(.pl)", "cyan"},
	".swift": {"Swift(.swift)", "orange"},
	".r": {"R(.r)", "cyan"},
	".scala": {"Scala(.scala)", "red"},
	".kt": {"Kotlin(.kt)", "orange"},
	".dart": {"Dart(.dart)", "cyan"},
	".asm": {"Assembly(.asm)", "red"},
	".lua": {"Lua(.lua)", "cyan"},
	".clj": {"Clojure(.clj)", "yellowgreen"},
	".coffee": {"CoffeeScript(.coffee)", "yellow"},
	".f90": {"Fortran(.f90)", "white"},
	".groovy": {"Groovy(.groovy)", "yellow"},
	".v": {"Verilog(.v)", "white"},
	".vhdl": {"VHDL(.vhdl)", "white"},
	".d": {"D(.d)", "red"},
	".nim": {"Nim(.nim)", "yellowgreen"},
	".pas": {"Pascal(.pas)", "white"},
	".tcl": {"Tcl(.tcl)", "white"},
	".raku": {"Raku(.raku)", "white"},
	".erl": {"Erlang(.erl)", "white"},
	".ex": {"Elixir(.ex)", "purple"},
	".exs": {"Elixir Script(.exs)", "purple"},
	".fs": {"F#(.fs)", "cyan"},
	".ml": {"ML(.ml)", "orange"},
	".m": {"Objective-C(.m)", "yellow"},
	".s": {"Assembly(.s)", "red"},
	".xsl": {"XSLT(.xsl)", "white"},
	".less": {"Less(.less)", "cyan"},
	".log": {"Log File(.log)", "white"},
	".ini": {"Initialization File(.ini)", "white"},
	".cfg": {"Configuration File(.cfg)", "gray"},
	".rtf": {"Rich Text Format(.rtf)", "white"},
	".doc": {"Microsoft Word Document(.doc)", "cyan"},
	".docx": {"Microsoft Word Document (Open XML)(.docx)", "cyan"},
	".pdf": {"Portable Document Format(.pdf)", "red"},
	".epub": {"Electronic Publication(.epub)", "white"},
}

func Write(cntResult counter.CntResult, distPaths []string, ignoreListMap map[string][]string) {
	pathsLen := len(distPaths)
	if pathsLen == 0 {
		PrintTable(cntResult, ignoreListMap)
	}
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
	largestFileType = len("FileType")
	largestSteps = len("Steps")
	largestBlanks = len("Blanks")
	largestComments = len("Comments")
	largestFiles = len("Files")
	largestBytes = len("Bytes")
	info := cntResult.Info
	for _, i := range info {
		if _, found := fileTypeList[i.Filetype]; found {
			fileTypeNum = len(fileTypeList[i.Filetype][0])
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