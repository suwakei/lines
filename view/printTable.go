package view

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/suwakei/lines/counter"
)

type largests struct {
	largestFileType int
	largestLines    int
	largestBlanks   int
	largestComments int
	largestFiles    int
	largestBytes    int
}

func RetTotals(cntResult counter.CntResult, ignoreListMap map[string][]string) (
	[]string,
	map[string][]string,
	) {
	var (
		AllIgnoreFiles []string
		AllIgnoreDirs []string
	)
	if len(ignoreListMap["file"]) != 0 {
		AllIgnoreFiles = ignoreListMap["file"]
	} else {
		AllIgnoreFiles = nil
	}

	if len(ignoreListMap["dir"]) != 0 {
		AllIgnoreDirs = ignoreListMap["dir"]
	} else {
		AllIgnoreDirs = nil
	}

	totalMap := make([]string, 0, 12)
	ignoreMap := make(map[string][]string, 2)
	totalMap = []string{
		"Target Abs Path: ", cntResult.InputPath,
		"Total Lines: ", fmt.Sprint(cntResult.TotalLines),
		"Total Blanks: ", fmt.Sprint(cntResult.TotalBlanks),
		"Total Comments: ", fmt.Sprint(cntResult.TotalComments),
		"Total Files: ", fmt.Sprint(cntResult.TotalFiles),
		"Total Bytes: ", fmt.Sprint(cntResult.TotalBytes),
	}

	ignoreMap = map[string][]string {
		"All ignore files: ": AllIgnoreFiles,
		"All ignore dirs: ": AllIgnoreDirs,
	}

	return totalMap, ignoreMap
}

func PrintTable(cntResult counter.CntResult, ignoreListMap map[string][]string) {
	totals, allIgnores := RetTotals(cntResult, ignoreListMap)

	for i := 0; i < len(totals); {
		fmt.Println(totals[i], totals[i+1])
		i += 2
	}

	for k, v := range allIgnores {
		if v == nil {
			fmt.Println(k, "none")
		} else {
			fmt.Println(k, v)
		}
	}

	fmt.Print("\n\n")

	largests := largestsNew(cntResult)

	header, body, footer := makeTable(cntResult, *largests)

	fmt.Println(header)
	fmt.Println(body)
	fmt.Println(footer)
}

func space(currentTarget string, largestLen int) string {
	currentLen := len(currentTarget)
	var spaceNum int
	if largestLen-currentLen > -1 {
		spaceNum = largestLen - currentLen
	} else {
		spaceNum = 0
	}
	return strings.Repeat(" ", spaceNum)
}

func makeTable(cntResult counter.CntResult, largests largests) (
	header string,
	body string,
	footer string,
) {
	var lineLen int
	numLen := getNumLen(cntResult)
	header, lineLen = makeHeader(largests, numLen)

	body = makeBody(cntResult, largests)
	footer = " " + strings.Repeat("¯", lineLen-2)
	return header, body, footer
}

func makeHeader(largests largests, numLen int) (string, int) {
	var header strings.Builder

	header.WriteString(fmt.Sprintf("|%s|  %s%s  |  %s%s  |  %s%s  |  %s%s  |  %s%s  |  %s%s  |",
		strings.Repeat("#", numLen),
		"FileType",
		space("FileType", largests.largestFileType),
		"Lines",
		space("Lines", largests.largestLines),
		"Blanks",
		space("Blanks", largests.largestBlanks),
		"Comments",
		space("Comments", largests.largestComments),
		"Files",
		space("Files", largests.largestFiles),
		"Bytes",
		space("Bytes", largests.largestBytes)),
	)
	l := header.Len()
	headerRoof := " " + strings.Repeat("_", l-2)
	headerContent := header.String()
	h := headerRoof + "\n" + headerContent
	header.Reset()
	return h, l
}

func makeBody(cntResult counter.CntResult, largests largests) string {
	var body strings.Builder
	var ln string = "\n"
	var fileType string
	cntResultLen := len(cntResult.Info)
	largestNumDigit := len(fmt.Sprint(cntResultLen))

	for i := 0; i < cntResultLen; i++ {
		target := cntResult.Info[i]
		if ft, found := counter.FileTypeList[target.FileType]; found {
			fileType = ft[0]
		} else {
			fileType = target.FileType
		}

		if i+1 == cntResultLen {
			ln = ""
		}

		body.WriteString(fmt.Sprintf("|%d%s|  %s%s  |  %d%s  |  %d%s  |  %d%s  |  %d%s  |  %s%s  |%s",
			i+1,
			space(fmt.Sprint(i+1), largestNumDigit),
			coloring(fileType, target),
			space(fileType, largests.largestFileType),
			target.Lines,
			space(fmt.Sprint(target.Lines), largests.largestLines),
			target.Blanks,
			space(fmt.Sprint(target.Blanks), largests.largestBlanks),
			target.Comments,
			space(fmt.Sprint(target.Comments), largests.largestComments),
			target.Files,
			space(fmt.Sprint(target.Files), largests.largestFiles),
			target.Bytes,
			space(fmt.Sprint(target.Bytes), largests.largestBytes),
			ln,
		))
	}
	s := body.String()
	body.Reset()
	return s
}

func largestsNew(cntResult counter.CntResult) *largests {
	largestFileType, largestLines, largestBlanks, largestComments, largestFiles, largestBytes := largest(cntResult, counter.FileTypeList)
	return &largests{
		largestFileType: largestFileType,
		largestLines:    largestLines,
		largestBlanks:   largestBlanks,
		largestComments: largestComments,
		largestFiles:    largestFiles,
		largestBytes:    largestBytes,
	}
}

func coloring(f string, info counter.FileInfo) string {
	c := info.FileColor

	switch c {
	case "Red":
		return color.RedString(f)
	case "HiRed":
		return color.HiRedString(f)
	case "Blue":
		return color.BlueString(f)
	case "HiBlue":
		return color.HiBlueString(f)
	case "Yellow":
		return color.YellowString(f)
	case "HiYellow":
		return color.HiYellowString(f)
	case "Green":
		return color.GreenString(f)
	case "HiGreen":
		return color.HiGreenString(f)
	case "Cyan":
		return color.CyanString(f)
	case "HiCyan":
		return color.HiCyanString(f)
	case "HiBlack":
		return color.HiBlackString(f)
	case "HiWhite":
		return color.HiWhiteString(f)
	case "Magenta":
		return color.MagentaString(f)
	case "HiMagenta":
		return color.HiMagentaString(f)
	}
	return color.WhiteString(f)
}

func getNumLen(cntResult counter.CntResult) int {
	num := len(cntResult.Info)
	numLen := len(fmt.Sprint(num))
	return numLen
}