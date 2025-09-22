package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	GeminiAPIKey   string
	GeminiModel    string
	GeminiMaxToken int
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Debug().Msg("Failed to load the env file")
	}

	viper.AutomaticEnv()

	geminiApiKey := mustGetString("GEMINI_API_KEY")
	geminiModel := mustGetString("GEMINI_MODEL")
	geminiMaxTokens := mustGetInt("GEMINI_MAX_TOKENS")

	return &Config{
		GeminiAPIKey:   geminiApiKey,
		GeminiModel:    geminiModel,
		GeminiMaxToken: geminiMaxTokens,
	}
}

func mustGetString(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Fatal().Msgf("The env var %s is missing", key)
	}

	return value
}

func mustGetInt(key string) int {
	value := viper.GetInt("GEMINI_MAX_TOKENS")
	if value == 0 {
		log.Fatal().Msgf("The env var %s is missing", key)
	}

	return value
}
