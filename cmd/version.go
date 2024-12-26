package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION string = "0.9.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print  version of this app",
	Long: `Print version of this app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
