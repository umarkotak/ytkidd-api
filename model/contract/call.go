package contract

type (
	StartCall struct {
		ChatRoomID   int64  `json:"-"`            //
		FromUserID   int64  `json:"-"`            //
		FromUserGuid string `json:"-"`            //
		ToUserGuid   string `json:"to_user_guid"` //
		CallType     string `json:"call_type"`    // Enum: voice, video
	}
)
