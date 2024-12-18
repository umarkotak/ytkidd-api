package ai_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/ollama_api"
)

func SendChat(ctx context.Context, params contract.AiChat) (resp_contract.AiMessage, error) {
	ollamaMessages := []ollama_api.ChatMessage{}
	for _, message := range params.Messages {
		ollamaMessages = append(ollamaMessages, ollama_api.ChatMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	ollamaChatResp, err := ollama_api.SendChat(ctx, ollama_api.ChatParams{
		Messages: ollamaMessages,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.AiMessage{}, err
	}

	return resp_contract.AiMessage{
		Role:    ollamaChatResp.Message.Role,
		Content: ollamaChatResp.Message.Content,
	}, nil
}
