package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/suwakei/steps/counter"
	"github.com/suwakei/steps/pathHandler"
	"github.com/suwakei/steps/view"
)

const VERSION string = "1.0.0"

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

			if len(args) == 0 {
				fmt.Println("[ERROR]: no path input!")
				return
			}

			inputPath, err := pathHandler.Parse(args[0])
			if err != nil {
				fmt.Println("[ERROR]: failed to parse input path!\n", err)
				return
			}

			ignoreFile, _ := cmd.Flags().GetString("ignore")
			exts, _ := cmd.Flags().GetStringSlice("ext")
			dists, _ := cmd.Flags().GetStringSlice("dist")

			ignoreListMap := map[string][]string{
				"file": {
					".exe",
					".com",
					".dll",
					".so",
					".dylib",
					".xls",
					".xlsx",
					"xlsm",
					".pdf",
					".doc",
					".docx",
					".ppt",
					".pptx",
					".msi",
					".jar",
					".gz",
					".tar",
					".png",
					".jpg",
					".jpeg",
					".svg",
					".gif",
					".bmp",
					".tiff",
					".webp",
				},
			}

			if cmd.Flags().Changed("ignore") {
				if ignoreFile == "" {
					ignoreFile = ".gitignore"
				}
				temp, err := pathHandler.MakeIgnoreList(ignoreFile)
				if err != nil {
					fmt.Println("[ERROR]: failed to make ignore listMap!\n", err)
					return
				}
				ignoreListMap["file"] = append(ignoreListMap["file"], temp["file"]...)
				ignoreListMap["dir"] = append(ignoreListMap["dir"], temp["dir"]...)
			}

			if cmd.Flags().Changed("ext") {
				temp, err := pathHandler.MakeIgnoreList(exts)
				if err != nil {
					fmt.Println("[ERROR]: failed to make ignore listMap!\n", err)
					return
				}
				ignoreListMap["file"] = append(ignoreListMap["file"], temp["file"]...)
				ignoreListMap["dir"] = append(ignoreListMap["dir"], temp["dir"]...)
			}

			files, err := pathHandler.Search(inputPath, ignoreListMap)
			fmt.Println(files)
			if err != nil {
				fmt.Println("[ERROR]: failed to get current directory!\n", err)
				return
			}

			if cmd.Flags().Changed("dist") {
				if len(dists) == 0 {
				} else {
					for i := 0; i < len(dists); i++ {
						dists[i], err = pathHandler.Parse(dists[i])
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			}
			countResult, err := counter.Count(files, inputPath)
			if err != nil {
				log.Fatal(err)
			}

			view.Write(countResult, dists, ignoreListMap)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("[ERROR]:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version of this app")
	rootCmd.Flags().StringSliceP("dist", "d", []string{}, "input filepath to output. output format [.json, .jsonc, .yml, .yaml, .toml, .txt]")
	rootCmd.Flags().StringP("only", "o", "", "By specifying an extension or file name, only files with that extension or name are targeted. \"-o=*.go\" or \"-o *.go\" or \"-o=test.txt\"")
	rootCmd.Flags().StringP("ignore", "i", "", "input your .gitignore file path. ignore extentions in .gitignore file. (default: .gitignore)")
	rootCmd.Flags().StringSliceP("ext", "e", []string{}, "input extension you don't want to count \"-e=test.json, *.js, *.go\" or \"-e=test.json -e=*.js -e=*.go\". (default: *.exe, *.com, *.dll, *.so, *.dylib, *.xls, *.xlsx, *.pdf, *.doc, *.docx, *.ppt, *.pptx)")
}
