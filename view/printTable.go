package view

import (
	"fmt"
    "strings"
	"github.com/suwakei/steps/counter"
	//"github.com/fatih/color"
)

type largests struct {
	largestFileType int
	largestSteps int
	largestBlanks int
	largestComments int
	largestFiles int
	largestBytes int
}

func PrintTable(cntResult counter.CntResult, ignoreListMap map[string][]string) {
	fmt.Println("Target Abs Path: ", cntResult.InputPath)
	fmt.Println("All Steps: ", cntResult.AllSteps)
	fmt.Println("All Blanks: ", cntResult.AllBlanks)
	fmt.Println("All Comments: ", cntResult.AllComments)
	fmt.Println("All Files: ")
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

	body, lineLen, numLen := makeBody(cntResult, *largests)
	roof, header, bottom := makeHeader(*largests, lineLen, numLen)
	footer := makeFooter(lineLen, numLen)

	fmt.Println(roof)
	fmt.Println(header)
	fmt.Println(bottom)
	for _, b := range body {
		fmt.Print(b)
	}
	fmt.Print(footer)
}

func space(currentTarget string, largestLen int) string {
    currentLen := len(currentTarget)
	var spaceNum int
	if largestLen - currentLen > -1 {
		spaceNum = largestLen - currentLen
	} else {
		spaceNum = 0
	}
    return strings.Repeat(" ", spaceNum)
}

func makeHeader(largests largests, lineLen int, numLen int) (string, string, string) {
	roof := " " + strings.Repeat("_", lineLen - 3) + " "
	numberHeader := "|" + strings.Repeat("#", numLen) + "|"

	var title string = "FileType"

	fileTypeHeader := "  " + title
	if len(title) <= largests.largestFileType {
		l := largests.largestFileType - len(title)
		fileTypeHeader += strings.Repeat(" ", l) + "  |"
	} else {
		fileTypeHeader += "  " + "|"
	}




	header := numberHeader + fileTypeHeader

	bottom := "|" + strings.Repeat("_", numLen) + "|"
	return roof, header, bottom
}

func makeBody(cntResult counter.CntResult, largests largests) (body []string, lineLen int, numLen int) {
	var (
		target counter.FileInfo
		fileType string
		cnt int
		cntResultLen int = len(cntResult.Info)
		largestNumDigit = len(fmt.Sprint(cntResultLen))
	)

	for i := 0; i < cntResultLen; i++ {
		target = cntResult.Info[i]
		if _, found := fileTypeList[target.Filetype]; found {
			fileType = fileTypeList[target.Filetype]
		} else {
			fileType = target.Filetype
		}

		if i == 0 {
			lineLen = len(fmt.Sprintf("|%d%s|  %s%s  |  %d%s  |  %d%s  |  %d%s  |  %d%s  |  %d(%dKB)%s   |\n",
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
		))
		}

		body = append(body, fmt.Sprintf("|%d%s|  %s%s  |  %d%s  |  %d%s  |  %d%s  |  %d%s  |  %d(%dKB)%s   |\n",
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
		))
		cnt++
	}
	return body, lineLen, len(fmt.Sprint(cnt))
}

func makeFooter(lineLen int, numLen int) string {
	return "|"+ strings.Repeat("_", numLen) + "|" + strings.Repeat("_", lineLen - 5) + "|"
}

func largestsNew(cntResult counter.CntResult) *largests {
	largestFileType, largestSteps, largestBlanks, largestComments,  largestFiles, largestBytes := largest(cntResult, fileTypeList)
	return &largests{
		largestFileType: largestFileType,
		largestSteps: largestSteps,
		largestBlanks: largestBlanks,
		largestComments: largestComments,
		largestFiles: largestFiles,
		largestBytes: largestBytes,
	}
}