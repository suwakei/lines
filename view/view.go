package view

import (
	"fmt"

	"github.com/suwakei/steps/counter"
)

var fileTypeList map[string]string = map[string]string{
	".c":  "C",
	".h":  "C",
	".cpp": "C++",
	".cc":  "C++",
	".hpp": "C++",
	".cs": "C#",
	".html": "HTML",
	".css": "CSS",
	".scss": "SCSS",
	".sass": "SASS",
	".py": "Python",
	".java": "Java",
	".js": "JavaScript",
	".ts": "TypeScript",
	".rb": "Ruby",
	".go": "Go",
	".php": "PHP",
	".xml": "XML",
	".json": "JSON",
	".yaml": "YAML",
	".yml": "YAML",
	".md": "Markdown",
	".txt": "Plain Text",
	".sql": "SQL",
	".sh": "Shell Script",
	".bat": "Batch File",
	".pl": "Perl",
	".swift": "Swift",
	".r": "R",
	".scala": "Scala",
	".kt": "Kotlin",
	".dart": "Dart",
	".asm": "Assembly",
	".lua": "Lua",
	".clj": "Clojure",
	".coffee": "CoffeeScript",
	".f90": "Fortran",
	".groovy": "Groovy",
	".v": "Verilog",
	".vhdl": "VHDL",
	".d": "D",
	".nim": "Nim",
	".pas": "Pascal",
	".tcl": "Tcl",
	".raku": "Raku",
	".erl": "Erlang",
	".ex": "Elixir",
	".exs": "Elixir Script",
	".fs": "F#",
	".ml": "ML",
	".m": "Objective-C",
	".s": "Assembly",
	".xsl": "XSLT",
	".less": "Less",
	".log": "Log File",
	".ini": "Initialization File",
	".cfg": "Configuration File",
	".rtf": "Rich Text Format",
	".doc": "Microsoft Word Document",
	".docx": "Microsoft Word Document (Open XML)",
	".pdf": "Portable Document Format",
	".epub": "Electronic Publication",
}

func Write(cntResult counter.CntResult, distPaths []string, ignoreListMap map[string][]string) {
	pathsLen := len(distPaths)
	if pathsLen == 0 {
		PrintTable(cntResult, ignoreListMap)
	} else {

	}
}

func b2kb(Bytes int) int {
	var kb int
	if Bytes >= 1024 {
		kb = Bytes / 1024
		return kb
	} else {
		return Bytes
	}
}

func largest(cntResult counter.CntResult) (
	largestFileType int,
	largestSteps int,
	largestBlanks int,
	largestComments int,
	largestFiles int,
	largestBytes int,
	) {
	info := cntResult.Info
	for _, i := range info {
		fileTypeNum := len(fmt.Sprint(i.Filetype))
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