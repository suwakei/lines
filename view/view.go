package view

import (
	"fmt"

	"github.com/suwakei/lines/counter"
)

func Write(cntResult counter.CntResult, distPaths []string, ignoreListMap map[string][]string) {
	pathsLen := len(distPaths)
	if pathsLen == 0 {
		PrintTable(cntResult, ignoreListMap)
	} else {

	}
}

func largest(cntResult counter.CntResult, fileTypeList map[string][]string) (
	largestFileType int,
	largestLines int,
	largestBlanks int,
	largestComments int,
	largestFiles int,
	largestBytes int,
) {
	var fileTypeNum int
	largestFileType = len("FileType")
	largestLines = len("Lines")
	largestBlanks = len("Blanks")
	largestComments = len("Comments")
	largestFiles = len("Files")
	largestBytes = len("Bytes")
	info := cntResult.Info
	for _, i := range info {
		if _, found := fileTypeList[i.FileType]; found {
			fileTypeNum = len(fileTypeList[i.FileType][0])
		} else {
			fileTypeNum = len(fmt.Sprint(i.FileType))
		}

		linesNum := len(fmt.Sprint(i.Lines))
		blanksNum := len(fmt.Sprint(i.Blanks))
		commentsNum := len(fmt.Sprint(i.Comments))
		filesNum := len(fmt.Sprint(i.Files))
		bytesNum := len(fmt.Sprint(i.Bytes))

		if largestFileType < fileTypeNum {
			largestFileType = fileTypeNum
		}

		if largestLines < linesNum {
			largestLines = linesNum
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
