package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suwakei/steps/pathHandler"



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
			var ignoreList []string
			inputPath := args[0]
			//outputPath, _ := cmd.Flags().GetString("output")
			ignoreFile, _ := cmd.Flags().GetString("ignore")
			exts, _ := cmd.Flags().GetStringSlice("ext")

			ignoreSpecified := cmd.Flags().Changed("ignore")
			extSpecified := cmd.Flags().Changed("ext")
			var err error

			if ignoreSpecified {
				if ignoreFile == "" {
					ignoreFile = ".gitignore"
					ignoreList, err = pathHandler.MakeIgnoreList[string](ignoreFile)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore list! with ignore flag\n", err)
						return
					}
				} else {
					ignoreList, err = pathHandler.MakeIgnoreList[string](ignoreFile)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore list!\n", err)
						return
					}
				}
			}

			if extSpecified {
				if len(ignoreList) == 0 {
					ignoreList, err = pathHandler.MakeIgnoreList[[]string](exts)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore list! with ext flag\n", err)
						return
					}
				} else {
					temp, err := pathHandler.MakeIgnoreList[[]string](exts)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore list! with ext flag\n", err)
						return
					}
					ignoreList = append(ignoreList, temp...)
				}
			}

			// search and apply ignorefile or ignore flag
			files, err := pathHandler.Search(inputPath, ignoreList)
			if err != nil {
				fmt.Println("[ERROR]: failed to get current directory!\n", err)
				return
			}

			fmt.Println(exts)
			for _, file := range files {
				fmt.Println(file)
			}
			for _, i := range ignoreList {
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
	rootCmd.Flags().StringP("output", "o", "", "input filepath to output. output format [.json, .jsonc, .yml, .yaml, .toml, .txt]")
	rootCmd.Flags().StringP("ignore", "i", "", "input your .gitignore file path. ignore extentions in .gitignore file. (default: .gitignore)")
	rootCmd.Flags().StringSliceP("ext", "e", []string{}, "input extension you don't want to count \"-e=test.json, *.js, *.go\" or \"-e=test.json -e=*.js -e=*.go\". (default: *.exe, *.com, *.dll, *.so, *.dylib, *.xls, *.xlsx, *.pdf, *.doc, *.docx, *.ppt, *.pptx)")
}



