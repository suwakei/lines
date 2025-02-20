package view

import (
	"fmt"
    "strings"
	"github.com/suwakei/steps/counter"
	//"github.com/fatih/color"
)

func PrintTable(cntResult counter.CntResult, ignoreListMap map[string][]string) {
	fmt.Println("Target Abs Path: ", cntResult.InputPath)
	fmt.Println("All Steps: ", cntResult.AllSteps)
	fmt.Println("All Blanks: ", cntResult.AllBlanks)
	fmt.Println("All Comments: ", cntResult.AllComments)
	fmt.Println("All Bytes: ", cntResult.AllBytes)
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
	
	var (
		target counter.FileInfo
		fileType string
		cntResultLen int = len(cntResult.Info)
		largestNumDigit = len(fmt.Sprint(cntResultLen))
		largestFileType, largestSteps, largestBlanks, largestComments,largestFiles, largestBytes = largest(cntResult)
	)
	for i := 0; i < cntResultLen; i++ {
		target = cntResult.Info[i]
		if _, found := fileTypeList[target.Filetype]; found {
			fileType = fileTypeList[target.Filetype]
		} else {
			fileType = target.Filetype
		}

		largest達を用いて空白を調整する
		fmt.Printf("|%d%s|  %s%s|  %d%s|  %d%s|  %d%s|  %d%s|  %d(%dKB)%s|\n",
			i+1,
  space(i+1)
			fileType,
			target.Steps,
			target.Blanks,
			target.Comments,
			target.Files,
			target.Bytes,
			b2kb(target.Bytes),
		)
		}
	}

func space()