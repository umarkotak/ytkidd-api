package contract

type (
	GenChatToken struct {
		UserGuid   string
		ToUserGuid string
		ChatRoomID int64
	}

	GetByChatRoomIDForHistory struct {
		ChatRoomID int64 `db:"chat_room_id"`
		LastID     int64 `db:"last_id"`
		Limit      int64 `db:"limit"`
	}

	GetChatRoomList struct {
		UserGuid string `db:"user_guid"`
		Limit    int64  `db:"limit"`
		Page     int64  `db:"page"`
		Offset   int64  `db:"offset"`

		UserID int64 `db:"user_id"`
	}
)
