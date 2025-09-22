package cmd

import (
	"github.com/spf13/cobra"
)

var RootFlags struct {
	Debug bool
}

var rootCmd = &cobra.Command{
	Use:   "gemini-cli-clone",
	Short: "A gemini-cli clone that I have built to learn more about CLI development in golang.",
	Long: `
A gemini-cli clone that I have built to learn more about CLI development in golang.

Examples:

gemini-cli-clone chat                       # Start interactive chat
gemini-cli-clone version                    # Show version info
gemini-cli-clone --debug                    # Enable debug mode
        `,
	Version: "0.0.1",
}

func init() {
	rootCmd.Flags().BoolVarP(&RootFlags.Debug, "debug", "d", false, "Enable debug mode")
}

func Execute() error {
	return rootCmd.Execute()
}
