package contract

type (
	AiChat struct {
		Messages []AiMessage `json:"messages"`
	}

	AiMessage struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
)
