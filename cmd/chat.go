package cmd

import (
	"github.com/giancarlosisasi/gemini-cli-clone/internal/config"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/gemini"
	"github.com/giancarlosisasi/gemini-cli-clone/internal/tui"
	"github.com/rs/zerolog/log"
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

		err = tui.InitTUI(geminiClient)
		if err != nil {
			log.Fatal().Err(err).Msg("error to run the TUI program")
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// specific flags for the chat
	chatCmd.Flags().StringP("model", "m", "gemini-pro", "Gemini model to use")
}
