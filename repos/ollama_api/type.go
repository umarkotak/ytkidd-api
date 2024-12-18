package ollama_api

type (
	PromptParams struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
	}

	PromptResponse struct {
		Model    string `json:"model"`
		Response string `json:"response"`
		Done     bool   `json:"done"`
	}

	ChatParams struct {
		Model    string        `json:"model"`
		Messages []ChatMessage `json:"messages"`
		Stream   bool          `json:"stream"`
	}

	ChatResponse struct {
		Model   string      `json:"model"`
		Message ChatMessage `json:"message"`
		Done    bool        `json:"done"`
	}

	ChatMessage struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
)
