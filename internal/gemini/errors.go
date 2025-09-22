package gemini

import "fmt"

type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Cause   error  `json:"-"`
}

func (e *APIError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (caused by %s)", e.Type, e.Message, e.Cause)
	}

	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Cause
}

var (
	ErrAPIKeyMissing     = &APIError{Type: "auth_error", Message: "API key is missing"}
	ErrAPIKeyInvalid     = &APIError{Type: "auth_error", Message: "API key is invalid"}
	ErrModelNotFound     = &APIError{Type: "model_error", Message: "specified model not found"}
	ErrRateLimitExceeded = &APIError{Type: "rate_limit", Message: "rate limit exceeded"}
	ErrNetworkError      = &APIError{Type: "network_error", Message: "network connection failed"}
)
