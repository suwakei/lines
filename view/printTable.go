package view

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/suwakei/steps/counter"
)

type largests struct {
    largestFileType  int
    largestSteps     int
    largestBlanks    int
    largestComments  int
    largestFiles     int
    largestBytes     int
}

func PrintTable(cntResult counter.CntResult, ignoreListMap map[string][]string) {
    fmt.Println("Target Abs Path: ", cntResult.InputPath)
    fmt.Println("All Steps: ", cntResult.AllSteps)
    fmt.Println("All Blanks: ", cntResult.AllBlanks)
    fmt.Println("All Comments: ", cntResult.AllComments)
    fmt.Println("All Files: ", cntResult.AllFiles)
    fmt.Printf("All Bytes: %s\n", cntResult.AllBytes)
    if len(ignoreListMap["file"]) != 0 {
        fmt.Println("All ignore files: ", ignoreListMap["file"])
    } else {
        fmt.Println("All ignore files: ", "None")
    }

    if len(ignoreListMap["dir"]) != 0 {
        fmt.Println("All ignore dirs: ", ignoreListMap["dir"])
    } else {
        fmt.Println("All ignore dirs: ", "None")
    }

    fmt.Println("")
    fmt.Println("")
    fmt.Println(cntResult.Info)

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

func getNumLen(cntResult counter.CntResult) int {
    num := len(cntResult.Info)
    numLen := len(fmt.Sprint(num))
    return numLen
}

func makeHeader(largests largests, numLen int) (string, int) {
    var header strings.Builder

    header.WriteString(fmt.Sprintf("|%s|  %s%s  |  %s%s  |  %s%s  |  %s%s  |  %s%s  |  %s%s  |",
    strings.Repeat("#", numLen),
    "FileType",
    space("FileType", largests.largestFileType),
    "Steps",
    space("Steps", largests.largestSteps),
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
            target.Steps,
            space(fmt.Sprint(target.Steps), largests.largestSteps),
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
    return body.String()
}

func largestsNew(cntResult counter.CntResult) *largests {
    largestFileType, largestSteps, largestBlanks, largestComments, largestFiles, largestBytes := largest(cntResult, counter.FileTypeList)
    return &largests{
        largestFileType: largestFileType,
        largestSteps: largestSteps,
        largestBlanks: largestBlanks,
        largestComments: largestComments,
        largestFiles: largestFiles,
        largestBytes: largestBytes,
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