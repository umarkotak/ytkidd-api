package resp_contract

import "time"

type (
	StartCall struct {
		ClientTrtcConfig ClientTrtcConfig `json:"client_trtc_config"`
	}

	CallCalling struct {
		ChatRoomID    int64     `json:"chat_room_id"`
		FromUserGuid  string    `json:"from_user_guid"`
		CallSession   string    `json:"call_session"`
		CallExpiredAt time.Time `json:"call_expired_at"`
	}
)
