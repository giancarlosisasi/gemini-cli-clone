package gemini

import "time"

type Response struct {
	Content      string        `json:"content"`
	Model        string        `json:"model"`
	Duration     time.Duration `json:"duration"`
	TokenCount   int           `json:"token_count"`
	FinishReason string        `json:"finish_reason"`
}

type GeminiChatStreamChunk struct {
	Text  string
	Error error
	Done  bool
}
