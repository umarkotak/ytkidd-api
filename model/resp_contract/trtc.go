package resp_contract

type (
	ClientTrtcConfig struct {
		AppID       int      `json:"app_id"`        //
		RoomID      string   `json:"room_id"`       //
		MyUserID    string   `json:"my_user_id"`    //
		MyUserSig   string   `json:"my_user_sig"`   //
		RoomUserIDs []string `json:"room_user_ids"` // valid user id list, if the ID is not exists, should not render the call
	}
)
