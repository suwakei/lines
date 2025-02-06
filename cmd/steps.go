package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION string = "0.1.0"

var (
	path string
	rootCmd = &cobra.Command{
		Use:   "steps [PATH] [OPTIONS]",
		Short: "steps counts the number of lines of dir or file that you input.",
		Long: `steps counts the number of lines of dir or file that you input.
		also this app can output result in [json, jsonc, yml, toml, txt] format
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
			if args[0] == "." {
				path, _ = os.Getwd()
			} else {
				path = args[0]
			}
		} else if len(args) == 0 {
			fmt.Println("[ERROR]: no path input!")
			return
		}
		fmt.Println(path)
	},

}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version of this app")
	
}


