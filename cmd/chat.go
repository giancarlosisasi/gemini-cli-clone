package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start interactive chat with Gemini AI",
	Long: `Start an interactive chat session with Google's Gemini AI.

Examples:
  gemini-cli-clone chat                         # Start interactive session
  gemini-cli-clone chat --debug                 # Enable debug logging
  gemini-cli-clone chat --model "gemini-pro"    # Enable debug logging
        `,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("TODO")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// specific flags for the chat
	chatCmd.Flags().StringP("model", "m", "gemini-pro", "Gemini model to use")
}
