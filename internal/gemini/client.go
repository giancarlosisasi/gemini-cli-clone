package gemini

import (
	"context"
	"fmt"

	"github.com/giancarlosisasi/gemini-cli-clone/internal/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/genai"
)

type Client struct {
	client *genai.Client
	config *config.Config
}

func NewGeminiClient(config *config.Config) (*Client, error) {
	client, err := genai.NewClient(
		context.Background(),
		&genai.ClientConfig{
			APIKey:  config.GeminiAPIKey,
			Backend: genai.BackendGeminiAPI,
		},
	)

	if err != nil {
		log.Debug().Err(err).Msg("new gemini client: failed to create the gemini client")
		return nil, fmt.Errorf("failed to create the gemini client")
	}

	return &Client{
		client: client,
		config: config,
	}, nil
}

func (c *Client) Chat(ctx context.Context, message string) (string, error) {
	config := &genai.GenerateContentConfig{
		Temperature: genai.Ptr[float32](0.5),
	}
	chat, err := c.client.Chats.Create(ctx, c.config.GeminiModel, config, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Error to create chat in gemini")
	}

	result, err := chat.SendMessage(ctx, genai.Part{Text: message})
	if err != nil {
		log.Fatal().Err(err).Msg("Error to send message to gemini api")
	}

	return result.Text(), nil
}

func extractFinishReason(result *genai.GenerateContentConfig) string {
	// TODO:
	return "stop"
}

func extractTokenCount(result *genai.GenerateContentResponse) int {
	// TODO:
	return 0
}
