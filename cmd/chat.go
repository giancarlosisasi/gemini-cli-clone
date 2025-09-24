package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/config"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/gemini"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const blue = "#1D56F4"
const purple = "#7D56F4"

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

		fmt.Print(styleTextColor("You: ", purple))

		var userMessage string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			userMessage = strings.TrimSpace(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal().Err(err).Msg("error reading the user input")

		}

		fmt.Println(styleTextColor("Gemini answer:", purple))

		ctx := context.Background()
		for chunk := range geminiClient.Chat(ctx, userMessage) {
			if chunk.Error != nil {
				log.Fatal().Err(chunk.Error).Msg("error to process answer")
			}

			if chunk.Done {
				break
			}

			fmt.Print(styleTextColor(chunk.Text, blue))
		}

		fmt.Println() // Final newline

		return nil

	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// specific flags for the chat
	chatCmd.Flags().StringP("model", "m", "gemini-pro", "Gemini model to use")
}

func styleTextColor(text string, color string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

	return style.Render(text)
}
