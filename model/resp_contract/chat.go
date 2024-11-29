package resp_contract

import "time"

type (
	ChatRoomData struct {
		ID    int64  `json:"id"`
		Title string `json:"title"`

		LatestChatMessage          string    `json:"latest_chat_message"`
		LatestChatSendedAt         time.Time `json:"latest_chat_sended_at"`
		LatestChatSendedAtUnixNano int64     `json:"latest_chat_sended_at_unix_nano"`
		ThumbnailImageUrl          string    `json:"thumbnail_image_url"`
		UnreadChatCount            int64     `json:"unread_chat_count"`
	}
)
