package cmd

import (
	"context"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/config"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/gemini"
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
		config := config.NewConfig()
		geminiClient, err := gemini.NewGeminiClient(config)
		if err != nil {
			return err
		}

		ctx := context.Background()

		answer, err := geminiClient.Chat(ctx, "What's the weather in Lima Peru?")
		if err != nil {
			return err
		}

		style := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			PaddingTop(1).
			PaddingLeft(4)

		fmt.Println(style.Render(answer))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// specific flags for the chat
	chatCmd.Flags().StringP("model", "m", "gemini-pro", "Gemini model to use")
}
