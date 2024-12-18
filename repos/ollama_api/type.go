package ollama_api

type (
	PromptParams struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
	}

	PromptResponse struct {
		Model    string `json:"model"`
		Response string `json:"response"`
		Done     bool   `json:"done"`
	}
)
