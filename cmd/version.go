package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  "Display the version, build date, and other information.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gemini-cli-clone version: %s\n", "0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
