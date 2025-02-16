package cmd

import (
	"os"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"github.com/suwakei/steps/pathHandler"
	cnter "github.com/suwakei/steps/counter"
)

const VERSION string = "0.1.0"

var (
	rootCmd = &cobra.Command{
		Use:   "steps [PATH] [OPTIONS]",
		Short: "steps counts the number of lines of dir or file that you input.",
		Long: `steps counts the number of lines of dir or file that you input.
		also this app can output result in [.json, .jsonc, .yml, .yaml, .toml, .txt] format
		and also can ignore only particular file of extension with input your .gitignore file.
		and so on ... please see flags help for more information.
		`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if version, _ := cmd.Flags().GetBool("version"); version {
				fmt.Println(VERSION)
				return 
		}

		if len(args) > 0 {
			inputPath, err := pathHandler.Parse(args[0])
			if err != nil {
				fmt.Println("[ERROR]: failed to parse input path!\n", err)
				return
			}

			//outputPath, _ := cmd.Flags().GetString("output")
			ignoreFile, _ := cmd.Flags().GetString("ignore")
			exts, _ := cmd.Flags().GetStringSlice("ext")

			ignoreSpecified := cmd.Flags().Changed("ignore")
			extSpecified := cmd.Flags().Changed("ext")

			var ignoreListMap map[string][]string = make(map[string][]string, 2)
			ignoreListMap["file"] = append(ignoreListMap["file"], "*.exe", "*.com", "*.dll", "*.so", "*.dylib", "*.xls", "*.xlsx", "*.pdf", "*.doc", "*.docx", "*.ppt", "*.pptx")

			if ignoreSpecified {
				if ignoreFile == "" {
					ignoreFile = ".gitignore"
					temp, err := pathHandler.MakeIgnoreList[string](ignoreFile)
					ignoreListMap["file"] = append(ignoreListMap["file"], temp["file"]...)
					ignoreListMap["dir"] = append(ignoreListMap["dir"], temp["dir"]...)
	
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore listMap! with ignore flag\n", err)
						return
					}

				} else {
					temp, err := pathHandler.MakeIgnoreList[string](ignoreFile)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore listMap! without ignore flag\n", err)
						return
					}
					ignoreListMap["file"] = append(ignoreListMap["file"], temp["file"]...)
					ignoreListMap["dir"] = append(ignoreListMap["dir"], temp["dir"]...)
				}
			}

			if extSpecified {
				if len(ignoreListMap["file"]) == 0 && len(ignoreListMap["dir"]) == 0 {
					temp, err := pathHandler.MakeIgnoreList[[]string](exts)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore listMap! with ext flag\n", err)
						return
					}

					ignoreListMap["file"] = append(ignoreListMap["file"], temp["file"]...)
					ignoreListMap["dir"] = append(ignoreListMap["dir"], temp["dir"]...)

				} else {
					temp, err := pathHandler.MakeIgnoreList[[]string](exts)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore list! with ext flag\n", err)
						return
					}
					ignoreListMap["file"] = append(ignoreListMap["file"], temp["file"]...)
					ignoreListMap["dir"] = append(ignoreListMap["dir"], temp["dir"]...)
				}
			}
			// search and apply ignorefile or ext flag
			files, err := pathHandler.Search(inputPath, ignoreListMap)
			if err != nil {
				fmt.Println("[ERROR]: failed to get current directory!\n", err)
				return
			}


			var result []cnter.CntResult
			lenFiles := len(files)
			if lenFiles >= 6 {
				var (
					aaa []string
					bbb []string
					ccc []string
					wg sync.WaitGroup
				)
				alen := (lenFiles+2) / 3
				blen := (lenFiles+1) / 3
				clen := (lenFiles) / 3

				aaa = make([]string, 0, alen)
				bbb = make([]string, 0, blen)
				ccc = make([]string, 0, clen)

				aaa = append(aaa, files[0:alen-1]...)
				bbb = append(bbb, files[alen:alen+blen-1]...)
				ccc = append(ccc, files[alen+blen:lenFiles-1]...)

			} else {
				for _, file := range files {
					r, err := cnter.Count(file)
					if err != nil {
						fmt.Println("[ERROR]: failed to culc line!\n", err)
						return
					}
					result = append(result, r)
				}

				fmt.Println(result)
			}

			fmt.Println(exts)
			fmt.Println("-----searchfiles-----")
			for _, file := range files {
				fmt.Println(file)
			}
			fmt.Println("-----ignoreFile-----")
			for _, i := range ignoreListMap["file"] {
				fmt.Println(i)
			}
			fmt.Println("-----ignoreDir-----")
			for _, i := range ignoreListMap["dir"] {
				fmt.Println(i)
			}

		} else if len(args) == 0 {
			fmt.Println("[ERROR]: no path input!")
			return
		}
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("[ERROR]:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version of this app")
	rootCmd.Flags().StringP("dist", "d", "", "input filepath to output. output format [.json, .jsonc, .yml, .yaml, .toml, .txt]")
	rootCmd.Flags().StringP("only", "o", "", "By specifying an extension or file name, only files with that extension or name are targeted. \"-o=*.go\" or \"-o *.go\" or \"-o=test.txt\"")
	rootCmd.Flags().StringP("ignore", "i", "", "input your .gitignore file path. ignore extentions in .gitignore file. (default: .gitignore)")
	rootCmd.Flags().StringSliceP("ext", "e", []string{}, "input extension you don't want to count \"-e=test.json, *.js, *.go\" or \"-e=test.json -e=*.js -e=*.go\". (default: *.exe, *.com, *.dll, *.so, *.dylib, *.xls, *.xlsx, *.pdf, *.doc, *.docx, *.ppt, *.pptx)")
}



