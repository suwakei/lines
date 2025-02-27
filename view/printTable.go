package view

import (
    "fmt"
    "strings"
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
    fmt.Printf("All Bytes: %d(%dKB)\n", cntResult.AllBytes, b2kb(int(cntResult.AllBytes)))
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

    largests := largestsNew(cntResult)

    roof, header, body, footer := makeTable(cntResult, *largests)

    fmt.Println(roof)
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

func makeTable(cntResult counter.CntResult, largests largests) (string, string, string, string) {
    lineLen, numLen := calculateLengths(cntResult, largests)
    roof := " " + strings.Repeat("_", lineLen-3) + " "
    header := makeHeader(largests, lineLen, numLen)
    body := makeBody(cntResult, largests)
    footer := "|" + strings.Repeat("_", numLen) + "|" + strings.Repeat("_", lineLen-5) + "|"
    return roof, header, body, footer
}

func calculateLengths(cntResult counter.CntResult, largests largests) (int, int) {
    cntResultLen := len(cntResult.Info)
    largestNumDigit := len(fmt.Sprint(cntResultLen))
    lineLen := len(fmt.Sprintf("|%d%s|  %s%s  |  %d%s  |  %d%s  |  %d%s  |  %d%s  |  %d(%dKB)%s   |\n",
        1, space(fmt.Sprint(1), largestNumDigit),
        "FileType", space("FileType", largests.largestFileType),
        0, space(fmt.Sprint(0), largests.largestSteps),
        0, space(fmt.Sprint(0), largests.largestBlanks),
        0, space(fmt.Sprint(0), largests.largestComments),
        0, space(fmt.Sprint(0), largests.largestFiles),
        0, b2kb(0), space(fmt.Sprint(0), largests.largestBytes),
    ))
    return lineLen, largestNumDigit
}

func makeHeader(largests largests, lineLen int, numLen int) string {
    numberHeader := "|" + strings.Repeat("#", numLen) + "|"
    fileTypeHeader := "  FileType" + space("FileType", largests.largestFileType) + "  |"

    return numberHeader + fileTypeHeader
}

func makeBody(cntResult counter.CntResult, largests largests) string {
    var body strings.Builder
	var ln string = "\n"
	var fileType string
    cntResultLen := len(cntResult.Info)
    largestNumDigit := len(fmt.Sprint(cntResultLen))

    for i := 0; i < cntResultLen; i++ {
        target := cntResult.Info[i]
        if ft, found := fileTypeList[target.Filetype]; found {
            fileType = ft[0]
        }

		if i+1 == cntResultLen {
			ln = ""
		}

        body.WriteString(fmt.Sprintf("|%d%s|  %s%s  |  %d%s  |  %d%s  |  %d%s  |  %d%s  |  %d(%dKB)%s   |%s",
            i+1,
            space(fmt.Sprint(i+1), largestNumDigit),
            fileType,
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
            b2kb(target.Bytes),
            space(fmt.Sprint(target.Bytes), largests.largestBytes),
			ln,
        ))
    }
    return body.String()
}

func largestsNew(cntResult counter.CntResult) *largests {
    largestFileType, largestSteps, largestBlanks, largestComments, largestFiles, largestBytes := largest(cntResult, fileTypeList)
    return &largests{
        largestFileType: largestFileType,
        largestSteps: largestSteps,
        largestBlanks: largestBlanks,
        largestComments: largestComments,
        largestFiles: largestFiles,
        largestBytes: largestBytes,
    }
}