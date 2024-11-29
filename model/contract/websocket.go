package contract

type (
	WsBaseMessage struct {
		MessageType string `json:"message_type"`
		Data        any    `json:"data"`
	}

	WsSendChatParams struct {
		Message string `json:"message"`
	}

	WsGetChatHistories struct {
		LastID int64 `json:"last_id"`
		Limit  int64 `json:"limit"`
	}
)
