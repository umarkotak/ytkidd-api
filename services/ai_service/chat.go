package ai_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/contract_resp"
	"github.com/umarkotak/ytkidd-api/repos/ollama_api"
)

func SendChat(ctx context.Context, params contract.AiChat) (contract_resp.AiMessage, error) {
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
		return contract_resp.AiMessage{}, err
	}

	return contract_resp.AiMessage{
		Role:    ollamaChatResp.Message.Role,
		Content: ollamaChatResp.Message.Content,
	}, nil
}
