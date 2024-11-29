package model

type (
	OtpData struct {
		UserID   int64  `json:"user_id"`
		OtpToken string `json:"otp_token"`
		OtpValue string `json:"otp_value"`
	}
)
