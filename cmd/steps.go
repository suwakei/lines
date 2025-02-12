package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suwakei/steps/handler"



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
			files, err := handler.Search(inputPath)

			fmt.Println(exts)

			if err != nil {
				fmt.Println("[ERROR]: failed to get current directory!\n", err)
				return
			}
			inputPath := args[0]
			files, err := handler.Search(inputPath)
			if err != nil {
				fmt.Println("[ERROR]: failed to get current directory!\n", err)
				return
			}
			// search and apply ignorefile or ignore flag
			for _, file := range files {
				fmt.Println(file)

			}

			if ignoreSpecified {
				if ignoreFile == "" {
					ignoreFile = ".gitignore"
					ignoreList, err = handler.MakeIgnoreList(ignoreFile)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore list!\n", err)
						return
					}
				} else {
					ignoreList, err = handler.MakeIgnoreList(ignoreFile)
					if err != nil {
						fmt.Println("[ERROR]: failed to make ignore list!\n", err)
						return
					}
				}
				for _, i := range ignoreList {
					fmt.Println(i)
				}
			}
			// search and apply ignorefile or ignore flag
			for _, file := range files {
				fmt.Println(file)
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
	rootCmd.Flags().StringP("ext", "e", "", "input extension you don't want to count. (default: none)")
}



