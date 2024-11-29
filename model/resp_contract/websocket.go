package resp_contract

import "time"

type (
	WsChatData struct {
		ChatLogID        int64     `json:"chat_log_id"`
		FromUserGuid     string    `json:"from_user_guid"`
		SendedAt         time.Time `json:"sended_at"`
		SendedAtUnixNano int64     `json:"sended_at_unix_nano"`
		ChatType         string    `json:"chat_type"`
		Message          string    `json:"message"`

		FromMe    bool   `json:"from_me"`    // generated on handle redis chat
		AckStatus string `json:"ack_status"` //
	}

	WsChatAckData struct {
		ChatLogID int64  `json:"chat_log_id"`
		AckType   string `json:"ack_type"`
	}

	WsIncomingCallData struct {
		ChatRoomID    int64     `json:"chat_room_id"`
		FromUserGuid  string    `json:"from_user_guid"`
		CallSession   string    `json:"call_session"`
		CallExpiredAt time.Time `json:"call_expired_at"`
	}
)
